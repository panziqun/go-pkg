package config

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	Name        string
	Version     string
	JwtSecret   string
	RuntimePath string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormt   string
}

type Database struct {
	Type          string
	User          string
	Password      string
	Host          string
	Port          string
	Name          string
	Charset       string
	ParseTime     string
	Loc           string
	SingularTable bool
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mail struct {
	Host     string
	Port     int
	User     string
	Password string
}

type Redis struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	Password    string
	IdleTimeout time.Duration
	Db          int
}

var cfg *ini.File

var AppConfig = &App{}
var DatabaseConfig = &Database{}
var ServerConfig = &Server{}
var MailConfig = &Mail{}
var RedisConfig = &Redis{}

func Setup() {
	var err error
	cfg, err = ini.Load("src/config/app.ini", "src/config/app.ini.local")
	if err != nil {
		fmt.Printf("fail to parse 'app.ini': %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("database", DatabaseConfig)
	mapTo("server", ServerConfig)
	mapTo("mail", MailConfig)
	mapTo("redis", RedisConfig)

	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("cfg.MapTo err: %v", err)
	}
}
