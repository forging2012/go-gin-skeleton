package app

import (
	"github.com/inhere/go-gin-skeleton/model"
	"os"
)

// allowed app env name
const (
	PROD  = "prod"
	AUDIT = "audit"
	TEST  = "test"
	DEV   = "dev"
)

// for application
const (
	Timezone = "PRC"
	BaseDate = "2006-01-02 15:04:05"

	Timeout     = 10
	PageSize    = 20
	PageSizeStr = "20"
	MaxPageSize = 100
)

var (
	Env      = "dev"
	Name     = "github.com/inhere/go-gin-skeleton"
	Debug    bool
	Hostname string
	RootPath string
	GitInfo  model.GitInfoData

	HttpPort = 59440
)

// the app work dir path
var WorkDir, _ = os.Getwd()

// IsEnv current env name check
func IsEnv(env string) bool {
	return env == Env
}

// AbsPath always return abs path.
func AbsPath(path string) string {
	if string(path[0]) == "/" {
		return path
	}

	return path
}
