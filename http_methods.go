package fintoc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Formats resource url with the base url
func formatUrl(resourceUrl string) string {
	return fmt.Sprintf("%s%s", BaseURL, resourceUrl)
}

// Function requestMethod for requests with custom errors
func (client *APIClient) requestMethod(reqMethod, resourceUrl string, reader io.Reader) (*http.Response, error) {
	url := formatUrl(resourceUrl)
	req, err := http.NewRequest(reqMethod, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{
		"Accept":        []string{"application/json"},
		"Authorization": []string{client.Secret},
	}
	if reqMethod == http.MethodPatch {
		req.Header.Add("Content-Type", "application/json")
	}
	res, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	// we manage the custom errors in this block
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
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
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// updateReq updates the link according a payload
func (client *APIClient) updateReq(url string, payload io.Reader) ([]byte, error) {
	res, err := client.requestMethod(http.MethodPatch, url, payload)
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

// deleteReq deletes a link
func (client *APIClient) deleteReq(url string) (int, error) {
	res, err := client.requestMethod(http.MethodDelete, url, nil)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
