package conf

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type Module struct {
	Database bool
	Mail     bool
	Redis    bool
	Mongodb  bool
	Log      bool
}

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

type Mongodb struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}

var cfg *ini.File

var ModuleConf = &Module{}
var AppConf = &App{}
var DatabaseConf = &Database{}
var ServerConf = &Server{}
var MailConf = &Mail{}
var RedisConf = &Redis{}
var MongodbConf = &Mongodb{}

func Setup() {
	var err error
	cfg, err = ini.Load("src/conf/app.ini", "src/conf/app.ini.local")
	if err != nil {
		fmt.Printf("fail to parse 'app.ini': %v", err)
	}

	mapTo("app", AppConf)
	mapTo("database", DatabaseConf)
	mapTo("server", ServerConf)
	mapTo("mail", MailConf)
	mapTo("redis", RedisConf)
	mapTo("mongodb", MongodbConf)
	mapTo("module", ModuleConf)

	ServerConf.ReadTimeout = ServerConf.ReadTimeout * time.Second
	ServerConf.WriteTimeout = ServerConf.WriteTimeout * time.Second
	RedisConf.IdleTimeout = RedisConf.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("cfg.MapTo err: %v", err)
	}
}
