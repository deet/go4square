package foursquare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"net/url"
)

type ClientConfig struct {
	ClientId     string
	ClientSecret string
	ClientVersion string
}

var CONFIG ClientConfig

func TestExploreNear(t *testing.T) {
	api := getApi()
	resp, error := api.Explore(url.Values{"near": {"New York"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}

func TestExploreLatLng(t *testing.T) {
	api := getApi()
	resp, error := api.Explore(url.Values{"ll": {"40.7,-74"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}

func getApi() *Client {
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
