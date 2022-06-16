package appctr

import (
	"log"

	"go.uber.org/zap"
)

func Log() *zap.Logger {
	return &lg
}

var lg zap.Logger

func prepareLog() {
	var err error
	var l *zap.Logger

	if env == EnvProd {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("logger can't be instanced")
	}

	lg = *l
}
