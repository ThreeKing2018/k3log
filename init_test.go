package golog

import (
	"testing"
	"github.com/ThreeKing2018/k3log/conf"
	"fmt"
	"log"
)

func Test_logge(t *testing.T) {
	defer Sync()
	SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
		conf.WithProjectName("Zelog日志"),          //设置项目名称
		conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
		conf.WithLogLevel(conf.ErrorLevel),        //设置日志级别,默认debug
		conf.WithMaxAge(30),                      //日志保存天数,默认30天
		conf.WithMaxSize(512),                    //多少M进行分隔日志,默认100M
		//conf.WithStacktrace(conf.PanicLevel),                   //设置堆栈级别
		conf.WithIsStdOut(true)) //是否同时输出控制台
	Debug("debug日志", 1)
	Info("info日志", 2)
	Warn("warn日志", 3)
	Error("error日志", 4)
	Panic("panic", 5)
	Fatal("fatal", 6)
}

func BenchmarkInfo(t *testing.B) {
	SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
		conf.WithProjectName("Zelog日志"),          //设置项目名称
		conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
		conf.WithLogLevel(conf.InfoLevel),        //设置日志级别,默认debug
		conf.WithMaxAge(30),                      //日志保存天数,默认30天
		conf.WithMaxSize(512),                    //多少M进行分隔日志,默认100M
		conf.WithStacktrace(2),                   //设置堆栈级别
		conf.WithIsStdOut(true)) //是否同时输出控制台
	for i := 0; i < t.N; i++ {
		Info("测试日志", "打印结果", 100)
	}
}
func TestDebug(t *testing.T) {
	SetLogger(conf.WithIsStdOut(true),
		conf.WithLogType(conf.LogJsontype))
	Debug("ddd", 200, "aa", 2001, "bb")
}
func TestError(t *testing.T) {
	Error("error", 1)
	fmt.Println("继续执行")
}
func TestFatal(t *testing.T) {
	Fatal("fatal", 1)
	fmt.Println("程序中止")
	log.Fatalln()
}
func TestPanic(t *testing.T) {
	Panic("panic", 2)
	fmt.Println("程序中止")
}
func TestInfo2(t *testing.T) {
	Info("aa", 11)
	SetLogLevel(conf.InfoLevel)
	Info("info", 100)
	Warn("warn", 200)
	SetLogLevel(conf.ErrorLevel)
	Info("info-100", 300) //这个无法输出,因为上面设置日志级别为:error
	Error("err", 400)
}