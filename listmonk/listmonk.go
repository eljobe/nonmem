package listmonk

import (
	"fmt"
	"log"

	"github.com/eljobe/nonmem/listmonk/lists"
	"github.com/eljobe/nonmem/listmonk/models"
)

type listMap map[models.ListName]models.List

type Listmonk struct {
	c    *client
	lmap listMap
}

func NewListmonkAt(url string) *Listmonk {
	return &Listmonk{
		c:    newClientAt(url),
		lmap: make(listMap),
	}
}

func NewListmonk() *Listmonk {
	return &Listmonk{
		c:    newClient(),
		lmap: make(listMap),
	}
}

func (l *Listmonk) loadLists() error {
	ls, err := lists.LookupLists(l.c)
	if err != nil {
		return err
	}
	for _, list := range ls {
		l.lmap[list.Name] = list
	}
	return nil
}

func (l *Listmonk) List(name models.ListName) (*models.List, error) {
	if val, ok := l.lmap[name]; ok {
		return &val, nil
	}
	err := l.loadLists()
	if err != nil {
		return nil, err
	}
	if val, ok := l.lmap[name]; ok {
		return &val, nil
	} else {
		return nil, fmt.Errorf("%s Not Found", name)
	}
}

func (l *Listmonk) MustList(name models.ListName) *models.List {
	found, err := l.List(name)
	if err != nil {
		log.Fatal(err)
	}
	return found
}
