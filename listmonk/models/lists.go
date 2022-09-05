package models

import (
	"fmt"
	"time"
)

type ListName string

type ListId int

func (i ListId) String() string {
	return fmt.Sprintf("%d", i)
}

type ListType string

const (
	Public  ListType = "public"
	Private          = "private"
)

type OptInType string

const (
	Single OptInType = "single"
	Dobule           = "double"
)

type List struct {
	Id              ListId    `json:"id"`
	Created         time.Time `json:"created_at"`
	Updated         time.Time `json:"updated_at"`
	Name            ListName  `json:"name"`
	Type            ListType  `json:"type"`
	OptIn           OptInType `json:"optin"`
	Tags            []string  `json:"tags"`
	SubscriberCount int       `json:"subscriber_count"`
}
