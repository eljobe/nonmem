package lists

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eljobe/nonmem/listmonk/models"
)

type data struct {
	Results []models.List `json:"results"`
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
	Total   int           `jsno:"total"`
}

type listData struct {
	Data data `json:"data"`
}

const url = "/lists?per_page=all"

type client interface {
	Get(string) (*http.Response, error)
}

func LookupLists(c client) ([]models.List, error) {
	res, err := c.Get(url)
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

	return resData.Data.Results, nil
}
