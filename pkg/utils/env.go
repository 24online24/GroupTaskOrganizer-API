package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(name string) string {
	err := godotenv.Load(".env")
	HandleError(err)

	return os.Getenv(name)
}
