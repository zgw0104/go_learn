package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

// 设置logger编码格式 json console...
func getEncoder() zapcore.Encoder {

	zapconfig := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewConsoleEncoder(zapconfig)
}

// 设置logger输出文件，以及写入方式
//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.OpenFile(
//		"./zap.log",
//		os.O_CREATE|os.O_APPEND|os.O_RDWR,
//		0744,
//	)
//	return zapcore.AddSync(file)
//}

func getLoggerWriter() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./zap.log",
		MaxSize:    0,
		MaxAge:     30, //最大备份天数
		MaxBackups: 5,  //最大备份数量
		LocalTime:  false,
		Compress:   false, //是否压缩
	}

	return zapcore.AddSync(lumberjackLogger)
}

func innitLogger() {
	//logger, _ = zap.NewProduction()
	//logger, _ = zap.NewDevelopment()
	//sugar = logger.Sugar()

	encoder := getEncoder()
	writeSyncer := getLoggerWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger = zap.New(core, zap.AddCaller()) //zao.AddCaller() -->  记录函数调用信息
	sugar = logger.Sugar()

}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugar.Error("error fetching url", zap.String("url", url), zap.Error(err))
	} else {
		sugar.Info("success", zap.String("status code", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}

//	func main() {
//		innitLogger()
//		defer logger.Sync()
//
//		simpleHttpGet("http://www.baidu.com")
//		simpleHttpGet("http://www.google.com")
//
// }

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func main() {
	//r := gin.Default()
	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true))
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	r.Run()
}
