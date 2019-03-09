package conf

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type ModuleConf struct {
	Database bool
	Mail     bool
	Redis    bool
	Mongodb  bool
	Log      bool
}

type AppConf struct {
	Name        string
	Version     string
	JwtSecret   string
	RuntimePath string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormt   string
}

type DatabaseConf struct {
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

type ServerConf struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MailConf struct {
	Host     string
	Port     int
	User     string
	Password string
}

type RedisConf struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	Password    string
	IdleTimeout time.Duration
	Db          int
}

type MongodbConf struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}

var cfg *ini.File

var Module = &ModuleConf{}
var App = &AppConf{}
var Database = &DatabaseConf{}
var Server = &ServerConf{}
var Mail = &MailConf{}
var Redis = &RedisConf{}
var Mongodb = &MongodbConf{}

func Setup() {
	var err error
	cfg, err = ini.Load("src/conf/app.ini", "src/conf/app.ini.local")
	if err != nil {
		fmt.Printf("fail to parse 'app.ini': %v", err)
	}

	mapTo("app", App)
	mapTo("database", Database)
	mapTo("server", Server)
	mapTo("mail", Mail)
	mapTo("redis", Redis)
	mapTo("mongodb", Mongodb)
	mapTo("module", Module)

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second
	Redis.IdleTimeout = Redis.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("cfg.MapTo err: %v", err)
	}
}
