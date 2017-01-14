package hotellook

// RoomTypes returns types rooms for the specified language.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#45
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/roomTypes.json?language=en&token=YOUR_TOKEN
func (a *HotellookApi) RoomTypes(language string) (roomTypes map[int]string, err error) {
	err = a.getJson(
		"static/roomTypes.json",
		map[string]string{
			"language": language,
		},
		&roomTypes,
	)
	return
}
