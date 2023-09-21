package logger

import (
	"github.com/sirupsen/logrus"
)

var _defaultLogger = logrus.New()

func Default() *logrus.Logger {
	return _defaultLogger
}

func Debugf(format string, args ...interface{})   { _defaultLogger.Debugf(format, args...) }
func Debugln(args ...interface{})                 { _defaultLogger.Debugln(args...) }
func Errorf(format string, args ...interface{})   { _defaultLogger.Errorf(format, args...) }
func Errorln(args ...interface{})                 { _defaultLogger.Errorln(args...) }
func Fatalf(format string, args ...interface{})   { _defaultLogger.Fatalf(format, args...) }
func Fatalln(args ...interface{})                 { _defaultLogger.Fatalln(args...) }
func Infof(format string, args ...interface{})    { _defaultLogger.Infof(format, args...) }
func Infoln(args ...interface{})                  { _defaultLogger.Infoln(args...) }
func Panicf(format string, args ...interface{})   { _defaultLogger.Panicf(format, args...) }
func Panicln(args ...interface{})                 { _defaultLogger.Panicln(args...) }
func Printf(format string, args ...interface{})   { _defaultLogger.Printf(format, args...) }
func Println(args ...interface{})                 { _defaultLogger.Println(args...) }
func Warningf(format string, args ...interface{}) { _defaultLogger.Warningf(format, args...) }
func Warningln(args ...interface{})               { _defaultLogger.Warningln(args...) }
