package yelp

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	// apiHost is the base URL for the Yelp API
	apiHost = "https://api.yelp.com"

	// searchPath is the path to search for businesses
	searchPath = "/v3/businesses/search"

	// businessPath is the path to get a business by its id
	businessPath = "/v3/businesses/%s"

	// tokenPath is the path to fetch the bearer's access token
	tokenPath = "/oauth2/token"
)

// Client defines the current available Yelp API requests that can be made.
type Client interface {
	Search(SearchOptions) (SearchResults, error)
	BusinessByID(businessID string) (Business, error)
}

// client implements the Client interface.
type client struct {
	*http.Client
	credentials Credentials
}

// New returns a new Yelp client.
func New(c *http.Client, credentials Credentials) *client {
	return &client{
		Client:      c,
		credentials: credentials,
	}
}

// FetchToken makes a post request to get the auth token for credentials passed
// to the client on initialization.
func (c *client) FetchToken() error {
	respBody := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}{}
	_, err := c.postForm(apiHost+tokenPath, c.credentials.URLValues(), &respBody)
	if err != nil {
		return err
	}

	c.credentials.AccessToken = respBody.AccessToken
	c.credentials.ExpiryDate = time.Now().Add(time.Duration(respBody.ExpiresIn) * time.Second)
	return nil
}

// Search makes a request given the options passed in.
func (c *client) Search(so SearchOptions) (SearchResults, error) {
	respBody := SearchResults{}
	if !so.IsValid() {
		return respBody, errors.New("SearchOptions provided is not valid. Please see yelp/search.go for more details.")
	}

	urlStr := apiHost + searchPath + "?" + so.URLValues().Encode()
	_, err := c.authedDo("GET", urlStr, nil, nil, &respBody)
	return respBody, err
}

// BusinessByID looks for a business information by its id.
func (c *client) BusinessByID(businessID string) (Business, error) {
	respBody := Business{}

	urlStr := apiHost + fmt.Sprintf(businessPath, businessID)
	_, err := c.authedDo("GET", urlStr, nil, nil, &respBody)
	return respBody, err
}

// authedDo fetches the access token again if it is expired and constructs a
// request with the Authorization Header set with the access token. The response
// body is decoded into v.
func (c *client) authedDo(method string, url string, body io.Reader, headers map[string]string, v interface{}) (*http.Response, error) {
	if !c.credentials.IsValid() {
		if err := c.FetchToken(); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("Authorization", "Bearer "+c.credentials.AccessToken)

	resp, err := c.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return resp, fmt.Errorf("Yelp request failed with status %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// postForm makes a POST request with form values and decodes the response body
// into v.
func (c *client) postForm(url string, data url.Values, v interface{}) (*http.Response, error) {
	resp, err := c.PostForm(url, data)
	if err != nil {
		return resp, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
