package lists

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/eljobe/nonmem/listmonk"
)

type ListName string

type list struct {
	Id   int      `json:"id"`
	Name ListName `json:"name"`
}

type data struct {
	Results []list `json:"results"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Total   int    `jsno:"total"`
}

type listData struct {
	Data data `json:"data"`
}

type listMap map[ListName]int

type lists struct {
	listNames listMap
}

const url = "/lists?per_page=all"

func LookupLists() (*lists, error) {
	listsUrl := listmonk.ApiUrl + url

	lmClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, listsUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := lmClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resData := listData{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}

	listNames := listMap{}

	for _, list := range resData.Data.Results {
		listNames[list.Name] = list.Id
	}

	return &lists{listNames}, nil
}

func (l *lists) Names() []ListName {
	names := []ListName{}
	for name, _ := range l.listNames {
		names = append(names, name)
	}
	return names
}

func (l *lists) Id(name ListName) int {
	return l.listNames[name]
}