package config

import (
	// "os"
	"golang/logging"
	"github.com/joho/godotenv"
)

// Init init env file
func Init() {
	err := godotenv.Load()
	if err != nil {
		logging.GetLog().Info("Can't load environtment")
		panic(err)
	}
}

// func XMLFile() {
// 	xml, err := os.Open("response.xml")
// }