package main

import (
	"log"

	"github.com/LunaTearR/Ledgar/database"
	"github.com/gofiber/fiber/v3"
)

func main() {

	dbUser := "admin"
	dbName := "ledgar"
	dbPassword := "password"
	dbHost := "postgres"

	db, err := database.Connect(dbUser, dbName, dbPassword, dbHost)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.InitTables(db)
	if err != nil {
		log.Fatalln("Error InitTables : ", err)
	}

	app := fiber.New()

	app.Get("/healthcheck", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})
}
