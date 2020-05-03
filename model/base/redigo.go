package base

import (
	"strconv"
	"time"

	"github.com/741369/go_utils/log"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

/*type RedisDB struct {
	RedisConn redis.Conn
}

var RC *RedisDB

func OpenRedis() *RedisDB {
	pool := redis.Pool{
		MaxIdle:   viper.GetInt("redis.maxidle"),
		MaxActive: viper.GetInt("redis.maxactive"), // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("redis.host"),
				redis.DialDatabase(viper.GetInt("redis.db")),
				redis.DialPassword(viper.GetString("redis.password")))
			if err != nil {
				log.Infof(nil, "OpenRedis connect redis error = %v", err)
			}
			return c, err
		},
	}
	//RC = &RedisDB{RedisConn: pool.Get()}
	//return RC
	return &RedisDB{RedisConn: pool.Get()}
}*/

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func OpenRedis() error {
	RedisConn = &redis.Pool{
		MaxIdle:     viper.GetInt("redis.maxidle"),
		MaxActive:   viper.GetInt("redis.maxactive"), // max number of connections
		IdleTimeout: 60 * time.Second,
		//IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("redis.host"),
				redis.DialDatabase(viper.GetInt("redis.db")),
				redis.DialPassword(viper.GetString("redis.password")))
			if err != nil {
				log.Infof(nil, "redis_connect_error, err = %v", err)
				return nil, err
			}
			if viper.GetString("redis.password") != "" {
				if _, err := c.Do("AUTH", viper.GetString("redis.password")); err != nil {
					log.Infof(nil, "redis_auth_error, err = %v", err)
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Infof(nil, "redis_ping_error, err = %v", err)
				return err
			} else {
				return nil
			}
		},
	}

	return nil
}

func RedisGet(key string, db ...int) string {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	v, err := redis.String(conn.Do("GET", key))
	if err != nil {
		//log.Infof("[RedisGet] get error, err = %v", err)
		return ""
	}
	return v
}

func RedisSet(key, value string, expire int, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	if expire > 0 { // 设置有效期
		_, err := conn.Do("SETEX", key, expire, value)
		return checkError(err)
	} else { // 长久有效
		_, err := conn.Do("SET", key, value)
		return checkError(err)
	}
}

func RedisIncr(key string, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	_, err := conn.Do("INCR", key)
	return checkError(err)
}

func RedisDel(key string, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	_, err := conn.Do("DEL", key)
	return checkError(err)
}

func RedisBatchSet(data map[string]interface{}, expire int, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	for k, v := range data {
		if expire > 0 { // 设置有效期
			_, err := conn.Do("SETEX", k, expire, v)
			if err != nil {
				log.Infof(nil, "RedisBatchSet redis operation error SETEX %v", err)
				return false
			}
		} else { // 长久有效
			_, err := conn.Do("SET", k, v)
			if err != nil {
				log.Infof(nil, "RedisBatchSet redis operation error SET %v", err)
				return false
			}
		}
	}
	return true
}

// getbit
func RedisGetBit(key, offset string, db ...int) int {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	v, err := conn.Do("GETBIT", key, offset)
	if err != nil {
		//log.Infof("[RedisGetBit] get error, err = %v", err)
		return 0
	}
	//vInt, err := utils.InterfaceToInt(v)
	vInt, err := strconv.Atoi(v.(string))
	if err != nil || vInt == 0 {
		//log.Infof("[RedisGetBit] err = %v, v = %v", err, v)
		return 0
	}
	return vInt
}

// setbit
func RedisSetBit(key, offset string, value int, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	_, err := conn.Do("SETBIT", key, offset, value)
	if err != nil {
		log.Infof(nil, "RedisSetBit %s, %s, %d, %v", key, offset, value, err)
		return false
	}
	return true
}

// getbit
func RedisBitCount(key, offset string, db ...int) int {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	v, err := conn.Do("GETBIT", key, offset)
	if err != nil {
		//log.Infof("[RedisGetBit] get error, err = %v", err)
		return 0
	}
	vInt, err := strconv.Atoi(v.(string))
	//vInt, err := utils.InterfaceToInt(v)
	if err != nil || vInt == 0 {
		//log.Infof("[RedisGetBit] err = %v, v = %v", err, v)
		return 0
	}
	return vInt
}

// zadd
func RedisZadd(key, member, score string, db ...int) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	_, err := conn.Do("ZADD", key, score, member)
	if err != nil {
		log.Infof(nil, "RedisZadd %s, %s, %s, %v", key, score, member, err)
		return false
	}
	return true
}

// zrevrank 查询排行榜
func RedisZrevrank(key, member string, db ...int) (int, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	v, err := conn.Do("ZREVRANK", key, member)
	if err != nil {
		// TODO 告警
		log.Infof(nil, "RedisZrevrank redis 操作失败, %s, %s, %v", key, member, err)
		return 0, err
	}
	if v != nil {
		//vInt, err := utils.InterfaceToInt(v)
		vInt, err := strconv.Atoi(v.(string))
		if err != nil {
			log.Infof(nil, "RedisZrevrank interface to int error, %s, %s, %v", key, member, err)
			return 0, err
		}
		return vInt + 1, nil
	} else {
		return 0, nil
	}
}

// zrevrange 查询前几名集合
func RedisZrevrange(key string, start, stop int, db ...int) ([]map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		conn.Do("SELECT", db[0])
	}
	// 设置
	res, err := redis.Values(conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		// TODO 告警
		log.Infof(nil, "RedisZrevrange redis 操作失败 %s, %d, %d, %v", key, start, stop, err)
		return nil, err
	}
	resData := make([]map[string]string, 0)
	tmp := make(map[string]string, 0)
	for i, v := range res {
		if i%2 == 0 { // 偶数为key，基数为value
			tmp["key"] = string(v.([]byte))
		} else {
			tmp["value"] = string(v.([]byte))
			resData = append(resData, tmp)
			tmp = make(map[string]string, 0)
		}
	}
	return resData, nil
}

/*func (rc *RedisDB) RedisGet(key string, db ...int) string {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	v, err := redis.String(rc.RedisConn.Do("GET", key))
	if err != nil {
		//log.Infof("[RedisGet] get error, err = %v", err)
		return ""
	}
	return v
}

func (rc *RedisDB) RedisSet(key, value string, expire int, db ...int) bool {

	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	if expire > 0 { // 设置有效期
		_, err := rc.RedisConn.Do("SETEX", key, expire, value)
		return checkError(err)
	} else { // 长久有效
		_, err := rc.RedisConn.Do("SET", key, value)
		return checkError(err)
	}
}

func (rc *RedisDB) RedisIncr(key string, db ...int) bool {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	_, err := rc.RedisConn.Do("INCR", key)
	return checkError(err)
}

func (rc *RedisDB) RedisDel(key string, db ...int) bool {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	_, err := rc.RedisConn.Do("DEL", key)
	return checkError(err)
}

func (rc *RedisDB) RedisBatchSet(data map[string]interface{}, expire int, db ...int) bool {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	for k, v := range data {
		if expire > 0 { // 设置有效期
			_, err := rc.RedisConn.Do("SETEX", k, expire, v)
			if err != nil {
				log.Infof(nil, "RedisBatchSet redis operation error SETEX %v", err)
				return false
			}
		} else { // 长久有效
			_, err := rc.RedisConn.Do("SET", k, v)
			if err != nil {
				log.Infof(nil, "RedisBatchSet redis operation error SET %v", err)
				return false
			}
		}
	}
	return true
}

// getbit
func (rc *RedisDB) RedisGetBit(key, offset string, db ...int) int {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	v, err := rc.RedisConn.Do("GETBIT", key, offset)
	if err != nil {
		//log.Infof("[RedisGetBit] get error, err = %v", err)
		return 0
	}
	vInt, err := utils.InterfaceToInt(v)
	if err != nil || vInt == 0 {
		//log.Infof("[RedisGetBit] err = %v, v = %v", err, v)
		return 0
	}
	return vInt
}

// setbit
func (rc *RedisDB) RedisSetBit(key, offset string, value int, db ...int) bool {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	_, err := rc.RedisConn.Do("SETBIT", key, offset, value)
	if err != nil {
		log.Infof(nil, "RedisSetBit %s, %s, %d, %v", key, offset, value, err)
		return false
	}
	return true
}

// getbit
func (rc *RedisDB) RedisBitCount(key, offset string, db ...int) int {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	v, err := rc.RedisConn.Do("GETBIT", key, offset)
	if err != nil {
		//log.Infof("[RedisGetBit] get error, err = %v", err)
		return 0
	}
	vInt, err := utils.InterfaceToInt(v)
	if err != nil || vInt == 0 {
		//log.Infof("[RedisGetBit] err = %v, v = %v", err, v)
		return 0
	}
	return vInt
}

// zadd
func (rc *RedisDB) RedisZadd(key, member, score string, db ...int) bool {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	_, err := rc.RedisConn.Do("ZADD", key, score, member)
	if err != nil {
		log.Infof(nil, "RedisZadd %s, %s, %s, %v", key, score, member, err)
		return false
	}
	return true
}

// zrevrank 查询排行榜
func (rc *RedisDB) RedisZrevrank(key, member string, db ...int) (int, error) {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	v, err := rc.RedisConn.Do("ZREVRANK", key, member)
	if err != nil {
		// TODO 告警
		log.Infof(nil, "RedisZrevrank redis 操作失败, %s, %s, %v", key, member, err)
		return 0, err
	}
	if v != nil {
		vInt, err := utils.InterfaceToInt(v)
		if err != nil {
			log.Infof(nil, "RedisZrevrank interface to int error, %s, %s, %v", key, member, err)
			return 0, err
		}
		return vInt + 1, nil
	} else {
		return 0, nil
	}
}

// zrevrange 查询前几名集合
func (rc *RedisDB) RedisZrevrange(key string, start, stop int, db ...int) ([]map[string]string, error) {
	if len(db) == 1 && db[0] >= 0 && db[0] <= 16 {
		rc.RedisConn.Do("SELECT", db[0])
	}
	// 设置
	res, err := redis.Values(rc.RedisConn.Do("ZREVRANGE", key, start, stop, "WITHSCORES"))
	if err != nil {
		// TODO 告警
		log.Infof(nil, "RedisZrevrange redis 操作失败 %s, %d, %d, %v", key, start, stop, err)
		return nil, err
	}
	resData := make([]map[string]string, 0)
	tmp := make(map[string]string, 0)
	for i, v := range res {
		if i%2 == 0 { // 偶数为key，基数为value
			tmp["key"] = string(v.([]byte))
		} else {
			tmp["value"] = string(v.([]byte))
			resData = append(resData, tmp)
			tmp = make(map[string]string, 0)
		}
	}
	return resData, nil
}*/

func checkError(err error) bool {
	if err != nil {
		log.Infof(nil, "checkError %v", err)
		return false
	}
	return true
}
