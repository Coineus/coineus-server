package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/coineus/coineus-server/internal/config"
	"github.com/coineus/coineus-server/internal/router"
	"github.com/coineus/coineus-server/internal/storage"
)

var (
	PORT = "8000"
)

func main() {
	dbCfg := os.Getenv("POSTGRE_DB_URI")

	retryTime, err := strconv.Atoi(os.Getenv("RETRY_COOLDOWN"))
	if err != nil || os.Getenv("RETRY_COOLDOWN") == "" {
		retryTime = 5
	}

	retryCount, err := strconv.Atoi(os.Getenv("RETRY_COUNT"))
	if err != nil || os.Getenv("RETRY_COUNT") == "" {
		retryCount = 5
	}

	// Waiting for other dependencies to be ready
	for i := 0; i < retryCount; i++ {
		if _, err := storage.CreatePool(dbCfg); err != nil {
			log.Println("Error connecting to postgres: ", err)
			time.Sleep(time.Second * time.Duration(retryTime))
		} else {
			break
		}
	}

	db, err := storage.CreatePool(dbCfg)
	if err != nil {
		log.Fatalf("db err : %v", err)
	}

	store := storage.New(db)

	r := router.New(store)
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	r.ServeHTTP(PORT)
}

func init() {
	// Check error in local
	_ = config.GetConfig()
}
