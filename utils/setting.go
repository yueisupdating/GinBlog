package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JWTKEY   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone       int
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Panic("未能正确读取config.ini文件", err)
	}
	InitServer(file)
	InitDatabase(file)
	InitQiniu(file)
}
func InitServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JWTKEY = file.Section("server").Key("JWTKEY").MustString("AZ&K4(hjIU")
}
func InitDatabase(file *ini.File) {
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbServer").MustString("localhost")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("ABCDE0000")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
func InitQiniu(file *ini.File) {
	Zone = file.Section("qiniu").Key("Zone").MustInt(1)
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("")
}
