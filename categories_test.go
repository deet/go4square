package foursquare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test that categories method returns results
func TestCategories(t *testing.T) {
	api := GetApi()
	resp, error := api.Categories()
	assert.Nil(t, error)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Response.Categories)
}

// Test explore method by passing no arguments
func TestCategoriesContains(t *testing.T) {
	api := GetApi()
	resp, error := api.Categories()
	assert.Nil(t, error)
	assert.NotNil(t, resp)

	assert.NotEmpty(t, resp.Response.Categories)

	topLevelIds := []string{}
	for _, cat := range resp.Response.Categories {
		topLevelIds = append(topLevelIds, cat.Id)
	}

	assert.Contains(t, topLevelIds, "4d4b7104d754a06370d81259")

	// the first category is Arts and Entertainment. This is brittle in that if Foursquare adds one before the test will fail
	aaeCat := resp.Response.Categories[0]
	assert.Equal(t, aaeCat.Id, "4d4b7104d754a06370d81259")
	assert.Equal(t, aaeCat.Name, "Arts & Entertainment")
	assert.Equal(t, aaeCat.PluralName, "Arts & Entertainment")
	assert.Equal(t, aaeCat.ShortName, "Arts & Entertainment")
	assert.NotEmpty(t, aaeCat.Categories)

	// the first category in Arts and Entertainment is Aquarium
	aqCat := aaeCat.Categories[0]
	assert.Equal(t, aqCat.Id, "4fceea171983d5d06c3e9823")
	assert.Equal(t, aqCat.Name, "Aquarium")
	assert.Equal(t, aqCat.PluralName, "Aquariums")
	assert.Equal(t, aqCat.ShortName, "Aquarium")
	assert.Empty(t, aqCat.Categories, "Acquarium has no sub-categories")
}
