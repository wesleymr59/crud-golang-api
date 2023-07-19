package main

import (
	"crud-golang-api/database"
	"crud-golang-api/routers"
)

func main() {
	database.ConnDataBase()
	routers.HandleRequests()
}
