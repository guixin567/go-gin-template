package log

import "go.uber.org/zap"

var ALog *zap.Logger

func Init() (err error) {
	ALog, err = zap.NewProduction()
	return
}
