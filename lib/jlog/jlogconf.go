package jlog

var logConf struct {
	logDir          string
	alsoLogToStderr string
	v               string
}

func configInit() {
	logConf.logDir = "./log"
	logConf.alsoLogToStderr = "true"
	logConf.v = "3"
}
