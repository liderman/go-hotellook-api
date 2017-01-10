package hotellook

type Country struct {
	Code string                       `json:"code" bson:"code"`
	Id   int                          `json:"id,string" bson:"id"`
	Name []map[string][]VariationName `json:"name" bson:"name"`
}

// Countries returns a list of countries.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#41
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/countries.json?token=YOUR_TOKEN
func (a *HotellookApi) Countries() (countries []Country, err error) {
	err = a.getJson("static/countries.json", map[string]string{}, &countries)
	return
}
