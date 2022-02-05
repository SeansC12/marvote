package logging

import (
	"go.uber.org/zap"
)

var zapLog *zap.Logger

func init() {
	var err error
	config := zap.NewDevelopmentConfig()
	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(args ...interface{}) {
	zapLog.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLog.Sugar().Infof(template, args...)
}

func Debug(args ...interface{}) {
	zapLog.Sugar().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapLog.Sugar().Debugf(template, args...)
}

func Error(args ...interface{}) {
	zapLog.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLog.Sugar().Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	zapLog.Sugar().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zapLog.Sugar().Fatalf(template, args...)
}
