package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Link struct {
	Id          string      `json:"id"`
	Object      string      `json:"object"`
	UserName    string      `json:"username"`
	LinkToken   string      `json:"link_token"`
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
	Update(string, bool) *NewLink
	Delete(linkId string)
}

type LinkClient struct {
	*APIClient
}

func (l *LinkClient) All() []Link {
	var links []Link
	dataBytes, _ := l.getReq(LinksAll)
	err := json.Unmarshal(dataBytes, &links)
	if err != nil {
		log.Fatal(err.Error())
	}
	return links
}

func (l *LinkClient) Get(linkToken string) *NewLink {
	var link Link
	url := fmt.Sprintf(Links, linkToken)
	dataBytes, _ := l.getReq(url)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}
	newL := &NewLink{
		client:    l.APIClient,
		Link:      link,
		linkToken: linkToken,
	}
	// The following populates the AccountClient struct in order to have it
	// ready for the AccountM interface to use its methods
	newL.Account = &AccountClient{NewLink: newL}

	return newL
}

// The Update method that acts upon the object generated
// by the Get method of the Link interface.
// Activates o deactivates the status of a link.
func (n *NewLink) Update(active bool) *NewLink {
	var link Link
	var strPayload string
	url := fmt.Sprintf(Links, n.linkToken)

	if active {
		strPayload = "{\"active\":true}"
	} else {
		strPayload = "{\"active\":false}"
	}
	payload := strings.NewReader(strPayload)
	dataBytes, _ := n.client.updateReq(url, payload)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}

	newL := &NewLink{
		client:    n.client,
		Link:      link,
		linkToken: n.linkToken,
	}
	// The following populates the AccountClient struct in order to have it
	// ready for the AccountM interface to use its methods
	newL.Account = &AccountClient{NewLink: newL}

	return newL
}

// The Update method that will act upon the *LinkClient object.
// Activates o deactivates the status of a link.
func (l *LinkClient) Update(linkToken string, active bool) *NewLink {
	var (
		link       Link
		newL       *NewLink
		strPayload string
	)
	url := fmt.Sprintf(Links, linkToken)

	if active {
		strPayload = "{\"active\":true}"
	} else {
		strPayload = "{\"active\":false}"
	}
	payload := strings.NewReader(strPayload)
	dataBytes, _ := l.updateReq(url, payload)
	err := json.Unmarshal(dataBytes, &link)
	if err != nil {
		log.Fatal(err.Error())
	}

	newL = &NewLink{
		client:    l.APIClient,
		Link:      link,
		linkToken: linkToken,
	}
	// The following populates the AccountClient struct in order to have it
	// ready for the AccountM interface to use its methods
	newL.Account = &AccountClient{NewLink: newL}

	return newL
}

// The Delete method that acts upon the object generated by
// the Get method of the Link interface
func (n *NewLink) Delete() {
	url := fmt.Sprintf(Links, n.Link.Id)
	httpResponse, _ := n.client.deleteReq(url)

	fmt.Printf("Status Code: %d\n", httpResponse)
}

// The Delete method that will act upon the *LinkClient object
// and will delete certain link provided its link_id
func (l *LinkClient) Delete(linkId string) {
	url := fmt.Sprintf(Links, linkId)
	httpResponse, _ := l.deleteReq(url)

	fmt.Printf("Status Code: %d\n", httpResponse)
}