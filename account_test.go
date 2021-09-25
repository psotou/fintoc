package fintoc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/psotou/fintoc-sdk/fixtures"
)

func TestAccountAll(t *testing.T) {
	jsonLink := fixtures.LinkObject
	jsonAccounts := fmt.Sprintf("[%v]", fixtures.AccountObject)
	rLink := io.NopCloser(bytes.NewReader([]byte(jsonLink)))
	rAccount := io.NopCloser(bytes.NewReader([]byte(jsonAccounts)))
	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rLink,
			}, nil
		},
	}

	client, err := NewClient("secret")
	if err != nil {
		t.Error("Testing failed!")
	}
	link := client.Link.Get("secret")

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rAccount,
			}, nil
		},
	}
	accounts := link.Account.All()

	if len(accounts) == 0 {
		t.Error("Failed. Empty object.")
		return
	}
}

func TestAccountGet(t *testing.T) {
	jsonLink := fixtures.LinkObject
	jsonAccount := fixtures.AccountObject
	rLink := io.NopCloser(bytes.NewReader([]byte(jsonLink)))
	rAccount := io.NopCloser(bytes.NewReader([]byte(jsonAccount)))
	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rLink,
			}, nil
		},
	}

	client, err := NewClient("secret")
	if err != nil {
		t.Error("Testing failed!")
	}
	link := client.Link.Get("secret")

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rAccount,
			}, nil
		},
	}
	account := link.Account.Get("accId")

	if account == nil {
		t.Error("Failed. Empty object.")
		return
	}
}
