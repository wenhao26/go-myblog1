package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	Env        string
	LogFile    string
	LogConsole bool
	LogLevel   string
	AppMode    string
	HttpPort   int
	JwtKey     string

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln("配置文件加载失败：", err)
	}

	LoadServerCnf(file)
	LoadDatabaseCnf(file)
}

func LoadServerCnf(file *ini.File) {
	section := file.Section("server")
	Env = section.Key("env").MustString("dev")
	LogFile = section.Key("logFile").MustString("")
	LogConsole = section.Key("logConsole").MustBool(true)
	LogLevel = section.Key("logLevel").MustString("debug")
	AppMode = section.Key("appMode").MustString("debug")
	HttpPort = section.Key("httpPort").MustInt(3000)
	JwtKey = section.Key("jwtKey").MustString("123456")
}

func LoadDatabaseCnf(file *ini.File) {
	section := file.Section("database")
	DbHost = section.Key("dbHost").MustString("127.0.0.1")
	DbPort = section.Key("dbPort").MustInt(3306)
	DbUser = section.Key("dbUser").MustString("root")
	DbPassWord = section.Key("dbPassWord").MustString("")
	DbName = section.Key("dbName").MustString("")
}

func IsDev() bool {
	return Env == "dev"
}
