package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
)

type Movement struct {
	Id               string    `json:"id"`
	Object           string    `json:"object"`
	Amount           int       `json:"amount"`
	PostDate         time.Time `json:"post_date"`
	Description      string    `json:"description"`
	TransactionDate  time.Time `json:"transaction_date"`
	Currency         string    `json:"currency"`
	ReferenceId      string    `json:"reference_id"`
	Type             string    `json:"type"`
	Pending          bool      `json:"pending"`
	RecipientAccount Recipient `json:"recipient_account"`
	SenderAccount    Sender    `json:"sender_account"`
	Comment          string    `json:"comment"`
}

type Recipient struct {
	HolderId    string      `json:"holder_id"`
	HolderName  string      `json:"holder_name"`
	Number      string      `json:"number"`
	Institution Institution `json:"institution"`
}

type Sender struct {
	HolderId    string      `json:"holder_id"`
	HolderName  string      `json:"holder_name"`
	Number      string      `json:"number"`
	Institution Institution `json:"institution"`
}

type Institution struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (ac *accountClient) GetMovements(opts ...MovementsOptions) []Movement {
	var movements []Movement
	var parsedURL string
	movementsURL := fmt.Sprintf(Accounts+MovementsAll, ac.Account.Id)
	parsedURL = ac.parseLinktokenParam(movementsURL)
	if len(opts) > 0 {
		parsedURL = opts[0].parseOptionParams(movementsURL, ac.linkToken)
	}

	response, err := ac.client.get(parsedURL)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response, &movements)

	return movements
}

func (ac *accountClient) GetMovementByID(movementID string) Movement {
	var movement Movement
	movementURL := fmt.Sprintf(Accounts+Movements, ac.Account.Id, movementID)
	parsedURL := ac.parseLinktokenParam(movementURL)
	response, err := ac.client.get(parsedURL)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response, &movement)

	return movement
}

func (ac *accountClient) parseLinktokenParam(resourceURL string) string {
	u, err := url.Parse(resourceURL)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("link_token", ac.linkToken)
	u.RawQuery = q.Encode()

	return u.String()
}

type MovementsOptions struct {
	Since   string
	Until   string
	PerPage string
}

func (mO *MovementsOptions) parseOptionParams(resourceURL, linkToken string) string {
	u, _ := url.Parse(resourceURL)
	q := u.Query()
	q.Add("link_token", linkToken)
	if mO.Since != "" {
		q.Add("since", mO.Since)
	}
	if mO.Until != "" {
		q.Add("until", mO.Until)
	}
	if mO.PerPage != "" {
		q.Add("per_page", mO.PerPage)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
