package util

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	Debug  bool
	DbName string
	DbUri  string
}

var (
	flagValDebug      = flag.Bool("debug", envBool("DEBUG", false), "enable debug logging")
	flagValBucketName = flag.String("bucketname", envStr("BUCKET_NAME", ""), "bucket name")

	Config config
)

func InitConfig() {
	// flags

	flag.Parse()

	Config = config{
		Debug:  *flagValDebug,
		DbName: *flagValBucketName,
	}

	// logger

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
