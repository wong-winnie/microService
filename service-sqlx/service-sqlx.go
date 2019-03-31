package service_sqlx

import (
	"database/sql/driver"
	"fmt"
	"gitee.com/winnie_gss/microService/service-log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
	"os"
)

var DB *sqlx.DB
var dataConfig DataConfig
var pwd string

//数据库配置
type DataConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
}

//TODO : 初始化数据库
func InitDb() {

	pwd, _ = os.Getwd()
	pwd += "//static"
	file, err := ini.Load(pwd + "//config.ini")
	service_log.ErrorLog(err)

	file.BlockMode = false
	err = file.Section("mysql").MapTo(&dataConfig)
	service_log.ErrorLog(err)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dataConfig.DbUser, dataConfig.DbPassword, dataConfig.DbHost, dataConfig.DbPort, dataConfig.DbName)
	DB, err = sqlx.Open("mysql", dsn)
	service_log.ErrorLog(err)
}

//TODO : 查询一条数据
func SelectOne(sql string, data interface{}, args ...interface{}) {
	service_log.InfoLog(fmt.Sprintln("sql : ", sql, " args :", args))
	err := DB.Get(data, sql, args...)
	service_log.ErrorLog(err)
}

//TODO : 查询所有数据
func SelectAll(sql string, data interface{}) {
	service_log.InfoLog(sql)
	err := DB.Select(data, sql)
	service_log.ErrorLog(err)
}

//TODO : 修改数据
func UpdateDate(sql string, args ...interface{}) (result driver.Result) {
	service_log.InfoLog(fmt.Sprintln("sql : ", sql, " args :", args))
	result, err := DB.Exec(sql, args...)
	service_log.ErrorLog(err)
	return
}
