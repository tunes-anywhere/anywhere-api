package util

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	Debug      bool
	TableName  string
	BucketName string
	YTAPIKey   string
}

var (
	Config config
)

func InitConfig() {
	Config = config{
		Debug:      envBool("DEBUG", false),
		TableName:  envStr("TABLE_NAME", ""),
		BucketName: envStr("BUCKET_NAME", ""),
		YTAPIKey:   envStr("YT_API_KEY", ""),
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if Config.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:          os.Stdout,
			PartsExclude: []string{"timestamp"},
		})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func envStr(envKey string, defaultVal string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}

	return defaultVal
}

func envBool(envKey string, defaultVal bool) bool {
	if val := os.Getenv(envKey); val != "" {
		return val == "true"
	}

	return defaultVal
}
