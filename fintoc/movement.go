package fintoc

import (
	"encoding/json"
	"fmt"
	"log"
)

type MovementClient struct {
	ApiClient *APIClient
}

type MovementM interface {
	All() []Movement
	// Get(string) Movement
}

// func (m *MovementClient) All() []Movement {
// 	var movements []Movement

// 	url := fmt.Sprintf(Accounts+MovementsAll+LinkToken, "d", m.ApiClient.LinkToken)
// 	dataBytes, _ := m.ApiClient.GetReq(url)
// 	err := json.Unmarshal(dataBytes, &movements)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	return movements
// }

func (n *newAccount) MovementAll() []Movement {
	var movements []Movement

	url := fmt.Sprintf(Accounts+MovementsAll+LinkToken, n.account.Id, n.linkToken)
	dataBytes, _ := n.client.GetReq(url)
	err := json.Unmarshal(dataBytes, &movements)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movements
}

func (n *newAccount) MovementOne(movementId string) Movement {
	var movement Movement
	url := fmt.Sprintf(Accounts+Movements+LinkToken, n.account.Id, movementId, n.linkToken)
	dataBytes, _ := n.client.GetReq(url)
	err := json.Unmarshal(dataBytes, &movement)
	if err != nil {
		log.Fatal(err.Error())
	}
	return movement
}
