package middleware

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	log *logrus.Logger
}

//func NewLogger() *Logger {
//	logger := logrus.New()
//	logger.SetFormatter(&logrus.JSONFormatter{})
//	logger.SetLevel(logrus.InfoLevel)
//	return &Logger{log: logger}
//}

func Init(method func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Логирование началось")
		logger := logrus.New()
		start := time.Now()
		fields := logrus.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"ip":         r.RemoteAddr,
			"user_agent": r.UserAgent(),
		}
		file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.Errorf("Ошибка при открытии файла: %v", err)
		}
		defer file.Close()
		logger.SetOutput(file)
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.Info("Это сообщение будет только в файле")
		ctx := logrus.WithFields(fields)
		ctx.Info("Received request")
		method(w, r)

		latency := time.Since(start)
		ctx.WithFields(logrus.Fields{
			"latency": latency,
		}).Info("Request completed")
		fmt.Println("Логирование закончилось")
	}
}
