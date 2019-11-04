package main

import (
	"log"

	"github.com/kancers/exsample-db/backend/xo/database"
	"github.com/kancers/exsample-db/backend/xo/models"
)

func main() {
	db, err := database.New("mysql://root:pass@127.0.0.1/demo?parseTime=true")
	if err != nil {
		log.Fatalf("failed to mysql open: %+v", err)
	}

	// select by id
	user, err := models.UserByID(db, 1)
	if err != nil {
		log.Fatalf("missing user: %+v", err)
	}

	log.Printf("got user: %+v", user)
}
