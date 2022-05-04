package services

import (
	"github.com/joho/godotenv"
)

func SetEnvVariables(filePath string) {
	godotenv.Load(filePath)
}
