package fintoc

import (
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestLinkGet(t *testing.T) {
	httpmock.Activate()
	// linkToken := "link_nMNejK7BT8oGbvO4_token_GLtktZX5SKphRtJFe_yJTDWT"
	// url := formatUrl(fmt.Sprintf(LinkURL, linkToken))
	url := formatUrl(LinksAll)

	response, _ := ioutil.ReadFile("fixtures/link_object.json")
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))
	defer httpmock.DeactivateAndReset()

	client, _ := NewClient("")
	link := client.Link.All()

	assert.NotEmpty(t, link)
}
