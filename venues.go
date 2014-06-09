package foursquare

import (
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
)

const (
	EXPLORE_URL = BASE_URL + "/venues/explore?"
)

func (self *Client) Explore(values url.Values) (resp *ExploreResponse, err error) {

	values = self.appendClientParams(values)

	r, err := http.Get(EXPLORE_URL + values.Encode())
	defer r.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Error during get: %s", err.Error)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading stream: %s", err.Error)
	}

	apiError := self.apiError(body)

	if  apiError.Meta.Code != 200 {
		return nil, fmt.Errorf("Api error (%s): %s", apiError.Meta.Code, apiError.Meta.ErrorDetail)
	}

	resp = new(ExploreResponse)
	json.Unmarshal(body, resp)

	return resp, nil
}