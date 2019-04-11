package bootconfig

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

// Name for adapter with beego official support
const (
	AdapterConsole = "console"
	AdapterFile    = "file"
)

var Log = logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
//初始化日志配置
func LogInit() {
	var cfg = beego.AppConfig
	config := make(map[string]interface{})
	logAdapter := cfg.String("log.engine")
	if logAdapter == AdapterConsole {
		color, _ := cfg.Bool("log.console.color")
		config["color"] = color
	} else if logAdapter == AdapterFile {
		config["filename"] = cfg.String("log.file.filename")
		config["maxlines"], _ = cfg.Int("log.file.maxlines")
		config["maxsize"],_ = cfg.Int("log.file.maxsize")
		config["daily"], _ = cfg.Bool("log.file.daily")
		config["maxdays"], _ = cfg.Int("log.file.maxdays")
		config["rotate"], _ = cfg.Bool("log.file.rotate")
	} else {
		fmt.Println("not support log adapter")
		os.Exit(1)
	}
	// map 转 json
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("logInit failed, marshal err:", err)
		os.Exit(1)
	}
	err = Log.SetLogger(logAdapter, string(configStr)) // 设置日志记录方式：控制台记录
	if err != nil {
		fmt.Println("logInit failed, set logger err:", err)
		os.Exit(1)
	}

	logLever := cfg.String("log.level")
	adapterLevel := adapterLogLevel(logLever)
	Log.SetLevel(adapterLevel)    // 设置日志写入缓冲区的等级：Debug级别（最低级别，所以所有log都会输入到缓冲区）
	Log.EnableFuncCallDepth(true) // 输出log时能显示输出文件名和行号（非必须）
	Log.Async(1e3)
	Log.Debug("set logger success")
}

//用配置文件中配置的日志级别适配beego的日志级别
func adapterLogLevel(configLevel string) int {
	var adapterLevel int
	switch configLevel {
	case "info":
		adapterLevel = logs.LevelInfo
	case "debug":
		adapterLevel = logs.LevelDebug
	case "error":
		adapterLevel = logs.LevelError
	case "warning":
		adapterLevel = logs.LevelWarn
	default:
		adapterLevel = logs.LevelInfo
	}
	return adapterLevel
}
