package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(name string) string { // obține o variabilă de mediu
	err := godotenv.Load(".env") // încarcă variabilele de mediu din fișierul .env
	HandleError(err)

	return os.Getenv(name)
}
