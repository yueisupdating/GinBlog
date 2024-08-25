package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Log() gin.HandlerFunc {
	filePath := "log/"

	logWriter, _ := rotatelogs.New(
		filePath+"%Y-%m-%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName("latest_log.log"),
	)

	// logger将捕捉以下情况下的log
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}
	// 时间戳为固定写法
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.AddHook(Hook)

	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		// 洋葱模型,request入栈,response出栈
		spendTime := fmt.Sprintf("%d ms", time.Since(startTime).Milliseconds()) // 记录服务器处理操作用时
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unkown"
		}
		code := ctx.Writer.Status()
		logs := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"SpendTime": spendTime,
			"Method":    ctx.Request.Method,
			"Ip":        ctx.ClientIP(),
			"status":    code,
			"DataSize":  ctx.Writer.Size(),
			"Agent":     ctx.Request.UserAgent(),
			"Path":      ctx.Request.RequestURI,
		})
		// 判断package自带的error
		if ctx.Errors != nil {
			logs.Error(ctx.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		// 根据返回码判断,包括错误警告以及正常

		if code >= 500 {
			logs.Error()
		} else if code >= 400 {
			logs.Warn()
		} else {
			logs.Info()
		}
	}
}
