package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk/lists"
	"github.com/eljobe/nonmem/listmonk/subscribers"
)

func main() {
	flag.Parse()

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
	toAdd := []subscribers.Subscriber{}
	for _, cnSub := range cnSubs {
		if cnSub.IsSubscribedTo(clubNewsId) && cnSub.IsEnabled() && !cnSub.IsMember() && !cnSub.IsNonMember() {
			toAdd = append(toAdd, cnSub)
		}
	}

	for _, ta := range toAdd {
		fmt.Println(ta)
	}

	if len(toAdd) > 0 {
		fmt.Println("Non-Members to add:", len(toAdd))
		err = subscribers.BulkSubscribe(toAdd, []lists.ListId{ls.Id("Non-Members")})
		if err != nil {
			log.Fatal(err)
		}
	}

	// Fetch the "Non-Members" list.
	// nonMembersId := ls.Id("Non-Members")
	// nmSubs, err := subscribers.OfList(nonMembersId)
	// if err != nil {
	// log.Fatal(err)
	// }

	// Build a temorary list by subtracting "Members" from "Club News"
	// Remove anyone from "Non-Members" who isn't in the temporary list
}
