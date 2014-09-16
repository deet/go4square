package foursquare

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
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
func TestExploreNoParams(t *testing.T) {
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

// Test explore method by passing a near argument
func TestSearchNear(t *testing.T) {
	api := GetApi()
	resp, error := api.Search(url.Values{"near": {"Chicago, IL"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Venues())
}

// Test Search method by passing no arguments
func TestSearchNoParams(t *testing.T) {
	api := GetApi()
	resp, error := api.Search(url.Values{})
	assert.NotNil(t, error)
	assert.Nil(t, resp)
	assert.Equal(t, error.Meta.Code, 400, "when no arguments are passed, error code should be 400")
}

// Test Search method by passing a latitude and longitude arguments
func TestSearchLatLng(t *testing.T) {
	api := GetApi()
	resp, error := api.Search(url.Values{"ll": {"40.7,-74"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Venues())
}
