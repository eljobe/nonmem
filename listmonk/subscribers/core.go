package subscribers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/eljobe/nonmem/listmonk"
	"github.com/eljobe/nonmem/listmonk/lists"
)

type List struct {
	SubStatus string         `json:"subscription_status"`
	Id        lists.ListId   `json:"id"`
	Name      lists.ListName `json:"name"`
}

type Subscriber struct {
	Email                   string `json:"email"`
	Name                    string `json:"name"`
	Status                  string `json:"status"` // enabled, disabled, or blocklisted
	Lists                   []List `json:"lists"`
	PreconfirmSubscriptions bool   `json:"preconfirm_subscriptions"`
}

type data struct {
	Results []Subscriber `json:"results"`
	Page    int          `json:"page"`
	PerPage int          `json:"per_page"`
	Total   int          `json:"total"`
}

type subData struct {
	Data data `json:"data"`
}

const url = "/subscribers"
const deleteUrl = "http://127.0.0.1:9000/api/subscribers/query/delete"
const enabledQuery = "\"query=subscribers.status = 'enabled'\""

func (s *Subscriber) String() string {
	retVal, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(retVal)
}

func (s *Subscriber) IsMember() bool {
	for _, l := range s.Lists {
		if l.Name == "Members" {
			return true
		}
	}
	return false
}

func (s *Subscriber) IsNonMember() bool {
	for _, l := range s.Lists {
		if l.Name == "Non-Members" {
			return true
		}
	}
	return false
}

func OfList(id lists.ListId) ([]Subscriber, error) {
	listUrl := listmonk.ApiUrl + url + "?list_id=" + id.String() + "&per_page=all"
	lmClient := http.Client{
		Timeout: time.Second * 10, //Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, listUrl, nil)
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

	resData := subData{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return nil, err
	}

	return resData.Data.Results, nil
}

func (s *Subscriber) Save() error {
	subUrl := listmonk.ApiUrl + url
	lmClient := http.Client{
		Timeout: time.Second * 10, //Timeout after 10 seconds
	}

	val, err := json.Marshal(s)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, subUrl, bytes.NewBuffer(val))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := lmClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
