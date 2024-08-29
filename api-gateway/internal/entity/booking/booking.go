package bookingentity

type (
	Booking struct {
		BookingID    string `json:"booking_id"`
		UserID       string `json:"user_id"`
		HotelID      string `json:"hotel_id"`
		RoomType     string `json:"room_type"`
		CheckInDate  string `json:"check_in_date"`
		CheckOutDate string `json:"check_out_date"`
		TotalAmount  string `json:"total_amount"`
		Status       string `json:"status"`
	}

	CreateBookingReq struct {
		UserID       string `json:"-"`
		HotelID      string `json:"-"`
		RoomType     string `json:"-"`
		CheckInDate  string `json:"check_in_date"`
		CheckOutDate string `json:"check_out_date"`
		TotalAmount  string `json:"total_amount"`
	}

	CreateBookingRes struct {
		Status string `json:"status"`
	}

	GetBookingAllReq struct {
		Field  string `json:"field"`
		Value  string `json:"value"`
		Offset int64  `json:"offset"`
		Limit  int64  `json:"limit"`
	}

	GetBookingAllRes struct {
		Bookings []Booking `json:"bookings"`
		Count    uint64    `json:"count"`
	}

	UpdateBookingReq struct {
		BookingID    string `json:"booking_id"`
		RoomType     string `json:"room_type"`
		CheckInDate  string `json:"check_in_date"`
		CheckOutDate string `json:"check_out_date"`
		TotalAmount  string `json:"total_amount"`
		Status       string `json:"status"`
	}

	UpdateBookingRes struct {
		Status string `json:"status"`
	}

	DeleteBookingReq struct {
		BookingID string `json:"booking_id"`
	}

	DeleteBookingRes struct {
		Status string `json:"status"`
	}

	AddUserToWaitingGroupReq struct {
		UserID   string `json:"user_id"`
		HotelID  string `json:"hotel_id"`
		RoomType string `json:"room_type"`
	}

	AddUserToWaitingGroupRes struct {
		Status string `json:"status"`
	}

	GetUserToWaitingGroupReq struct {
		UserID string `json:"user_id"`
	}

	WaitingGroup struct {
		WaitGroupId     string `json:"wait_group_id"`
		UserID          string `json:"user_id"`
		HotelID         string `json:"hotel_id"`
		RoomType        string `json:"room_type"`
		TimeToStartWait string `json:"time_to_start_waiting"`
	}

	GetUserToWaitingGroupRes struct {
		WaitingGroup WaitingGroup `json:"waiting_group"`
	}

	GetAllWaitingGroupReq struct {
		Field  string `json:"field"`
		Value  string `json:"value"`
		Offset int64  `json:"offset"`
		Limit  int64  `json:"limit"`
	}

	GetAllWaitingGroupRes struct {
		WaitingGroups []*WaitingGroup `json:"waiting_groups"`
	}
)
