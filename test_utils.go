package foursquare

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
		r, error := ioutil.ReadFile("api_test.cfg")
		if error != nil {
			//File doesn't exist, let's try environment variables (works well for travis)
			CONFIG = *new(ClientConfig)
			CONFIG.ClientId = os.Getenv("FOURSQUARE_CLIENT_ID")
			CONFIG.ClientSecret = os.Getenv("FOURSQUARE_CLIENT_SECRET")
			CONFIG.ClientVersion = os.Getenv("FOURSQUARE_CLIENT_VERSION")
		} else {
			// No error, the file exists
			json.Unmarshal(r, &CONFIG)
		}
	}
	return CONFIG
}