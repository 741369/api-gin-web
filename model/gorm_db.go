package model

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	TestDB *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	//config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Local")
		"Asia%2fShanghai")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Printf("open mysql db error,  db_name = %s, config = %s, err = %v \n", name, config, err)
	}
	// set for db connection
	setupDB(db)

	return db
}

func openDB2(username, password, addr, name, charset string) *gorm.DB {
	//config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		charset,
		true,
		"Asia%2fShanghai")
	//"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Printf("open mysql db error,  db_name = %s, config = %s, err = %v \n", name, config, err)
	}
	// set for db connection
	setupDB(db)

	return db
}

type FileModel struct {
	Id       int
	FileName string
	FileSize int
}

func setupDB(db *gorm.DB) {
	if os.Getenv("ENV_GO") != "" {
		db.LogMode(true) // 显示mysql日志
	}
	db.DB().SetMaxOpenConns(1000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(100)  // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func (db *Database) Init() {
	DB = &Database{
		TestDB: InitTestDB(),
	}
}

func InitTestDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func (db *Database) OtherDB(dbName string) *gorm.DB {
	openDB := openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"), dbName)
	return openDB
}

func (db *Database) OtherDB3(dbName string) *gorm.DB {
	openDB := openDB(viper.GetString("db.username3"),
		viper.GetString("db.password3"),
		viper.GetString("db.addr3"), dbName)
	return openDB
}

func (db *Database) Close() {
	DB.TestDB.Close()
}

// 组装表
func TableName(tableName string, flag int) string {
	if flag < 0 || flag > 99 {
		return tableName
	}
	return tableName + "_" + strconv.Itoa(flag)
}
