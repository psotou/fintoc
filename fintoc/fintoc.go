package fintoc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	BaseURL      = "https://api.fintoc.com/v1/"
	Accounts     = "accounts/%s"    // ?link_token=%s" // %s: {account_id}
	AccountsAll  = "accounts/"      // ?link_token=%s"   //
	Movements    = "/movements/%s"  // %s: {movement_id}
	MovementsAll = "/movements/"    // ?link_token=%s" // %s: {link_token}
	LinkToken    = "?link_token=%s" //
	LinkURL      = "links/%s"       // %s: {link_token}
	LinksAll     = "links/"         //
)

// Fintoc API client
type APIClient struct {
	Secret string
	Client *http.Client

	// The old way, when I attempted to include interfaces
	// Link     LinkM
	// Account  AccountM
	// Movement MovementM
}

type Link struct {
	Id          string      `json:"id"`
	Object      string      `json:"object"`
	UserName    string      `json:"username"`
	LinkToken   string      `json:"link_token"`
	Mode        string      `json:"mode"`
	Active      bool        `json:"active"`
	Status      string      `json:"status"`
	HolderType  string      `json:"holder_type"`
	CreatedAt   string      `json:"created_at"`
	Institution Institucion `json:"institution"`
	Accounts    []Account   `json:"accounts"`
}

type Account struct {
	Id           string `json:"id"`
	Object       string `json:"object"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	Number       string `json:"number"`
	HolderId     string `json:"holder_id"`
	HolderName   string `json:"holder_name"`
	Type         string `json:"type"`
	Currency     string `json:"currency"`
	Balance      Blnc   `json:"balance"`
	RefreshedAt  string `json:"refreshed_at"`
}

type Blnc struct {
	Available int `json:"available"`
	Current   int `json:"current"`
	Limit     int `json:"limit"`
}

type Movement struct {
	Id               string       `json:"id"`
	Object           string       `json:"object"`
	Amount           int          `json:"amount"`
	PostDate         string       `json:"post_date"`
	Description      string       `json:"description"`
	TransactionDate  string       `json:"transaction_date"`
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

type Institucion struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// NewClient populates the APIClient
func NewClient(secret string) (*APIClient, error) {
	return &APIClient{Secret: secret, Client: &http.Client{}}, nil
}

// Formats resource url with the base url
func FormatUrl(resourceUrl string) string {
	return fmt.Sprintf("%s%s", BaseURL, resourceUrl)
}

// Function Reqwest for requests with custom errors
func (client *APIClient) Reqwest(resourceUrl string) (*http.Response, error) {
	url := FormatUrl(resourceUrl)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", client.Secret)
	res, err := client.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// we manage the custom errors in this block
	if res.StatusCode != http.StatusOK {
		var apiErr Error
		err := json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}
		return nil, &apiErr
	}

	return res, nil
}

// GetReq takes the response with the custom error and handles it appropriately
func (client *APIClient) GetReq(url string) ([]byte, error) {
	res, err := client.Reqwest(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
