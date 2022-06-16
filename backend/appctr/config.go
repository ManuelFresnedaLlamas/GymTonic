package appctr

import (
	"github.com/spf13/viper"
	"os"
)

func Cfg() *viper.Viper {
	return &cfg
}

func Env() string {
	return env
}

func Domain() string {
	return domain
}

const (
	envPrefix = "GYM"

	envVar    = "env"
	domainVar = "domain"

	envUploadDir = "/uploads"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

var cfg viper.Viper
var env string
var domain string

func prepareCfg() {
	cfg = *viper.New()
	cfg.SetEnvPrefix(envPrefix)
	cfg.AutomaticEnv()

	env = cfg.GetString(envVar)
	domain = cfg.GetString(domainVar)
}

func prepareUpload() bool {
	path, _ := os.Getwd()

	if _, err := os.Stat(path + envUploadDir); os.IsNotExist(err) {
		errDir := os.Mkdir(path+envUploadDir, 0755)
		if errDir != nil {
			return false
		}
	}

	return true
}
