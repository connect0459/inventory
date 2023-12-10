// config/config.go
package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config はconfig.jsonの構造を表します
type Config struct {
	DevMode bool `json:"DevMode"`
	DSN     struct {
		Dev      string `json:"Dev"`
		Production string `json:"Production"`
	} `json:"DSN"`
}

var AppConfig Config

func init() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file")
	}

	err = json.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatal("Error unmarshalling config file")
	}
}