package golog

import (
	"github.com/ThreeKing2018/k3log/conf"
	"github.com/ThreeKing2018/k3log/plugins/zaplog"
)

//默认
var l Loger = zaplog.New()

//设置
func SetLogger(opts ...conf.Option) {
	l = zaplog.New(opts...)
}

//目前只有zap生效
func SetLogLevel(level conf.Level) {
	l.SetLogLevel(level)
}

//目前只有zap生效
func Sync() {
	l.Sync()
}

//key value
func Debug(keysAndValues ...interface{}) {
	l.Debug(keysAndValues...)
}

func Info(keysAndValues ...interface{}) {
	l.Info(keysAndValues...)
}

func Warn(keysAndValues ...interface{}) {
	l.Warn(keysAndValues...)
}

func Error(keysAndValues ...interface{}) {
	l.Error(keysAndValues...)
}

func Panic(keysAndValues ...interface{}) {
	l.Panic(keysAndValues...)
}

func Fatal(keysAndValues ...interface{}) {
	l.Fatal(keysAndValues...)
}
