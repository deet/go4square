package foursquare

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/url"
)

func TestExploreNear(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{"near": {"New York"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}

func TestNoParams(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{})
	assert.NotNil(t, error)
	assert.Nil(t, resp)
	assert.Equal(t, error.Meta.Code, 400, "when no arguments are passed, error code should be 400")
}

func TestExploreLatLng(t *testing.T) {
	api := GetApi()
	resp, error := api.Explore(url.Values{"ll": {"40.7,-74"}})
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}
