package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Panic("未能正确读取config.ini文件", err)
	}
	InitServer(file)
	InitDatabase(file)
}
func InitServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}
func InitDatabase(file *ini.File) {
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbServer").MustString("localhost")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("ABCDE0000")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
