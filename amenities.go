package hotellook

import "fmt"

type Amenity struct {
	Id        int    `json:"id,string" bson:"id"`
	Name      string `json:"name" bson:"name"`
	GroupName string `json:"groupName" bson:"groupName"`
}

// Amenities returns a list of amenities.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#43
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/amenities.json?token=YOUR_TOKEN
func (a *HotellookApi) Amenities(language string) (amenities []Amenity, err error) {
	url := "static/amenities.json"
	if language != "en" {
		url = fmt.Sprintf("static/amenities/%s.json", language)
	}

	err = a.getJson(url, map[string]string{}, &amenities)
	return
}
