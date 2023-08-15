package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var TikTok Config

type (
	Config struct {
		server
		app
		jwt
		dataSource
		oss
	}

	server struct {
		RunMode      string        //运行模式
		HTTPPort     string        //HTTP端口
		ReadTimeout  time.Duration //读取超时时间
		WriteTimeout time.Duration //写入超时时间
	}

	app struct {
		DefaultPageSize       uint32        //默认分页大小
		MaxPageSize           uint32        //最大页大小
		DefaultContextTimeout time.Duration //默认上下文超时时间
		Log                   struct {
			SavePath   string //日志保存路径
			FileName   string //日志文件名
			FileExt    string //日志文件拓展名
			MaxSize    int    //日志文件最大大小
			MaxBackUps int    //保留日志文件的最大个数
			MaxAge     int    //保留日志文件的最大天数
		}
	}

	jwt struct {
		Secret     string        //JWT密钥
		Issuer     string        //JWT发行
		Timeout    time.Duration //过期时间
		MaxRefresh time.Duration //最大刷新时间
	}

	dataSource struct {
		MySQL struct {
			Host        string //主机地址
			UserName    string //用户名
			Password    string //密码
			DBType      string //数据库类型
			DBName      string //数据库名称
			MaxOpenConn int    //最大打开连接数
			MaxIdleConn int    //最大限制连接数
			Charset     string //字符集
			ParseTime   string
			TimeLocal   string //时间区域
		}
		Redis struct {
			Addr     string //redis服务地址
			Password string //redis密码
			DBIndex  int    //数据库下标
		}
	}

	oss struct {
		AccessKey string
		SecretKey string
		URL       string
	}
)

func init() {
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("no such config file: %v\n", err)
		} else {
			log.Fatalf("read config errof: %v\n", err)
		}
	}
	if err := vp.UnmarshalKey("dataSource", &TikTok.dataSource); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	if err := vp.UnmarshalKey("app", &TikTok.app); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	if err := vp.UnmarshalKey("server", &TikTok.server); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}
	TikTok.server.ReadTimeout *= time.Second
	TikTok.server.WriteTimeout *= time.Second

	if err := vp.UnmarshalKey("auth", &TikTok.jwt); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}
	TikTok.Timeout *= time.Minute
	TikTok.MaxRefresh *= time.Minute

	if err := vp.UnmarshalKey("oss", &TikTok.oss); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}

	if err := vp.UnmarshalKey("jwt", &TikTok.jwt); err != nil {
		log.Fatalf("读取配置文件失败:%v\n", err)
	}
	TikTok.Timeout *= time.Hour
	TikTok.MaxRefresh *= time.Hour
}
