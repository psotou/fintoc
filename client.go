package fintoc

import (
	"net/http"
)

const (
	BaseURL          = "https://api.fintoc.com/v1/"
	Accounts         = "accounts/%s"      // %s: {account_id}
	AccountsAll      = "accounts/"        //
	Movements        = "/movements/%s"    // %s: {movement_id}
	MovementsAll     = "/movements/"      //
	Links            = "links/%s"         // %s: {link_token}
	LinksAll         = "links/"           //
	Subscriptions    = "subscriptions/%s" // %s: {subscription_id}
	SubscriptionsAll = "subscriptions"    //
)

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
