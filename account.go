package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
)

type Account struct {
	Id           string    `json:"id"`
	Object       string    `json:"object"`
	Name         string    `json:"name"`
	OfficialName string    `json:"official_name"`
	Number       string    `json:"number"`
	HolderId     string    `json:"holder_id"`
	HolderName   string    `json:"holder_name"`
	Type         string    `json:"type"`
	Currency     string    `json:"currency"`
	Balance      Blnc      `json:"balance"`
	RefreshedAt  time.Time `json:"refreshed_at"`
}

type Blnc struct {
	Available int `json:"available"`
	Current   int `json:"current"`
	Limit     int `json:"limit"`
}

type NewAccount struct {
	client *APIClient
	// We make this field anonymous in order to directly get
	// the attributes of the Account object
	Account
	linkToken string
	// Since Get() method returns a *NewAccount, we need to add
	// the MovementM interface in this struct to allow for a
	// syntax like account.Movement.Method()
	Movement MovementM
}
type AccountClient struct {
	*NewLink
}

type AccountM interface {
	All() []Account
	Get(string) *NewAccount
}

func (a *AccountClient) All() []Account {
	u, _ := url.Parse(AccountsAll)
	q := u.Query()
	q.Add("link_token", a.linkToken)
	u.RawQuery = q.Encode()

	var accounts []Account
	dataBytes, _ := a.client.getReq(u.String())
	err := json.Unmarshal(dataBytes, &accounts)
	if err != nil {
		log.Fatal(err.Error())
	}
	return accounts
}

func (a *AccountClient) Get(accountId string) *NewAccount {
	uri := fmt.Sprintf(Accounts, accountId)
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("link_token", a.linkToken)
	u.RawQuery = q.Encode()

	var account Account
	byteData, _ := a.client.getReq(u.String())
	err := json.Unmarshal(byteData, &account)
	if err != nil {
		log.Fatal(err.Error())
	}
	newA := &NewAccount{
		client:    a.client,
		Account:   account,
		linkToken: a.linkToken,
	}
	// The following populates the MovementClient struct in order to have it
	// ready for the MovementM interface to use its methods
	newA.Movement = &MovementClient{NewAccount: newA}
	return newA
}
