package fintoc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/psotou/fintoc-sdk/fixtures"
)

func TestLinkAll(t *testing.T) {
	jsonResponse := fmt.Sprintf("[%v]", fixtures.LinkObject)
	r := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	client, err := NewClient("secret")
	if err != nil {
		t.Error("Testing failed!")
	}
	result := client.Link.All()

	if len(result) == 0 {
		t.Error("Failed. Empty object.")
		return
	}
	if result[0].Accounts == nil {
		t.Error("Link without account object.")
		return
	}
}

func TestLinkGet(t *testing.T) {
	jsonResponse := fixtures.LinkObject
	r := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	client, err := NewClient("secret")
	if err != nil {
		t.Error("Testing failed!")
	}
	result := client.Link.Get("linkToken")

	if result.Accounts == nil {
		t.Error("Link without account object.")
		return
	}
}
