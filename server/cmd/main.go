package main

import (
	"github.com/GO-CHAT/server/db"
	"github.com/GO-CHAT/server/internal/user"
	"github.com/GO-CHAT/server/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8888")
}
