package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk/lists"
	"github.com/eljobe/nonmem/listmonk/subscribers"
	"github.com/eljobe/nonmem/zest"
)

func isImplicitNonMember(s subscribers.Subscriber, ls *lists.Lists) bool {
	cnId := ls.Id(zest.ClubNews)
	mId := ls.Id(zest.Members)
	nmId := ls.Id(zest.NonMembers)
	return s.IsEnabled() && s.IsSubscribedTo(cnId) && !s.IsSubscribedTo(mId) && !s.IsSubscribedTo(nmId)
}

func main() {
	flag.Parse()

	// List the lists.
	ls, err := lists.LookupLists()
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the "Club News" list.
	clubNewsId := ls.Id(zest.ClubNews)
	cnSubs, err := subscribers.OfList(clubNewsId)
	if err != nil {
		log.Fatal(err)
	}

	// Add the Non-Members that need to be added.
	toAdd := []subscribers.Subscriber{}
	for _, cnSub := range cnSubs {
		if isImplicitNonMember(cnSub, ls) {
			toAdd = append(toAdd, cnSub)
		}
	}

	for _, ta := range toAdd {
		fmt.Println(ta)
	}

	if len(toAdd) > 0 {
		fmt.Println("Non-Members to add:", len(toAdd))
		err = subscribers.BulkSubscribe(toAdd, []lists.ListId{ls.Id(zest.NonMembers)})
		if err != nil {
			log.Fatal(err)
		}
	}
}
