package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

type MovementClient struct {
	*NewAccount
}

type MovementM interface {
	All() []Movement
	Get(string) Movement
}

func (n *NewAccount) All() []Movement {
	var movements []Movement

	url := fmt.Sprintf(Accounts+MovementsAll+LinkToken, n.Id, n.linkToken)
	dataBytes, _ := n.client.GetReq(url)
	err := json.Unmarshal(dataBytes, &movements)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movements
}

func (n *NewAccount) Get(movementId string) Movement {
	var movement Movement
	url := fmt.Sprintf(Accounts+Movements+LinkToken, n.Id, movementId, n.linkToken)
	dataBytes, _ := n.client.GetReq(url)
	err := json.Unmarshal(dataBytes, &movement)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movement
}
