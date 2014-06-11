package foursquare

import "fmt"

func (v Venue) String() string { 
	return fmt.Sprintf("ID: %s\nName: %s\nLocation: %s\nCategories: %s", v.Id, v.Name, v.Location, v.Categories)
}