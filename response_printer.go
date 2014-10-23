package foursquare

import "fmt"

func (v Venue) String() string {
	return fmt.Sprintf("ID: %s\nName: %s\nLocation: %s\nCategories: %s\nVerified: %t\nStats: %d", v.Id, v.Name, v.Location, v.Categories, v.Verified, v.Stats)
}
