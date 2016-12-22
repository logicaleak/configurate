package configurate

import (
	"os"
	"encoding/json"
	"flag"
	"strings"
	log "github.com/Sirupsen/logrus"
)

var profile *string
var config Config

type Config struct {
	configMap map[string]interface{}
}

func getConfig(configKey string) interface{} {
	parts := strings.Split(configKey, ".")
	var current interface{} = config.configMap[parts[0]]
	for _, part := range parts[1:] {
		currentMap := current.(map[string]interface{})
		current = currentMap[part]
	}
	return current
}

func GetString(configKey string) string {
	return getConfig(configKey).(string)
}

func GetFloat(configKey string) float64 {
	return getConfig(configKey).(float64)
}

func GetInt(configKey string) int {
	return int(GetFloat(configKey))
}


func loadConfig(mustLoad bool) error {
	if profile == nil {
		profile = flag.String("profile", "", "Give the profile to switch the config")
	}

	log.WithFields(log.Fields{
		"profile" : *profile,
	}).Info("Loading config with profile")

	configMap := make(map[string]interface{})

	var file *os.File
	var err error
	if *profile == "" {
		log.Debug("Profile was empty")
		file, err = os.Open("config.json")
	} else {
		file, err = os.Open("config-" + *profile + ".json")
	}

	if err != nil {
		if mustLoad {
			panic(err)
		} else {
			return err
		}
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configMap)

	if err != nil {
		if mustLoad {
			panic(err)
		} else {
			return err
		}
	}

	log.Info("Successfully loaded config")
	config = Config{configMap: configMap}

	return nil
}

func Load() error {
	return loadConfig(false)
}

func MustLoad() {
	loadConfig(true)
}
