package config

import "os"

var (
	TOKEN = os.Getenv("TOKEN")
	PORT  = os.Getenv("PORT")
)

const portConf string = "3000"

func Port() string {
	if PORT == "" {
		PORT = portConf
	}
	return PORT
}
