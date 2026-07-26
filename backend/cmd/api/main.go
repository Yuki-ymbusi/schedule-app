package main

import (
	"log"

	"schedule-api/internal/config"
	"schedule-api/internal/database"
	"schedule-api/internal/model"
	"schedule-api/internal/router"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 引数の型でテーブルを生成
	if err := db.AutoMigrate(&model.Schedule{}); err != nil {
		log.Fatal(err)
	}

	r := router.SetupRouter()
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}

}
