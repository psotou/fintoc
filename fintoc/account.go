package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

func (n *NewLink) AccountAll() []Account {
	return n.Accounts
}

func (n *NewLink) AccounOne(accountId string) *NewAccount {
	var account Account
	url := fmt.Sprintf(Accounts+LinkToken, accountId, n.linkToken)
	byteData, _ := n.client.GetReq(url)
	err := json.Unmarshal(byteData, &account)
	if err != nil {
		log.Fatal(err.Error())
	}
	newA := &NewAccount{client: n.client, Account: account, linkToken: n.linkToken}
	return newA
}

type NewAccount struct {
	client *APIClient
	// We make this field anonymous in order to directly get
	// the attributes of the Account object
	Account
	linkToken string
}
