package configs

import (
	"errors"
	"log"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("database.port", "5432")
}

func Load() (DBConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Arquivo de configuração não encontrado: %v", err)
		} else {
			log.Printf("Erro ao ler o arquivo de configuração: %v", err)
			return DBConfig{}, err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return cfg.DB, nil
}

func GetDB() (*DBConfig, error) {
	if cfg == nil || cfg.DB == (DBConfig{}) {
		return nil, errors.New("Configurações do banco de dados não inicializadas ou vazias")
	}
	return &cfg.DB, nil
}

func GetServerPort() string {
	if cfg == nil {
		return ""
	}
	return cfg.API.Port
}

func NewEcho() (*echo.Echo, error) {
	e := echo.New()
	return e, nil
}

func StartServer(e *echo.Echo) error {
	port := GetServerPort()
	if err := e.Start(":" + port); err != nil {
		return err
	}
	return nil
}
