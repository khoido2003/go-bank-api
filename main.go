package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func seedAccount(store Storage, fname string, lname string, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)

	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {

		log.Fatal(err)
	}

	fmt.Println("new account: ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "anthony", "gg", "12345678")
}

func main() {

	godotenv.Load()

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {

		// run this ./go-bank-api --seed
		fmt.Println("seed database")
		// seed data
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
