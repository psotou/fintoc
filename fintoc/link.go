package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

type LinkM interface {
	All() []Link
	Get(string) *newLink
}

// type LinkClient struct {
// 	client *APIClient
// }

// func (l *LinkClient) All() []Link {
// 	var links []Link
// 	dataBytes, _ := l.client.GetReq(LinksAll)
// 	err := json.Unmarshal(dataBytes, &links)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return links
// }

func (a *APIClient) LinkAll() []Link {
	var links []Link
	dataBytes, _ := a.GetReq(LinksAll)
	err := json.Unmarshal(dataBytes, &links)
	if err != nil {
		log.Fatal(err.Error())
	}
	return links
}

// func (l *LinkClient) Get(linkdId string) *newLink {
// 	var link Link
// 	url := fmt.Sprintf(LinkUrl, linkdId)
// 	dataBytes, _ := l.client.GetReq(url)
// 	err := json.Unmarshal(dataBytes, &link)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	newL := &newLink{client: l.client, link: link}
// 	return newL
// }

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

// func (l *LinkClient) Get(linkdId string) LinkClient {
// 	var link Link
// 	url := fmt.Sprintf(LinkUrl, linkdId)
// 	dataBytes, _ := l.ApiClient.GetReq(url)
// 	err := json.Unmarshal(dataBytes, &link)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	acc := LinkClient{ApiClient: l.ApiClient, Link: link}
// 	return acc
// }

// func (l *LinkClient) AccAll() *AqqClient {
// 	var account []Account
// 	url := fmt.Sprintf(AccountsAll+LinkToken, l.ApiClient.LinkToken)
// 	dataBytes, _ := l.ApiClient.GetReq(url)
// 	err := json.Unmarshal(dataBytes, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	acc := &AqqClient{client: l.ApiClient, Acc: account}
// 	return acc
// }

// type AqqClient struct {
// 	client LinkClient
// 	Lin    Link
// }
