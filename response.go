package foursquare

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
}

// https://developer.foursquare.com/overview/responses
type ApiError struct {
	Meta		MetaInfo
}

type MetaInfo struct {
	Code 		int32
	ErrorDetail	string
	ErrorType	string
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

func (resp *ExploreResponse) GetVenues() (venues []Venue) {
	for _, group := range resp.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues
}