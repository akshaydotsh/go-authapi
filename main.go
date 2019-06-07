package main

import (
	"github.com/joho/godotenv"
	"github.com/theakshaygupta/go-authapi/config"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/http_server"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("File .env not found, reading configuration from ENV")
	}
	config.SetConfig()
}

func main() {
	db, err := dbo.Connect()
	if err != nil {
		log.Fatalf("FatalError:Cannot connect to DB: \n%v", err)
	}
	defer db.Close()

	e := http_server.NewEchoServer(db)
	log.Fatal(e.Start(":" + config.Config.Port))
}
