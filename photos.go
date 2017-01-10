package hotellook

import "fmt"

// HotelPhotoUrl generates a URL hotel photo.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#47
//
// Example photo URL:
// https://photo.hotellook.com/image_v2/limit/h418520_1/800/520.auto
func (a *HotellookApi) HotelPhotoUrl(format string, hotelId, photoId, width, height int) string {
	if format != "crop" {
		format = "limit"
	}

	size := fmt.Sprintf("%d", width)
	if height > 0 {
		size = fmt.Sprintf("%d/%d", width, height)
	}

	return fmt.Sprintf("https://photo.hotellook.com/image_v2/%s/h%d_%d/%s.auto", format, hotelId, photoId, size)
}

// RoomsSpritePhotoUrl generates url photos of rooms (in a sprite).
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#48
//
// Generate rules:
// http://photo.hotellook.com/rooms/sprite/h<hotelId>_<groupId>/<firstWidth>/<photosCount>/<restPhotosWidth>.auto
//
// Example photo URLs:
// http://photo.hotellook.com/rooms/sprite/h4_1/100/3/50.auto
// http://photo.hotellook.com/rooms/sprite/h4_12/100x50/3/50x25.auto

func (a *HotellookApi) RoomsSpritePhotoUrl(hotelId, groupId int, firstWidth string, photosCount int, restPhotosWidth string) string {
	return fmt.Sprintf(
		"http://photo.hotellook.com/rooms/sprite/h%d_%d/%s/%d/%s.auto",
		hotelId, groupId, firstWidth, photosCount, restPhotosWidth,
	)
}

// RoomsPhotoUrl generates url of the room photo of a given size.
//
// Documentation:
// https://support.travelpayouts.com/hc/en-us/articles/115000343268-Hotels-data-API#48
//
// Generate rules:
// http://photo.hotellook.com/rooms/<format>/h<hotelId>_<groupId>_<photoIdx>/<width>/<height>.auto
//
// Example photo URL:
// http://photo.hotellook.com/rooms/crop/h4_12_1/200/200.auto

func (a *HotellookApi) RoomsPhotoUrl(format string, hotelId, groupId, photoIdx, width, height int) string {
	if format != "crop" {
		format = "limit"
	}

	return fmt.Sprintf(
		"http://photo.hotellook.com/rooms/%s/h%d_%d_%d/%d/%d.auto",
		format, hotelId, groupId, photoIdx, width, height,
	)
}
