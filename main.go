package main

import (
	"caturandi-labs/golang-starter/config"
	"caturandi-labs/golang-starter/ent"
	"caturandi-labs/golang-starter/ent/migrate"
	"caturandi-labs/golang-starter/handlers"
	"caturandi-labs/golang-starter/middleware"
	"caturandi-labs/golang-starter/routes"
	"caturandi-labs/golang-starter/utils"
	"context"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Configurations
	conf := config.New()

	// DB Connection
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Name))

	if err != nil {
		log.Fatalf("Failed Opening Connection to Database : %s", err)
	}

	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		utils.Fatalf("Migration failed : ", err)
	}

	// Set router
	app := fiber.New()
	middleware.SetMiddleware(app)


	handler := handlers.NewHandler(client, conf)

	routes.SetupApiV1(app, handler)

	port := "8000"

	addr  := flag.String("addr", port, "http service address")
	log.Fatal(app.Listen(":" + *addr))
}
