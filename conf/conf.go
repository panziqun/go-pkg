package conf

import (
	"fmt"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type AppConf struct {
	Name        string
	Version     string
	JwtSecret   string
	RuntimePath string

	LogPath      string
	LogFileExt   string
	LogTimeFormt string
	LogMail      string
	LogPanicPath string
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
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

var file *ini.File

var App = &AppConf{}
var Database = &DatabaseConf{}
var Server = &ServerConf{}
var Mail = &MailConf{}
var Redis = &RedisConf{}
var Mongodb = &MongodbConf{}

func Setup() {
	LoadSources("conf/app.ini", "conf/app.local.ini")
}

func LoadSources(source string, local string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err = ini.Load(wd+"/"+source, wd+"/"+local)
	if err != nil {
		fmt.Printf("fail to parse 'app.ini': %v", err)
		panic(err)
	}

	mapTo("app", App)
	mapTo("database", Database)
	mapTo("server", Server)
	mapTo("mail", Mail)
	mapTo("redis", Redis)
	mapTo("mongodb", Mongodb)

	Server.ReadTimeout = Server.ReadTimeout * time.Second
	Server.WriteTimeout = Server.WriteTimeout * time.Second
	Redis.IdleTimeout = Redis.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := file.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("file.MapTo err: %v", err)
		panic(err)
	}
}

func Section(name string) map[string]string {
	sec := file.Section(name)
	return sec.KeysHash()
}
