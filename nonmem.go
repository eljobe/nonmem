package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk"
	"github.com/eljobe/nonmem/listmonk/lists"
)

func main() {
	flag.Parse()
	fmt.Println("Updating a Non-Member's List.")
	fmt.Println("Talking to Listmonk on ", listmonk.ApiUrl)

	// Fetch the "Club News" list.
	ls, err := lists.LookupLists()
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range ls.Names() {
		fmt.Println(name, "-", ls.Id(name))
	}

	// Fetch the "Members" list.
	// Fetch the "Non-Members" list.

	// Build a temorary list by subtracting "Members" from "Club News"
	// Remove anyone from "Non-Members" who isn't in the temporary list
	// Add anyone to "Non-Members" who is in the temporary list.
}
