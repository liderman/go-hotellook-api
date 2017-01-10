package hotellook

type Location struct {
	Code      string                       `json:"code" bson:"code"`
	Id        int                          `json:"id,string" bson:"id"`
	CountryId int                          `json:"countryId,string" bson:"countryId"`
	Latitude  float64                      `json:"latitude,string" bson:"latitude"`
	Longitude float64                      `json:"longitude,string" bson:"longitude"`
	Name      []map[string][]VariationName `json:"name" bson:"name"`
}

// Locations returns a list of cities (locations).
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#42
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/locations.json?token=YOUR_TOKEN
func (a *HotellookApi) Locations() (locations []Location, err error) {
	err = a.getJson("static/locations.json", map[string]string{}, &locations)
	return
}
