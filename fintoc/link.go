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

func (a *APIClient) LinkOne(linkdId string) *NewLink {
	var link Link
	url := fmt.Sprintf(LinkURL, linkdId)
	dataBytes, _ := a.GetReq(url)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}
	newL := &NewLink{client: a, Link: link, linkToken: linkdId}
	return newL
}

type NewLink struct {
	client *APIClient
	// We make this field anonymous in order to directly get
	// the attributes of the Link object
	Link
	linkToken string
}
