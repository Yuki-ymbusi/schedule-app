package main

import (
	"log"

	"schedule-api/internal/router"
)

func main() {
	r := router.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
