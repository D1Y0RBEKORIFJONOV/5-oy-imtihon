package handler

import (
	bookingentity "apigateway/internal/entity/booking"
	bookingusecase "apigateway/internal/usecase/booking"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Booking struct {
	booking *bookingusecase.BookingUseCase
}

func NewBooking(booking *bookingusecase.BookingUseCase) *Booking {
	return &Booking{
		booking: booking,
	}
}

// CreateBooking godoc
// @Summary CreateBooking
// @Description CreateBooking hotel
// @Tags booking
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param room_type path string true "room_type"
// @Param body body bookingentity.CreateBookingReq true " "
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /booking/create/{hotel_id}/{room_type} [post]
func (b *Booking) CreateBooking(c *gin.Context) {
	var req bookingentity.CreateBookingReq
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	req.UserID = id.(string)
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	req.HotelID = hotelId

	roomType := c.Param("room_type")
	if roomType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_type is required"})
		return
	}
	req.RoomType = roomType

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := b.booking.CreateBooking(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Status)
}

// AddUserToWaitList godoc
// @Summary AddUserToWaitList
// @Description AddUserToWaitList hotel
// @Tags booking
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param room_type path string true "room_type"
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /booking/add/waitlist/{hotel_id}/{room_type} [post]
func (b *Booking) AddUserToWaitList(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}

	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}

	roomType := c.Param("room_type")
	if roomType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_type is required"})
		return
	}
	idStr, ok := id.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	res, err := b.booking.AddUserToWaitingGroup(c.Request.Context(), &bookingentity.AddUserToWaitingGroupReq{
		UserID:   idStr,
		HotelID:  hotelId,
		RoomType: roomType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Status)
}
