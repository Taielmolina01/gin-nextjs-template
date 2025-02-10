package application

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const (
	MESSAGE_ERROR_LOADING_ENV   = "Error loading the environment variables: %v"
	MESSAGE_SUCCESS_LOADING_ENV = "Success loading the environment variables"
)

type Configuration struct {
	Port         string
	DbDsn        string
	JwtAlgorithm string
	JwtSecretKey    string
}

var config *Configuration

func LoadConfig() *Configuration {
	envFile := filepath.Join(".", ".env")

	err := godotenv.Load(envFile)

	if err != nil {
		fmt.Errorf(MESSAGE_ERROR_LOADING_ENV, err)
		return nil
	}

	fmt.Println(MESSAGE_SUCCESS_LOADING_ENV)

	config = &Configuration{
		Port:         os.Getenv("PORT"),
		DbDsn:        os.Getenv("DB_DSN"),
		JwtAlgorithm: os.Getenv("JWT_ALGORITHM"),
		JwtSecretKey:    os.Getenv("JWT_SECRET"),
	}
	return config
}

func LoadConfigTest(port, dbDsn, jwtAlgorithm, jwtSecretKey string) *Configuration {
	config = &Configuration{
		Port: port,
		DbDsn: dbDsn,
		JwtAlgorithm: jwtAlgorithm,
		JwtSecretKey: jwtSecretKey,
	}
	return config
}

func GetConfiguration() *Configuration {
	return config
}