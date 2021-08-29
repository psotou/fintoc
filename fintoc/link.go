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
	// Since Get() method returns a *NewLink, we need to add
	// the AccountM interface in this struct to allow for a
	// syntax like link.Account.Method()
	Account AccountM
}

// All the available Link methods
type LinkM interface {
	All() []Link
	Get(string) *NewLink
}

type LinkClient struct {
	*APIClient
}

func (l *LinkClient) All() []Link {
	var links []Link
	dataBytes, _ := l.GetReq(LinksAll)
	err := json.Unmarshal(dataBytes, &links)
	if err != nil {
		log.Fatal(err.Error())
	}
	return links
}

func (l *LinkClient) Get(linkdId string) *NewLink {
	var link Link
	url := fmt.Sprintf(LinkURL, linkdId)
	dataBytes, _ := l.GetReq(url)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}
	newL := &NewLink{client: l.APIClient, Link: link, linkToken: linkdId}
	// The following populates the AccountClient struct in order to have it
	// ready for the AccountM interface to use its methods
	newL.Account = &AccountClient{NewLink: newL}

	return newL
}
