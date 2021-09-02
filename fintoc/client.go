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
	Accounts     = "accounts/%s"   // %s: {account_id}
	AccountsAll  = "accounts/"     //
	Movements    = "/movements/%s" // %s: {movement_id}
	MovementsAll = "/movements/"   //
	// LinkToken    = "?link_token=%s" //
	LinkURL  = "links/%s" // %s: {link_token}
	LinksAll = "links/"   //
)

// Fintoc API client
type APIClient struct {
	Secret string
	Client *http.Client
	// We add the MovementM interface in this struct to allow
	// for a syntax like client.Link.Method()
	Link LinkM
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
	c := &APIClient{Secret: secret, Client: &http.Client{}}
	// The following populates the LinkClient struct in order to have it
	// ready for the LinkM interface to use its methods
	c.Link = &LinkClient{APIClient: c}
	return c, nil
}

// Formats resource url with the base url
func formatUrl(resourceUrl string) string {
	return fmt.Sprintf("%s%s", BaseURL, resourceUrl)
}

// Function requestMethod for requests with custom errors
func (client *APIClient) requestMethod(reqMethod, resourceUrl string, reader io.Reader) (*http.Response, error) {
	url := formatUrl(resourceUrl)
	req, err := http.NewRequest(reqMethod, url, reader)
	if err != nil {
		log.Fatal(err.Error())
	}

	if reqMethod == http.MethodPatch {
		req.Header.Add("Content-Type", "application/json")
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

// getReq takes the response with the custom error and handles it appropriately
func (client *APIClient) getReq(url string) ([]byte, error) {
	res, err := client.requestMethod(http.MethodGet, url, nil)
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

// updateReq updates the link according a payload
func (client *APIClient) updateReq(url string, payload io.Reader) ([]byte, error) {
	res, err := client.requestMethod(http.MethodPatch, url, payload)
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

// deleteReq deletes a link
func (client *APIClient) deleteReq(url string) (int, error) {
	res, err := client.requestMethod(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
