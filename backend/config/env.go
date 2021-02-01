package config

import (
	"os"
)

type Env interface {
	WebServerPort() string
	MongoHost() string
	MongoDatabase() string
	AuthSecret() string

}

func NewEnv() Env {
	return &env{}
}

type env struct {
}

func (env) WebServerPort() string {
	return getString("SERVER_PORT", "8080")
}

func (env) MongoHost() string {
	return getString("MONGO_HOST", "mongodb://localhost:27017")
}

func (env) MongoDatabase() string {
	return getString("MONGO_DATABASE", "todo-list")
}

func (env) AuthSecret() string {
	return getString("AUTH_SECRET", "f0022cb73b9f478189e38bbe50e07a3b")
}


func getString(name string, defaultValue string) string {
	envValue := os.Getenv(name)

	if envValue == "" {

		return defaultValue
	}

	return envValue
}
