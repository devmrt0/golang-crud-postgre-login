package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBURL     = ""
	DBDRIVER  = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env: %v", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		fmt.Println(err)
		PORT = 7000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")

	//DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local&", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	DBURL = "host=localhost user=postgres password=artelnk dbname=CloudRestDb port=5432 sslmode=disable"
}
