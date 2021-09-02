package fintoc

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAccountGet(t *testing.T) {
	httpmock.Activate()
	// linkToken := "link_nMNejK7BT8oGbvO4_token_GLtktZX5SKphRtJFe_yJTDWT"
	url := formatUrl(fmt.Sprintf(Accounts, "acc_nMNejK7BT8oGbvO4"))
	response, _ := ioutil.ReadFile("fixtures/account_object.json")
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient("")
	// link := client.Link.Get()
	// account := client.Link.Get(linkToken).Account.Get("acc_nMNejK7BT8oGbvO4")
	assert.NotEmpty(t, client)
}
