package main

import (
	"flag"
	"time"
	"runtime"
	"github.com/Terry-Mao/goconf"
)

var (Conf *Config
	gconf *goconf.Config
	confFile string)

func init(){
	flag.StringVar(&confFile, "c", "./mysql.conf", "set mysql connect config")
}

type Config struct {
	// base section
	PidFile  string `goconf:"base:pidfile"`
	Dir      string `goconf:"base:dir"`
	Log      string `goconf:"base:log"`
	MaxProc  int    `goconf:"base:maxproc"`
	Debug    bool   `goconf:"base:debug"`
	HttpHost string `goconf:"http:host"`

	RedisAddrs        string      `goconf:"redis:addrs:,"`
	RedisDialTimeout  time.Duration `goconf:"redis:dial.timeout:time"`
	RedisReadTimeout  time.Duration `goconf:"redis:read.timeout:time"`
	RedisWriteTimeout time.Duration `goconf:"redis:write.timeout:time"`
	RedisPoolSize     int           `goconf:"redis:pool.size"`
	RedisPoolTimeout  time.Duration `goconf:"redis:pool.timeout:time"`

	MySqlIP string `goconf:"mysql:ip"`
	MySqlPort string `goconf:"mysql:port"`
	MysqlUser string `goconf:"mysql:user"`
	MysqlDb string `goconf:"mysql:db"`
	MysqlPasswd string `goconf:"mysql:passwd"`
	MysqlMaxConn string `goconf:"mysql:max.conn"`
	MysqlMaxIdle string `goconf:"mysql:max.idle"`
}

func NewConfig() *Config{
	return &Config{
		PidFile: "/tmp/login.pid",
		Dir: "./",
		Log:  "./login-log.xml",
		MaxProc: runtime.NumCPU(),
		Debug: true,
		HttpHost: "127.0.0.1:8080",

		RedisAddrs: "127.0.0.1:6379//0",
		RedisDialTimeout:  10 * time.Second,
		RedisReadTimeout:  2 * time.Second,
		RedisWriteTimeout: 2 * time.Second,
		RedisPoolSize:     10,
		RedisPoolTimeout:  30 * time.Second,
	}
}

func InitConfig()(err error){
	Conf=NewConfig()
	gconf=goconf.New()
	if err=gconf.Parse(&confFile);err != nil{
		return err
	}
	if err = gconf.Unmarshal(Conf);err != nil{
		return err
	}
	return nil
}

func ReloadConfig()(*Config,error){
	Conf:=NewConfig()
	ngconf,err :=gconf.Reload()
	if err != nil{
		return nil,err
	}
	if err:=gconf.Unmarshal(Conf);err!=nil{
		return nil,err
	}
	gconf=ngconf
	return Conf,nil

}
