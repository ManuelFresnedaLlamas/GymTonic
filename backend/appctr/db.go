package appctr

import (
	"log"
	"time"

	"go.uber.org/zap"

	dbx "github.com/go-ozzo/ozzo-dbx"

	_ "github.com/jackc/pgx/stdlib"
)

func DB() *dbx.DB {
	return &db
}

var db dbx.DB

func prepareDB() {
	str := cfg.GetString("db")
	log.Println(str)

	d, err := dbx.MustOpen("pgx", str)
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	if env == EnvDev {
		d.LogFunc = lg.Sugar().Debugf
	}

	go pingForStayConnected(d)

	db = *d

	lg.Debug("db is ok")
}

func pingForStayConnected(d *dbx.DB) {
	for {
		time.Sleep(5 * time.Minute)
		if err := d.DB().Ping(); err != nil {
			lg.Error("db ping error", zap.Error(err))
		}
	}
}
