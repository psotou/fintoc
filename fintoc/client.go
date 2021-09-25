package fintoc

import (
	"net/http"
)

const (
	BaseURL      = "https://api.fintoc.com/v1/"
	Accounts     = "accounts/%s"   // %s: {account_id}
	AccountsAll  = "accounts/"     //
	Movements    = "/movements/%s" // %s: {movement_id}
	MovementsAll = "/movements/"   //
	LinkURL      = "links/%s"      // %s: {link_token}
	LinksAll     = "links/"        //
)

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

// Fintoc API client
type APIClient struct {
	Secret string
	// We add the LinkM interface in this struct to allow
	// for a syntax like client.Link.Method()
	Link LinkM
}

// HTTP client interface to allow us to set instances of
// either http.Client or our mock http client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// MockClient sets the function that our mock Do method will run instead
// instead of the http.Client.Do method
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do method that overrides the http.Client.Do method
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// NewClient populates the APIClient
func NewClient(secret string) (*APIClient, error) {
	c := &APIClient{Secret: secret}
	// The following populates the LinkClient struct in order to have it
	// ready for the LinkM interface to use its methods
	c.Link = &LinkClient{APIClient: c}
	return c, nil
}
