package logger

import "github.com/sirupsen/logrus"

var _defaultLogger = logrus.New()

func Default() *logrus.Logger {
	return _defaultLogger
}
