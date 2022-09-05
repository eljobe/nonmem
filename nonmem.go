package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk"
	"github.com/eljobe/nonmem/listmonk/models"
	"github.com/eljobe/nonmem/listmonk/subscribers"
	"github.com/eljobe/nonmem/zest"
)

func isImplicitNonMember(s subscribers.Subscriber, lm *listmonk.Listmonk) bool {
	cnId := lm.MustList(zest.ClubNews).Id
	mId := lm.MustList(zest.Members).Id
	nmId := lm.MustList(zest.NonMembers).Id
	return s.IsEnabled() && s.IsSubscribedTo(cnId) && !s.IsSubscribedTo(mId) && !s.IsSubscribedTo(nmId)
}

func main() {
	flag.Parse()

	lm := listmonk.NewListmonk()

	// Fetch the "Club News" list.
	cnList := lm.MustList(zest.ClubNews)
	cnSubs, err := subscribers.OfList(cnList.Id)
	if err != nil {
		log.Fatal(err)
	}

	// Add the Non-Members that need to be added.
	toAdd := []subscribers.Subscriber{}
	for _, cnSub := range cnSubs {
		if isImplicitNonMember(cnSub, lm) {
			toAdd = append(toAdd, cnSub)
		}
	}

	for _, ta := range toAdd {
		fmt.Println(ta)
	}

	if len(toAdd) > 0 {
		fmt.Println("Non-Members to add:", len(toAdd))
		nmList := lm.MustList(zest.NonMembers)
		err = subscribers.BulkSubscribe(toAdd, []models.ListId{nmList.Id})
		if err != nil {
			log.Fatal(err)
		}
	}
}
