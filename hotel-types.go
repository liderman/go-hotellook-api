package hotellook

// HotelTypes returns types hotels for the specified language.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#46
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/hotelTypes.json?language=ru&token=YOUR_TOKEN
func (a *HotellookApi) HotelTypes(language string) (hotelTypes map[int]string, err error) {
	err = a.getJson(
		"static/hotelTypes.json",
		map[string]string{
			"language": language,
		},
		&hotelTypes,
	)
	return
}
