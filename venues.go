package foursquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	EXPLORE_URL = BASE_URL + "/venues/explore?"
	SEARCH_URL  = BASE_URL + "/venues/search?"
)

func (self *Client) Explore(values url.Values) (resp *ExploreResponse, error *ApiError) {

	values = self.appendClientParams(values)

	r, err := http.Get(EXPLORE_URL + values.Encode())
	defer r.Body.Close()

	if err != nil {
		error.msg = fmt.Sprintf("Error during get: %s", err.Error)
		return nil, error
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error.msg = fmt.Sprintf("Error reading stream: %s", err.Error)
		return nil, error
	}

	error = self.apiError(body)

	if error.Meta.Code != 200 {
		return nil, error
	}

	resp = new(ExploreResponse)
	json.Unmarshal(body, resp)

	return resp, nil
}

func (self *Client) Search(values url.Values) (resp *SearchResponse, error *ApiError) {

	values = self.appendClientParams(values)

	r, err := http.Get(SEARCH_URL + values.Encode())
	defer r.Body.Close()

	if err != nil {
		error.msg = fmt.Sprintf("Error during get: %s", err.Error)
		return nil, error
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error.msg = fmt.Sprintf("Error reading stream: %s", err.Error)
		return nil, error
	}

	error = self.apiError(body)

	if error.Meta.Code != 200 {
		return nil, error
	}

	resp = new(SearchResponse)
	json.Unmarshal(body, resp)

	return resp, nil
}
