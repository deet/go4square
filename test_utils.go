package foursquare

import (
	"encoding/json"
	"io/ioutil"
)

type ClientConfig struct {
	ClientId     string
	ClientSecret string
	ClientVersion string
}

var CONFIG ClientConfig

func GetApi() *Client {
	config := getClientConfig()
	return New(config.ClientId, config.ClientSecret, config.ClientVersion)
}

func getClientConfig() ClientConfig {
	if CONFIG.ClientId == "" {
		r, _ := ioutil.ReadFile("api_test.cfg")
		json.Unmarshal(r, &CONFIG)
	}
	return CONFIG
}