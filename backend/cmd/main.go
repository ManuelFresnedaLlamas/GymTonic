package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/ManuelFresnedaLlamas/GymTonic/fixtures"

	"github.com/ManuelFresnedaLlamas/GymTonic/appctr"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	go startProf()

	appctr.Start()

	appctr.Log().Info("environment is " + appctr.Env())
	if appctr.Env() == appctr.EnvDev {
		checkOrDoMigrations()
		checkOrDoFixtures()
	} else if appctr.Env() == appctr.EnvProd {
		checkOrDoProdMigrations()
	}

	appctr.UseMiddlewares()

	prepareIoC()

	http.Handle("/", appctr.Router())

	port := appctr.Cfg().GetString("port")
	appctr.Log().Info("starting server " + port)

	log.Fatalln(http.ListenAndServe(port, nil))
}

func checkOrDoMigrations() {
	if appctr.Cfg().GetString("migrations") == "true" {
		appctr.Log().Debug("starting migrations")

		migrate.SetTable("MIGRATIONS")

		ms := &migrate.FileMigrationSource{
			Dir: "migrations",
		}

		n, err := migrate.ExecMax(appctr.DB().DB(), "postgres", ms, migrate.Down, 0)
		if err != nil {
			appctr.Log().Fatal(err.Error())
		}
		appctr.Log().Debug(strconv.Itoa(n) + " migrations removed")

		n, err = migrate.Exec(appctr.DB().DB(), "postgres", ms, migrate.Up)
		if err != nil {
			appctr.Log().Fatal(err.Error())
		}
		appctr.Log().Debug(strconv.Itoa(n) + " migrations applied")
	}
}

func checkOrDoFixtures() {
	if appctr.Cfg().GetString("fixtures") == "true" {
		fixtures.Make()
	}
}
func checkOrDoProdMigrations() {
	if appctr.Cfg().GetString("migrations") != "true" {
		return
	}

	appctr.Log().Debug("starting prod migrations")

	folder := "/dst/migrations/prod"

	d, err := os.Open(folder)
	if err != nil {
		appctr.Log().Fatal(err.Error())
	}
	files, err := d.Readdir(-1)
	if err != nil {
		appctr.Log().Fatal(err.Error())
	}

	var n int

	for i := range files {
		if !strings.HasSuffix(files[i].Name(), ".sql") {
			continue
		}

		name := folder + "/" + files[i].Name()

		cmd := exec.Command("psql", "-U", "postgres", "-h", "gymº-db", "-d", "gym", "-w", "-a", "-q", "-f", name)
		if err := cmd.Run(); err != nil {
			appctr.Log().Fatal(err.Error())
		}
		n++
	}

	appctr.Log().Debug(strconv.Itoa(n) + " migrations applied")
}
func startProf() {
	r := http.NewServeMux()
	r.HandleFunc("/pprof", pprof.Index)
	r.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/pprof/profile", pprof.Profile)
	r.HandleFunc("/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/pprof/trace", pprof.Trace)
	log.Fatal(http.ListenAndServe(":7777", r))
}
