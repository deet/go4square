package foursquare

import "fmt"

// https://developer.foursquare.com/docs/responses/venue
type Venue struct {
	Id         string
	Name       string
	Location   Location
	Categories []Category
}

type Location struct {
	Address     string
	CrossStreet string
	City        string
	State       string
	PostalCode  string
	Country     string
	Lat         float64
	Lng         float64
	Distance    float64
}

type Category struct {
	Id         string
	Name       string
	PluralName string
	ShortName  string
	Primary    bool
	Categories []Category
}

// https://developer.foursquare.com/overview/responses

type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred after reading Offset bytes
}

type ApiError struct {
	Meta MetaInfo
	msg  string
}

func (e *ApiError) Error() string {
	if e.msg == "" {
		return e.msg
	} else {
		return fmt.Sprintf("Api error (%d): %s", e.Meta.Code, e.Meta.ErrorDetail)
	}
}

type MetaInfo struct {
	Code        int32
	ErrorDetail string
	ErrorType   string
}

// https://developer.foursquare.com/docs/venues/explore
type ExploreResponse struct {
	Response ExploreBody
}

type ExploreBody struct {
	Groups []ExploreGroup
}

type ExploreGroup struct {
	Items []ExploreItem
}

type ExploreItem struct {
	Venue Venue
}

func (resp *ExploreResponse) Venues() (venues []Venue) {
	for _, group := range resp.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues
}

type CategoriesBody struct {
	Categories []Category
}

type CategoriesResponse struct {
	Response CategoriesBody
}
