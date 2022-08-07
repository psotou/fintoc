package fintoc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

type Link struct {
	Id          string      `json:"id"`
	Object      string      `json:"object"`
	UserName    string      `json:"username"`
	LinkToken   string      `json:"link_token,omitempty"`
	Mode        string      `json:"mode"`
	Active      bool        `json:"active"`
	Status      string      `json:"status"`
	HolderType  string      `json:"holder_type"`
	CreatedAt   time.Time   `json:"created_at"`
	Institution Institucion `json:"institution"`
	Accounts    []Account   `json:"accounts"`
}

type Institucion struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (c *client) GetLinks() []Link {
	var links []Link
	response, err := c.get(LinksAll)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response, &links)

	return links
}

type linkClient struct {
	client
	Link
}

// Links are funny... to retrieve one you don't pass a link ID, no, no, no:
// you pass a link_token instead.
func (c *client) GetLinkByID(linkToken string) linkClient {
	var link Link
	response, err := c.get(fmt.Sprintf(Links, linkToken))
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response, &link)

	lc := client{secretKey: c.secretKey, linkToken: linkToken}
	return linkClient{lc, link}
}

func (c *client) DeleteLinkByID(linkID string) (int, error) {
	return c.delete(fmt.Sprintf(Links, linkID))
}

func (lc *linkClient) ToggleLinkActiveStatus() linkClient {
	var link Link
	var response []byte
	var payload io.Reader
	var err error
	payload = strings.NewReader(`{"active":true}`)
	response, err = lc.patch(fmt.Sprintf(Links, lc.linkToken), payload)
	if isLinkActive(*lc) {
		payload = strings.NewReader(`{"active":false}`)
		response, err = lc.patch(fmt.Sprintf(Links, lc.linkToken), payload)
	}
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response, &link)

	return linkClient{lc.client, link}
}

func isLinkActive(link linkClient) bool {
	return link.Link.Status == "active"
}
