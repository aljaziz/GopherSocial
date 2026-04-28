package main

import (
	"log"

	"github.com/aljaziz/GopherSocial/internal/db"
	"github.com/aljaziz/GopherSocial/internal/env"
	"github.com/aljaziz/GopherSocial/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
