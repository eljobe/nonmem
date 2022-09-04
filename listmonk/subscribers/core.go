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
	Id                      int    `json:"id"`
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

type bulkSubscribe struct {
	SubscriberIds []int          `json:"ids"`
	Action        string         `json:"action"`
	ListIds       []lists.ListId `json:"target_list_ids"`
	Status        string         `json:"status"`
}

type respData struct {
	Message string `json:"message"`
}

const url = "/subscribers"

func (s *Subscriber) String() string {
	retVal, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return string(retVal)
}

func (s *Subscriber) IsEnabled() bool {
	return s.Status == "enabled"
}

func (s *Subscriber) IsSubscribedTo(id lists.ListId) bool {
	for _, l := range s.Lists {
		if l.Id == id {
			return l.SubStatus != "unsubscribed"
		}
	}
	return false
}

func BulkSubscribe(subs []Subscriber, listIds []lists.ListId) error {
	subUrl := listmonk.ApiUrl + url + "/lists"
	subIds := []int{}
	for _, s := range subs {
		subIds = append(subIds, s.Id)
	}
	content := bulkSubscribe{
		SubscriberIds: subIds,
		Action:        "add",
		ListIds:       listIds,
		Status:        "confirmed",
	}

	lmClient := http.Client{
		Timeout: time.Second * 10, //Timeout after 10 seconds
	}

	val, err := json.Marshal(content)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, subUrl, bytes.NewBuffer(val))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := lmClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resData := respData{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return err
	}
	return nil
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
	_, err = lmClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
