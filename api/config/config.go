package config

import (
	"encoding/json"
	"os"
	"time"
)

type configuration struct {
	app      `json:"app"`
	server   `json:"server"`
	database `json:"database"`
	apiKey   `json:"apiKey"`
}

type app struct {
	Env                        string `json:"env"`
	JWTSecret                  string `json:"jwtSecret"`
	TokenExpireDuration        int    `json:"tokenExpireDuration"`
	RefreshTokenExpireDuration int    `json:"refreshTokenExpireDuration"`
}

type server struct {
	RunMode      string        `json:"runMode"`
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
}

type database struct {
	MySqlDSN string `json:"mySqlDSN"`
}

type apiKey struct {
	ApiIp string `json:"apiIp"`
}

var (
	AppConfig      app
	ServerConfig   server
	DatabaseConfig database
	ApiKeyConfig   apiKey
)

func Setup(filename string) error {
	var config configuration
	configFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return err
	}

	AppConfig = config.app
	ServerConfig = config.server
	DatabaseConfig = config.database
	ApiKeyConfig = config.apiKey

	return nil
}
