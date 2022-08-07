package fintoc

import (
	"fmt"
	"io"
	"net/http"
)

const (
	BaseURL      = "https://api.fintoc.com/v1/"
	Accounts     = "accounts/%s"   // %s: {account_id}
	AccountsAll  = "accounts"      //
	Movements    = "/movements/%s" // %s: {movement_id}
	MovementsAll = "/movements"    //
	Links        = "links/%s"      // %s: {link_token}
	LinksAll     = "links"         //
)

type client struct {
	secretKey string
	linkToken string
}

func NewClient(secret string) (*client, error) {
	return &client{secretKey: secret}, nil
}

func formatURL(resourceURL string) string {
	return fmt.Sprintf("%s%s", BaseURL, resourceURL)
}

func (c *client) requestMethod(reqMethod, resourceURL string, reader io.Reader) (*http.Response, error) {
	url := formatURL(resourceURL)
	req, err := http.NewRequest(reqMethod, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Authorization": []string{c.secretKey},
	}
	if reqMethod == http.MethodPatch {
		req.Header.Add("Content-Type", "application/json")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *client) get(url string) ([]byte, error) {
	res, err := c.requestMethod(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *client) patch(url string, payload io.Reader) ([]byte, error) {
	res, err := c.requestMethod(http.MethodPatch, url, payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *client) delete(url string) (int, error) {
	res, err := c.requestMethod(http.MethodDelete, url, nil)
	if err != nil {
		return -1, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
