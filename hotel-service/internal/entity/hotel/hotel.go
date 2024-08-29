package hotelentity

type (
	Hotel struct {
		HotelID  string `json:"hotel_id"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Rating   string `json:"rating"`
		Address  string `json:"address"`
		Rooms    []Room `json:"rooms"`
	}

	Room struct {
		RoomID        string `json:"room_id"`
		RoomType      string `json:"room_type"`
		PricePerNight string `json:"price_per_night"`
		Availability  string `json:"availability"`
		RoomIsBusy    bool   `json:"room_is_busy"`
		HotelID       string `json:"hotel_id"`
	}

	CreateHotelServiceReq struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Rating   string `json:"rating"`
		Address  string `json:"address"`
	}

	CreateHotelServiceRes struct {
		Hotel Hotel `json:"hotel"`
	}

	GetAllHotelsReq struct {
		Field  string `json:"field"`
		Value  string `json:"value"`
		Offset int32  `json:"offset"`
		Limit  int32  `json:"limit"`
	}

	GetAllHotelRes struct {
		Hotels []*Hotel `json:"hotels"`
		Count  int64    `json:"count"`
	}

	GetHotelsRoomsReq struct {
		HotelID string `json:"hotel_id"`
		Field   string `json:"field"`
		Value   string `json:"value"`
		Offset  int32  `json:"offset"`
		Limit   int32  `json:"limit"`
	}

	GetHotelsRoomsRes struct {
		Rooms []*Room `json:"rooms"`
		Count int64   `json:"count"`
	}

	AddRoomToHotelsReq struct {
		HotelID       string `json:"hotel_id"`
		RoomType      string `json:"room_type"`
		PricePerNight string `json:"price_per_night"`
		Availability  string `json:"availability"`
	}

	AddRoomToHotelsRes struct {
		Room Room `json:"room"`
	}

	UpdateHotelReq struct {
		HotelID  string `json:"hotel_id"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Rating   string `json:"rating"`
		Address  string `json:"address"`
	}

	UpdateHotelRes struct {
		Hotel Hotel `json:"hotel"`
	}

	UpdateHotelsRoomReq struct {
		RoomID        string `json:"room_id"`
		HotelID       string `json:"hotel_id"`
		RoomType      string `json:"room_type"`
		PricePerNight string `json:"price_per_night"`
		Availability  string `json:"availability"`
	}

	UpdateHotelsRoomRes struct {
		Room Room `json:"room"`
	}

	BookingRoomReq struct {
		HotelID string `json:"hotel_id"`
		RoomID  string `json:"room_id"`
		Book    bool   `json:"book"`
	}

	BookingRoomRes struct {
		Room Room `json:"room"`
	}

	DeleteHotelReq struct {
		HotelID string `json:"hotel_id"`
	}

	DeleteHotelRes struct {
		Status string `json:"status"`
	}

	DeleteHotelRoomsReq struct {
		HotelID string `json:"hotel_id"`
		RoomID  string `json:"room_id"`
	}

	DeleteHotelRoomsRes struct {
		Status string `json:"status"`
	}
)
