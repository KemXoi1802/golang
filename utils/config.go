package utils

import (
	"github.com/joho/godotenv"
)

// Init init env file
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		GetLog().Info("Can't load environtment")
		panic(err)
	}
}
