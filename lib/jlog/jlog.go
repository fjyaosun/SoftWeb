package jlog

import (
	"SoftWeb/lib/jsmtp"
	"flag"
	"fmt"
	"runtime"

	"github.com/golang/glog"
)

func init() {
	configInit()
	initLog()
}

func initLog() {
	flag.Set("alsologtostderr", logConf.alsoLogToStderr)
	flag.Set("log_dir", logConf.logDir)
	flag.Set("v", logConf.v)
	flag.Parse()

}

func Call() string {
	buf := make([]byte, 1<<20)
	buf = buf[:runtime.Stack(buf, true)]
	logStr := fmt.Sprintf("PANIC STASK: %s", buf)
	return logStr
}

func LogError(err error) string {
	logStr := fmt.Sprintf("ERROR: %s\n%s", err.Error(), Call())
	glog.Errorln(logStr)
	glog.Flush()
	return logStr
}

func LogErrorSendMail(err error) {
	jsmtp.SendDefaultMail(LogError(err))
}

func LogWarnning(err error) {
	logStr := fmt.Sprintf("WARNNING: %s\n%s", err.Error(), Call())
	glog.Warningln(logStr)
	glog.Flush()
}

func LogInfo(err error) {
	logStr := fmt.Sprintf("INFO: %s\n%s", err.Error(), Call())
	glog.Infoln(logStr)
	glog.Flush()
}
