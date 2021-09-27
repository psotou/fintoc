package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"
)

type Movement struct {
	Id               string       `json:"id"`
	Object           string       `json:"object"`
	Amount           int          `json:"amount"`
	PostDate         string       `json:"post_date"`
	Description      string       `json:"description"`
	TransactionDate  time.Time    `json:"transaction_date"`
	Currency         string       `json:"currency"`
	ReferenceId      string       `json:"reference_id"`
	Type             string       `json:"type"`
	Pending          bool         `json:"pending"`
	RecipientAccount RcpntAccount `json:"recipient_account"`
	SenderAccount    Sender       `json:"sender_account"`
	Comment          string       `json:"comment"`
}

type RcpntAccount struct {
	HolderId    string      `json:"holder_id"`
	HolderName  string      `json:"holder_name"`
	Number      string      `json:"number"`
	Institution Institucion `json:"institution"`
}

type Sender struct {
	HolderId    string      `json:"holder_id"`
	HolderName  string      `json:"holder_name"`
	Number      string      `json:"number"`
	Institution Institucion `json:"institution"`
}

type MovementClient struct {
	*NewAccount
}

type MovementM interface {
	All(opts ...Params) []Movement
	Get(string) Movement
}

type Params struct {
	Since   string
	Until   string
	PerPage string
}

func (n *NewAccount) All(opts ...Params) []Movement {
	uri := fmt.Sprintf(Accounts+MovementsAll, n.Id)
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("link_token", n.linkToken)

	// Since opts is a slice of elements of type Params and
	// we know that we're ONLY accepting one element
	// we proceed and extract by hand the element to be used
	if len(opts) > 0 {
		if opts[0].Since != "" {
			q.Add("since", opts[0].Since)
		}
		if opts[0].Until != "" {
			q.Add("until", opts[0].Until)
		}
		if opts[0].PerPage != "" {
			q.Add("per_page", opts[0].PerPage)
		}
	}

	u.RawQuery = q.Encode()
	var movements []Movement
	dataBytes, _ := n.client.getReq(u.String())
	err := json.Unmarshal(dataBytes, &movements)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movements
}

func (n *NewAccount) Get(movementId string) Movement {
	var movement Movement
	uri := fmt.Sprintf(Accounts+Movements, n.Id, movementId)
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("link_token", n.linkToken)
	u.RawQuery = q.Encode()
	dataBytes, _ := n.client.getReq(u.String())
	err := json.Unmarshal(dataBytes, &movement)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movement
}
