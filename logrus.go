package main

import "github.com/sirupsen/logrus"

// go的另一个日志库，功能强大，完全兼容标准的log库，docker中也使用了该库

// logrus支持更多的日志级别，logrus对于日志级别的限制是高于这个级别的不会被显示，默认是InfoLevel，最高的级别是Trace
func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)                 // 设置在输出日志中添加文件名和方法信息
	logrus.SetFormatter(&logrus.JSONFormatter{}) //设置日志格式为json格式，默认为文本格式

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
