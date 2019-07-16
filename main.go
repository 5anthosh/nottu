package main

import (
	"nottu/db"
	"nottu/server"
)

func main() {
	db.ReinitDB()
	server.Run()
}
