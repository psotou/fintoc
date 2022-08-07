package fintoc

import (
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
	Balance      Balance   `json:"balance"`
	RefreshedAt  time.Time `json:"refreshed_at"`
}

type Balance struct {
	Available int `json:"available"`
	Current   int `json:"current"`
	Limit     int `json:"limit"`
}

// func (c *client) parseLinktoken(resourceURL string) string {
// 	u, err := url.Parse(resourceURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	q := u.Query()
// 	q.Add("link_token", c.linkToken)
// 	u.RawQuery = q.Encode()

// 	return u.String()
// }

func (lc *linkClient) GetAccounts() []Account {
	return lc.Accounts
}

type accountClient struct {
	client
	Account
}

func (lc *linkClient) GetAccountByID(accountID string) accountClient {
	var accountIndex int
	for idx, account := range lc.Accounts {
		if accountID == account.Id {
			accountIndex = idx
		}
	}

	return accountClient{lc.client, lc.Accounts[accountIndex]}

	// why should I hit the endpoint again when the Link struct already has all
	// the information I need about the accounts associated to that Link??
}
