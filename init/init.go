/**********************************************
** @Des:
** @Author: lg1024
** @Last Modified time: 2019/11/15 上午11:43
***********************************************/

package init

import (
	"github.com/741369/go_utils/log"
	"net/http"
	"time"

	"github.com/spf13/pflag"
	"os"
	"strings"

	"api-gin-web/model"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path.")
)

func init() {
	pflag.Parse()

	// init config
	if err := Init(*cfg); err != nil {
		panic(err)
	}

	// add monitor
	// Ping the server to make sure the router is working.
	/*go func() {
		if err := monitor.PingServer(); err != nil {
			// TODO 监控告警
			log.Error("The router has no response, or it might took too long to start up.", err)
		}
		log.Infof(nil,nil,"The router has been deployed successfully.")
	}()*/

	// init db
	go model.DB.Init()
	go model.OpenRedis()
	//defer model.DB.Close()

	// http请求过期时间
	http.DefaultClient.Timeout = time.Minute * 10
}

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	//c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf")
		//viper.AddConfigPath(os.Getenv("GOPATH") + "/src/api-gin-web/conf")
		if os.Getenv("ENV_GO") != "" {
			viper.SetConfigName("config-" + os.Getenv("ENV_GO"))
		} else {
			viper.SetConfigName("config")
		}
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // 读取匹配的环境变量
	viper.SetEnvPrefix("ONETHINGPCS")

	replcer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replcer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	/*viper.SetConfigName("extra")
	viper.AddConfigPath("conf")
	viper.MergeInConfig()*/

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

//func (c *Config) watchConfig() {
//	viper.WatchConfig()
//	viper.OnConfigChange(func(e fsnotify.Event) {
//		log.Infof(nil, "Config file changed: %s", e.Name)
//	})
//}
