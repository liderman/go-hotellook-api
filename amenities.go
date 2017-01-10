package hotellook

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
func (a *HotellookApi) Amenities() (amenities []Amenity, err error) {
	err = a.getJson("static/amenities.json", map[string]string{}, &amenities)
	return
}
