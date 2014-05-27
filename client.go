package foursquare

import (
	"encoding/json"
	"net/url"
)

const (
	BASE_URL = "https://api.foursquare.com/v2"
)

type Client struct {
	clientId     string
	clientSecret string
	apiVersion   string
}

func New(clientId string, clientSecret string, apiVersion string) *Client {
	client := new(Client)
	client.clientId = clientId
	client.clientSecret = clientSecret
	client.apiVersion = apiVersion
	return client
}

func (self *Client) appendClientParams(params url.Values) url.Values {
	params.Set("client_id", self.clientId)
	params.Set("client_secret", self.clientSecret)
	params.Set("v", self.apiVersion)
	return params
}

func (self *Client) apiError(content []byte) *ApiError {
	error := new(ApiError)
	json.Unmarshal(content, &error)
	return error
}