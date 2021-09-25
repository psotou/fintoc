package fintoc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/psotou/fintoc/fixtures"
)

func TestMovementAll(t *testing.T) {
	jsonLink := fixtures.LinkObject
	jsonAccount := fixtures.AccountObject
	jsonMovements := fmt.Sprintf("[%v]", fixtures.MovementObject)
	rLink := io.NopCloser(bytes.NewReader([]byte(jsonLink)))
	rAccount := io.NopCloser(bytes.NewReader([]byte(jsonAccount)))
	rMovements := io.NopCloser(bytes.NewReader([]byte(jsonMovements)))
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
	link := client.Link.Get("secret") // .Account.Get("accountId").Movement.All()

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rAccount,
			}, nil
		},
	}
	account := link.Account.Get("accountId")

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rMovements,
			}, nil
		},
	}
	movements := account.Movement.All()

	if len(movements) == 0 {
		t.Error("Failed. Empty object.")
		return
	}
}

func TestMovementGet(t *testing.T) {
	jsonLink := fixtures.LinkObject
	jsonAccount := fixtures.AccountObject
	jsonMovement := fixtures.MovementObject
	rLink := io.NopCloser(bytes.NewReader([]byte(jsonLink)))
	rAccount := io.NopCloser(bytes.NewReader([]byte(jsonAccount)))
	rMovements := io.NopCloser(bytes.NewReader([]byte(jsonMovement)))
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
	link := client.Link.Get("secret") // .Account.Get("accountId").Movement.All()

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rAccount,
			}, nil
		},
	}
	account := link.Account.Get("accountId")

	Client = &MockClient{
		DoFunc: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       rMovements,
			}, nil
		},
	}
	movement := account.Movement.Get("movId")

	// we use the zero-valued struct (Movement{})
	if movement == (Movement{}) {
		t.Error("Failed. Empty object.")
		return
	}
}
