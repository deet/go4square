package foursquare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	CATEGORIES_URL = BASE_URL + "/venues/categories?"
)

func (self *Client) Categories() (resp *CategoriesResponse, error *ApiError) {

	values := self.appendClientParams(url.Values{})

	r, err := http.Get(CATEGORIES_URL + values.Encode())
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

	resp = new(CategoriesResponse)
	json.Unmarshal(body, resp)

	return resp, nil
}
