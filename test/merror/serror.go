package merror

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"runtime"
	"strconv"
	"v_tools/log"
	"vtoken_digiccy_go/test/config"
)

type ErrorLevel int

const (
	LogLevel ErrorLevel = iota
	WarnLevel
	CrashLevel
)

var errorLog *log.Log

func init() {
	errorLog = log.Init("20060102.error")
}

func errorLogOutput(level ErrorLevel, a interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if level == LogLevel {
		//output log
		logrus.Info(a, file, line)
		errorLog.Println(a, file, line)
	} else if level == WarnLevel {
		logrus.Warn(a, file, line)
		errorLog.Println(a, file, line)
		go SendEmail("warn", file+"--"+strconv.Itoa(line)+"--"+fmt.Sprint(a), config.ConsulCfg.AlertEmail)
	} else if level == CrashLevel {
		logrus.Warn(a, file, line)
		errorLog.Println(a, file, line)
		go SendEmail("crash", file+"--"+strconv.Itoa(line)+"--"+fmt.Sprint(a), config.ConsulCfg.AlertEmail)
		//TODO: 企业微信推送
	}
}

//推荐格式为 merchant,uid,...interface
func Log(a ...interface{}) {
	errorLogOutput(LogLevel, a)
}

//推荐格式为 merchant,uid,...interface
func Warn(a ...interface{}) {
	errorLogOutput(WarnLevel, a)
}

//推荐格式为 merchant,uid,...interface
func Crash(a ...interface{}) {
	errorLogOutput(CrashLevel, a)
}
