package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk"
	"github.com/eljobe/nonmem/listmonk/lists"
	"github.com/eljobe/nonmem/listmonk/subscribers"
)

func main() {
	flag.Parse()
	fmt.Println("Updating a Non-Member's List.")
	fmt.Println("Talking to Listmonk on ", listmonk.ApiUrl)

	// List the lists.
	ls, err := lists.LookupLists()
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the "Club News" list.
	clubNewsId := ls.Id("Club News")
	cnSubs, err := subscribers.OfList(clubNewsId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Club News", len(cnSubs), "subscribers")
	toAdd := []subscribers.Subscriber{}
	for _, cnSub := range cnSubs {
		if !cnSub.IsMember() && !cnSub.IsNonMember() {
			toAdd = append(toAdd, cnSub)
		}
	}
	fmt.Println("Non-Members to add:", len(toAdd))

	// Fetch the "Non-Members" list.
	nonMembersId := ls.Id("Non-Members")
	nmSubs, err := subscribers.OfList(nonMembersId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Non-Memebers", len(nmSubs), "subscribers")

	// Build a temorary list by subtracting "Members" from "Club News"
	// Remove anyone from "Non-Members" who isn't in the temporary list
	// Add anyone to "Non-Members" who is in the temporary list.
}
