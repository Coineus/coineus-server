package main

import (
	"log"
	"os"

	"github.com/coineus/coineus-server/internal/config"
	"github.com/coineus/coineus-server/internal/router"
	"github.com/coineus/coineus-server/internal/storage"
)

func main() {
	dbCfg := os.Getenv("POSTGRE_DB_URI")

	db := storage.CreatePool(dbCfg)

	store := storage.New(db)
	r := router.New(store)
	r.ServeHTTP("8000")
}

func init() {
	err := config.GetConfig()
	log.Fatal("config fail : ", err)
}
