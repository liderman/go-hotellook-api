package hotellook

import "fmt"

type HotelsResponse struct {
	GenTimestamp float64 `json:"gen_timestamp" bson:"gen_timestamp"`
	Hotels       []Hotel `json:"hotels" bson:"hotels"`
	Pois         []Poi   `json:"pois" bson:"pois"`
}

type Hotel struct {
	Address          map[string]string `json:"address" bson:"address"`
	CheckIn          string            `json:"checkIn" bson:"checkIn"`
	CheckOut         string            `json:"checkOut" bson:"checkOut"`
	CityId           int               `json:"cityId" bson:"cityId"`
	CntFloors        int               `json:"cntFloors" bson:"cntFloors"`
	CntRooms         int               `json:"cntRooms" bson:"cntRooms"`
	CntSuites        int               `json:"cntSuites" bson:"cntSuites"`
	Distance         int               `json:"distance" bson:"distance"`
	Facilities       []int             `json:"facilities" bson:"facilities"`
	Id               int               `json:"id" bson:"id"`
	Link             string            `json:"link" bson:"link"`
	Location         Coordinates       `json:"location" bson:"location"`
	Name             map[string]string `json:"name" bson:"name"`
	PhotoCount       int               `json:"photoCount" bson:"photoCount"`
	Photos           []Photo           `json:"photos" bson:"photos"`
	PhotosByRoomType map[int]int       `json:"photosByRoomType" bson:"photosByRoomType"`
	PoiDistance      map[int]int       `json:"poi_distance" bson:"poi_distance"`
	Popularity       int               `json:"popularity" bson:"popularity"`
	Pricefrom        int               `json:"pricefrom" bson:"pricefrom"`
	PropertyType     int               `json:"propertyType" bson:"propertyType"`
	Rating           int               `json:"rating" bson:"rating"`
	ShortFacilities  []string          `json:"shortFacilities" bson:"shortFacilities"`
	Stars            int               `json:"stars" bson:"stars"`
	YearOpened       int               `json:"yearOpened" bson:"yearOpened"`
	YearRenovated    int               `json:"yearRenovated" bson:"yearRenovated"`
}

type Poi struct {
	Category  string      `json:"category" bson:"category"`
	Confirmed bool        `json:"confirmed" bson:"confirmed"`
	Geom      Geom        `json:"geom" bson:"geom"`
	Id        int         `json:"id" bson:"id"`
	Location  Coordinates `json:"location" bson:"location"`
	Name      string      `json:"name" bson:"name"`
	Rating    int         `json:"rating" bson:"rating"`
}

type Geom struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

type Photo struct {
	Url    string `json:"url" bson:"url"`
	Width  int    `json:"width" bson:"width"`
	Height int    `json:"height" bson:"height"`
}

type Coordinates struct {
	Lon float64 `json:"lon" bson:"lon"`
	Lan float64 `json:"lat" bson:"lat"`
}

// Hotels returns a list of hotels by locationId.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#44
//
// Example URI:
// http://engine.hotellook.com/api/v2/static/hotels.json?locationId=895&token=YOUR_TOKEN
func (a *HotellookApi) Hotels(locationId int) (hotels HotelsResponse, err error) {
	err = a.getJson(
		"static/hotels.json",
		map[string]string{
			"locationId": fmt.Sprintf("%d", locationId),
		},
		&hotels,
	)
	return
}
