package main

import (
	"github.com/labstack/gommon/log"
	"github.com/theakshaygupta/go-authapi/dbo"
	"github.com/theakshaygupta/go-authapi/http_server"
)

func main() {
	db, err := dbo.Connect()
	if err != nil {
		log.Fatalf("Cannot connect to DB: \n%v", err)
	}
	defer db.Close()

	e := http_server.NewEchoServer(db)
	e.Logger.Fatal(e.Start(":5000"))
}
