# k3log

#### 介绍
> 取名Three King Log
- 由uber zap 日志扩展而来
- 实现分隔,异步,动态级别打印,json/txt
- 以key-value形式打印日志,适合项目里使用

#### 日志级别
- Debug 调度使用, 程序继续运行
- Info 提示使用, 程序继续运行
- Warn 警告使用, 程序继续运行
- Error 错误使用, 程序继续运行
- Panic 恐慌的,退出函数,不会退出应用唾弃,会执行defer
- Fatal 致命的,退出应用程序,不会执行defer,因为底层多一个os.Exit

#### 设置参数
- WithFilename    string //日志保存路径
- WithLogLevel    Level  //日志记录级别
- WithMaxSize     int    //日志分割的尺寸 MB
- WithMaxAge      int    //分割日志保存的时间 day
- WithStacktrace  Level  //记录堆栈的级别
- WithIsStdOut    bool   //是否标准输出console输出
- WithProjectName string //项目名称
- WithLogType     string //日志类型,普通 或 json

> 使用方法

- 简单使用

```golang
Debug("debug日志", 1)
Info("info日志", 2)
Warn("warn日志", 3)
Error("error日志", 4)
Panic("panic", 5)
Fatal("fatal", 6)
```

- 自测使用
```golang
SetLogger(conf.WithIsStdOut(true),
		conf.WithLogType(conf.LogJsontype))
	Debug("self test", 100)
```

- 线上测试使用
```golang
SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
    conf.WithProjectName("Zelog日志"),          //设置项目名称
    conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
Debug("debug日志", 1)
Info("info日志", 2)
Warn("warn日志", 3)
Error("error日志", 4)
Panic("panic", 5)
Fatal("fatal", 6)
```

- 生产使用

```golang
SetLogger(conf.WithLogType(conf.LogJsontype), //打印json格式
		conf.WithProjectName("Zelog日志"),          //设置项目名称
		conf.WithFilename("log.txt"),             //设置输出文件名,或输出的路径
		conf.WithLogLevel(conf.InfoLevel),        //设置日志级别,默认debug
		conf.WithMaxAge(30),                      //日志保存天数,默认30天
		conf.WithMaxSize(512),                    //多少M进行分隔日志,默认100M
		conf.WithIsStdOut(false)) //是否同时输出控制台
	defer Sync()
	Debug("debug日志", 1)
	Info("info日志", 2)
	Warn("warn日志", 3)
	Error("error日志", 4)
	Panic("panic", 5)
	Fatal("fatal", 6)
```

#### 动态改变日志的打印级别
```golang
Info("aa", 11)
SetLogLevel(conf.InfoLevel)
Info("info", 100)
Warn("warn", 200)
SetLogLevel(conf.ErrorLevel)
Info("info-100", 300) //这个无法输出,因为上面设置日志级别为:error
Error("err", 400)
```