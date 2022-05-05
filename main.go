package main

import (
	"strateegy/billing-service/database"
	"strateegy/billing-service/server"
)

func main() {
	s := server.NewServer()
	database.StartDB()
	s.Run()
}
