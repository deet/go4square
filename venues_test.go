package foursquare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/url"
)

// Test explore method by passing a near argument
func TestExploreNear(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{"near": {"Chicago, IL"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Venues())
}

// Test explore method by passing no arguments
func TestNoParams(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{})
	assert.NotNil(t, error)
	assert.Nil(t, resp)
	assert.Equal(t, error.Meta.Code, 400, "when no arguments are passed, error code should be 400")
}

// Test explore method by passing a latitude and longitude arguments
func TestExploreLatLng(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{"ll": {"40.7,-74"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Venues())
}
