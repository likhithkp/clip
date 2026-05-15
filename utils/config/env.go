package config

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_const "github.com/likhithkp/clip/utils/const"
)

type Env struct {
	Port          string
	MongodbUri    string
	DbName        string
	DeploymentEnv string
	JwtSecretKey  string
}

func GetEnv() (*Env, error) {
	deploymentEnv := strings.TrimSpace(os.Getenv("DEPLOYMENT_ENV"))
	if deploymentEnv != string(_const.Deployment_Production) {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors.New("HTTP_PORT not provided")
	}

	mongodbUri := os.Getenv("MONGODB_URI")
	if len(mongodbUri) == 0 {
		return nil, errors.New("MONGODB_URI not provided")
	}

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		return nil, errors.New("DB_NAME not provided")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET")
	if len(jwtSecretKey) == 0 {
		return nil, errors.New("JWT_SECRET not provided")
	}

	return &Env{
		Port:          port,
		MongodbUri:    mongodbUri,
		DbName:        dbName,
		DeploymentEnv: deploymentEnv,
		JwtSecretKey:  jwtSecretKey,
	}, nil
}
