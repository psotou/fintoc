package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

func (n *newLink) AccountAll() []Account {
	return n.link.Accounts
}

func (n *newLink) AccounOne(accountId string) *newAccount {
	var account Account
	url := fmt.Sprintf(Accounts+LinkToken, accountId, n.linkToken)
	byteData, _ := n.client.GetReq(url)
	err := json.Unmarshal(byteData, &account)
	if err != nil {
		log.Fatal(err.Error())
	}
	newA := &newAccount{client: n.client, account: account, linkToken: n.linkToken}
	return newA
}

type newAccount struct {
	client    *APIClient
	account   Account
	linkToken string
}
