package config

import (
	"WMB/authenticator"
	"WMB/manager"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	path           string
	name           string
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	ConfigToken    authenticator.Token
}

func (c *Config) readConfigFile() *viper.Viper {
	viper.SetConfigName(c.name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config File Not Found")
		} else {
			panic("Config File Error")
		}
	}
	return viper.GetViper()
}

func (c Config) Get(key string) string {
	return c.readConfigFile().GetString(key)
}

func New(path string, configFileName string) *Config {
	config := new(Config)
	config.name = configFileName
	config.path = path
	dbHost := config.Get("wmb.database.host")
	dbPort := config.Get("wmb.database.port")
	dbName := config.Get("wmb.database.name")
	dbUser := config.Get("wmb.database.user")
	dbPassword := config.Get("wmb.database.password")
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	infraManager := manager.NewInfra(dataSourceName)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager
	config.ConfigToken = newTokenConfig()
	//config.name = configFileName
	//config.path = path
	return config
}

func newTokenConfig() authenticator.Token {
	tokenConfig := authenticator.TokenConfig{
		AplicationName:      "Warung Makan Bahari",
		JwtSignatureKey:     "P@ssw0rd",
		JwtSignatureMethod:  jwt.SigningMethodHS256,
		AccessTokenDuration: 600 * time.Second,
	}
	return authenticator.NewToken(tokenConfig)
}
