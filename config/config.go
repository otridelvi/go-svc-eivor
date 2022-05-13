package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Apps     Apps           `json:"apps"`
	Database PostgresConfig `json:"database"`
}

type Apps struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	HttpPort int    `json:"httpPort"`
	Version  string `json:"version"`
}

type PostgresConfig struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	Name               string `json:"name"`
	Schema             string `json:"schema"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	MinIdleConnections int    `json:"minIdleConnections"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	DebugMode          bool   `json:"debugMode"`
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf("%s:%d", c.Apps.Address, c.Apps.HttpPort)
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ...")
	viper.SetConfigFile(path)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error ReadInConfig")
		panic(err)
	}

	conf := Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println("Error Unmarshal Config")
		panic(err)
	}

	configuration, _ := json.Marshal(conf)
	fmt.Println(string(configuration))
	return &conf

}
