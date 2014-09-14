go4square[![GoDoc](https://godoc.org/github.com/rfsbraz/go4square?status.png)](https://godoc.org/github.com/rfsbraz/go4square) [![Build Status](https://travis-ci.org/rfsbraz/go4square.svg?branch=master)](https://travis-ci.org/rfsbraz/go4square)
=========

Go wrapper for the foursquare v2 API

Example
-------

```go
package main

import (
        "github.com/rfsbraz/go4square"
)

func main() {
	values := url.Values{}

	api := go4square.New("ID_HERE", "SECRET_HERE", "VERSION_HERE")
	resp, err := api.Explore(url.Values{"ll": {"44.3,37.2"}})
	if err != nil {
		// Handle the error
	} else {
		// Enjoy the venues
	}

	categoryResp, err := api.Categories()
	if err != nil {
		// Handle the error
	} else {
		// Enjoy the categories
	}
}
```

Implemented
--------

* Nothing just yet, but the venue endpoint will be the first to be complete.
* Categories response, excluding URLs to icons.
* Categories from Venues response may not have sub categories populated.

TODO
-----------

* Every foursquare endpoint
