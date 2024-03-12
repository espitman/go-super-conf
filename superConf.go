package superConf

import (
	"os"
	"strings"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

func init() {
	env := "dev"
	if os.Getenv("APP_ENV") != "" {
		env = strings.ToLower(os.Getenv("APP_ENV"))
	}
	config.WithOptions(config.ParseEnv)
	config.AddDriver(json.Driver)
	err := config.LoadFiles("config/default.json")
	if err != nil {
		panic(err)
	}
	_ = config.LoadFiles("config/" + env + ".json")
	_ = config.LoadFiles("config/local.json")
}

func Get(key string) string {
	envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
	envVar := os.Getenv(envKey)
	if envVar != "" {
		return envVar
	}
	value := config.String(key)
	return value
}
