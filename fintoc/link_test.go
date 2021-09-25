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
	jsonLinks := fmt.Sprintf("[%v]", fixtures.LinkObject)
	r := io.NopCloser(bytes.NewReader([]byte(jsonLinks)))
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
	links := client.Link.All()

	if len(links) == 0 {
		t.Error("Failed. Empty object.")
		return
	}
	if links[0].Accounts == nil {
		t.Error("Link without account object.")
		return
	}
}

func TestLinkGet(t *testing.T) {
	jsonLink := fixtures.LinkObject
	r := io.NopCloser(bytes.NewReader([]byte(jsonLink)))
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
	link := client.Link.Get("linkToken")

	if link == nil {
		t.Error("Link object empty.")
		return
	}

	if link.Accounts == nil {
		t.Error("Link without account object.")
		return
	}
}
