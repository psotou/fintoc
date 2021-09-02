package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

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
	url := fmt.Sprintf(Accounts+Movements+LinkToken, n.Id, movementId, n.linkToken)
	dataBytes, _ := n.client.getReq(url)
	err := json.Unmarshal(dataBytes, &movement)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movement
}
