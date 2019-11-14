package main

import (
	"context"
	"log"

	"github.com/kancers/exsample-db/backend/sqlboiler/database"
	"github.com/kancers/exsample-db/backend/sqlboiler/models"
)

func main() {

	db, err := database.New("root:pass@tcp(127.0.0.1:3306)/demo?parseTime=true")
	if err != nil {
		log.Fatalf("failed to mysql open %+v", err)
	}

	// set global database
	// boil.SetDB(db)

	// select by id
	ctx := context.Background()
	user, err := models.Users().One(ctx, db)

	if err != nil {
		log.Fatalf("missing user: %+v", err)
	}

	log.Printf("got user: %+v", user)
}