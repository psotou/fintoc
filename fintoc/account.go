package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

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
	return a.Accounts
}

func (a *AccountClient) Get(accountId string) *NewAccount {
	uri := fmt.Sprintf(Accounts, accountId)
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("link_token", a.linkToken)
	u.RawQuery = q.Encode()

	var account Account
	byteData, _ := a.client.GetReq(u.String())
	err = json.Unmarshal(byteData, &account)
	if err != nil {
		log.Fatal(err.Error())
	}
	newA := &NewAccount{client: a.client, Account: account, linkToken: a.linkToken}
	// The following populates the MovementClient struct in order to have it
	// ready for the MovementM interface to use its methods
	newA.Movement = &MovementClient{NewAccount: newA}
	return newA
}
