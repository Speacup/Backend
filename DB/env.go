package DB

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Loadenv() {
	err := godotenv.Load(os.ExpandEnv("DB/.env"))
	if err != nil {
		log.Fatal(err)
	}

}
func env(name string) string {
	Loadenv()
	enn := os.Getenv(name)
	return enn
}
func envint(name string) int {
	star := env(name)
	i, err := strconv.Atoi(star)
	if err != nil {
		panic(err)
	}
	return i
}
