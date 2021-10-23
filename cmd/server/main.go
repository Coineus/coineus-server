package main

import (
	"log"
	"os"

	"github.com/coineus/coineus-server/internal/config"
	"github.com/coineus/coineus-server/internal/router"
	"github.com/coineus/coineus-server/internal/storage"
)

var (
	PORT = "8000"
)

func main() {
	dbCfg := os.Getenv("POSTGRE_DB_URI")

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
	err := config.GetConfig()
	if err != nil {
		log.Fatal("config fail : ", err)
	}
}
