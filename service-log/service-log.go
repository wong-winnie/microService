package service_log

import (
	"fmt"
	"runtime"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

var logConfig LogConfig
var folderPath string
var filePath string
var pwd string

type LogConfig struct {
	LogLevel int
}

//初始化日记库
func InitLog() {
	pwd, _ = os.Getwd()
	pwd += "//static"
	file, err := ini.Load(pwd + "//config.ini")
	simplePanic(err)

	file.BlockMode = false
	err = file.Section("log").MapTo(&logConfig)
	simplePanic(err)
}

//调试信息
func DebugLog(str string) {
	if logConfig.LogLevel > 1 {
		return
	}
	WriteLog(str, "Debug")
}

//运行信息
func InfoLog(str string) {
	if logConfig.LogLevel > 2 {
		return
	}
	WriteLog(str, "Info")
}

//错误信息
func ErrorLog(err error) {
	if err != nil {
		WriteLog(err, "Error")
	}
}

func WriteLog(str interface{}, level string) {

	//生成文件
	folderPath = pwd + "//runtime//" + time.Now().Format("200601")
	filePath = folderPath + "//" + time.Now().Format("02.log")

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0777)
	}

	//写入日记文件
	_, files, line, _ := runtime.Caller(2)

	//1-excute 2-write 4-read
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	simplePanic(err)

	content := fmt.Sprintf("[ %s ] %s : %d\r\n[ %s ] %s\r\n---------------------------------------------------------------\r\n", time.Now().Format("2006-01-02 15:04:05"), files, line, level, str)
	file.Write([]byte(content))
}

func simplePanic(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println("----------------------------------------")
		fmt.Println(file, line)
		fmt.Println(err)
		fmt.Println("----------------------------------------")
		runtime.Goexit()
	}
}
