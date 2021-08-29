package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

type AccountClient struct {
	ApiClient *APIClient
}

type AccountM interface {
	All() []Account
	Get(string) *newAccount
}

// func (a *AccountClient) All() []Account {
// 	var account []Account
// 	url := fmt.Sprintf(AccountsAll+LinkToken, a.ApiClient.LinkToken)
// 	dataBytes, _ := a.ApiClient.GetReq(url)
// 	err := json.Unmarshal(dataBytes, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return account
// }

func (n *newLink) AccountAll() []Account {
	return n.link.Accounts
}

// // Old way with just one return value
// func (a *AccountClient) Get(accountId string) Account {
// 	var account Account
// 	url := fmt.Sprintf(Accounts+LinkToken, accountId, a.ApiClient.LinkToken)
// 	byteData, _ := a.ApiClient.GetReq(url)
// 	err := json.Unmarshal(byteData, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return account
// }

// New way with two return values
// func (a *AccountClient) Get(accountId string) (Account, []Movement) {
// 	var (
// 		account Account
// 		mov     []Movement
// 	)
// Account
// 	url := fmt.Sprintf(Accounts+LinkToken, accountId, a.ApiClient.LinkToken)
// 	byteData, _ := a.ApiClient.GetReq(url)
// 	err := json.Unmarshal(byteData, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	// Movements of the account
// 	urlM := fmt.Sprintf(Accounts+MovementsAll+LinkToken, accountId, a.ApiClient.LinkToken)
// 	byteDataM, _ := a.ApiClient.GetReq(urlM)
// 	errM := json.Unmarshal(byteDataM, &mov)
// 	if errM != nil {
// 		log.Fatal(errM.Error())
// 	}
// 	return account, mov
// }

// func (a *AccountClient) Get(accountId string) *newAccount {
// 	var account Account
// 	url := fmt.Sprintf(Accounts+LinkToken, accountId, a.ApiClient.LinkToken)
// 	byteData, _ := a.ApiClient.GetReq(url)
// 	err := json.Unmarshal(byteData, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	newA := &newAccount{client: a.ApiClient, aa: account}
// 	return newA
// }

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

// func (n *newAccount) AllMov() []Movement {
// 	var m []Movement
// 	urlM := fmt.Sprintf(Accounts+MovementsAll+LinkToken, n.aa.Id, n.client.ApiClient.LinkToken)
// 	byteDataM, _ := n.client.ApiClient.GetReq(urlM)
// 	errM := json.Unmarshal(byteDataM, &m)
// 	if errM != nil {
// 		log.Fatal(errM.Error())
// 	}
// 	return m
// }

// // Thw way to return a function (not much use tbh)
// func (a *AccountClient) Geti(accountId string) (Account, func() []Movement) {
// 	var account Account
// 	url := fmt.Sprintf(Accounts+LinkToken, accountId, a.ApiClient.LinkToken)
// 	byteData, _ := a.ApiClient.GetReq(url)
// 	err := json.Unmarshal(byteData, &account)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return account, func() []Movement {
// 		var movements []Movement

// 		url := fmt.Sprintf(Accounts+MovementsAll+LinkToken, accountId, a.ApiClient.LinkToken)
// 		dataBytes, _ := a.ApiClient.GetReq(url)
// 		err := json.Unmarshal(dataBytes, &movements)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 		return movements
// 	}

// }
