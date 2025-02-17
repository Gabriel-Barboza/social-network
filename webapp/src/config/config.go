package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = " "
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

func Carregar() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
