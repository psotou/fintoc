package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

func (a *APIClient) LinkAll() []Link {
	var links []Link
	dataBytes, _ := a.GetReq(LinksAll)
	err := json.Unmarshal(dataBytes, &links)
	if err != nil {
		log.Fatal(err.Error())
	}
	return links
}

func (a *APIClient) LinkOne(linkdId string) *newLink {
	var link Link
	url := fmt.Sprintf(LinkURL, linkdId)
	dataBytes, _ := a.GetReq(url)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}
	newL := &newLink{client: a, link: link, linkToken: linkdId}
	return newL
}

type newLink struct {
	client    *APIClient
	link      Link
	linkToken string
}
