package handler

import (
	"apigateway/internal/entity"
	hotelusecase "apigateway/internal/usecase/hotel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HotelService struct {
	hotel *hotelusecase.HotelServiceUseCaseImpl
}

func NewHotelService(hotel *hotelusecase.HotelServiceUseCaseImpl) *HotelService {
	return &HotelService{
		hotel: hotel,
	}
}

// CreateHotel godoc
// @Summary CreateHotel
// @Description Register a new Hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param body body entity.CreateHotelServiceReq true "User registration information"
// @Security ApiKeyAuth
// @Success 201 {object} entity.Hotel
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/create [post]
func (h *HotelService) CreateHotel(c *gin.Context) {
	var req entity.CreateHotelServiceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.hotel.CreateHotel(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res.Hotel)
}

// GetHotels godoc
// @Summary GetHotels
// @Description getting hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param field header string false "Field"
// @Param value header string false "Value"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Security ApiKeyAuth
// @Success 201 {object} entity.GetAllHotelRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/get [get]
func (h *HotelService) GetHotels(c *gin.Context) {
	var req entity.GetAllHotelsReq

	req.Field = c.GetHeader("Field")

	req.Value = c.GetHeader("Value")

	offsetStr := c.Query("offset")
	if offsetStr == "" {
		req.Offset = 0
	} else {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		req.Offset = int32(offset)
	}
	limitStr := c.Query("limit")
	if limitStr == "" {
		req.Limit = 0
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		req.Limit = int32(limit)
	}

	hotels, err := h.hotel.GetHotels(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotels)
}

// UpdateHotel godoc
// @Summary UpdateHotel
// @Description UpdateHotel hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param body body entity.UpdateHotelReq true " "
// @Success 201 {object} entity.Hotel
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/update [patch]
func (h *HotelService) UpdateHotel(c *gin.Context) {
	var req entity.UpdateHotelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hotel, err := h.hotel.UpdateHotel(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotel.Hotel)
}

// DeleteHotel godoc
// @Summary DeleteHotel
// @Description DeleteHotel hotel
// @Tags hotel
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/delete/{hotel_id} [delete]
func (h *HotelService) DeleteHotel(c *gin.Context) {
	var req entity.DeleteHotelReq
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	req.HotelID = hotelId
	status, err := h.hotel.DeleteHotel(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status.Status)
}

// AddRoomToHotels godoc
// @Summary AddRoomToHotels
// @Description AddRoomToHotels hotel
// @Tags room
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param body body entity.AddRoomToHotelsReq true " "
// @Security ApiKeyAuth
// @Success 201 {object} entity.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/room/create/{hotel_id} [post]
func (h *HotelService) AddRoomToHotels(c *gin.Context) {
	var req entity.AddRoomToHotelsReq
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	req.HotelID = hotelId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.hotel.AddRoomToHotels(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Room)
}

// GetHotelsRooms godoc
// @Summary GetHotelsRooms
// @Description GetHotelsRooms hotel
// @Tags room
// @Accept json
// @Produce json
// @Param field header string false "Field"
// @Param value header string false "Value"
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param hotel_id path string true "hotel_id"
// @Security ApiKeyAuth
// @Success 201 {object} entity.GetHotelsRoomsRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/room/gets/{hotel_id} [get]
func (h *HotelService) GetHotelsRooms(c *gin.Context) {
	var req entity.GetHotelsRoomsReq
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	req.HotelID = hotelId

	req.Field = c.GetHeader("Field")

	req.Value = c.GetHeader("Value")

	offsetStr := c.Query("offset")
	if offsetStr == "" {
		req.Offset = 0
	} else {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
		req.Offset = int32(offset)
	}
	limitStr := c.Query("limit")
	if limitStr == "" {
		req.Limit = 0
	} else {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
		req.Limit = int32(limit)
	}
	rooms, err := h.hotel.GetHotelsRooms(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

// UpdateHotelsRooms godoc
// @Summary UpdateHotelsRooms
// @Description UpdateHotelsRooms hotel
// @Tags room
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param room_id path string true "room_id"
// @Param body body entity.UpdateHotelsRoomReq true " "
// @Security ApiKeyAuth
// @Success 201 {object} entity.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/room/update/{hotel_id}/{room_id} [patch]
func (h *HotelService) UpdateHotelsRooms(c *gin.Context) {
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	roomId := c.Param("room_id")
	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id is required"})
		return
	}
	var req entity.UpdateHotelsRoomReq
	req.HotelID = hotelId
	req.RoomID = roomId
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.hotel.UpdateHotelsRoom(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, room.Room)
}

// DeleteHotelsRooms godoc
// @Summary DeleteHotelsRooms
// @Description DeleteHotelsRooms hotel
// @Tags room
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param room_id path string true "room_id"
// @Security ApiKeyAuth
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/room/delete/{hotel_id}/{room_id} [delete]
func (h *HotelService) DeleteHotelsRooms(c *gin.Context) {
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	roomId := c.Param("room_id")
	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id is required"})
		return
	}
	status, err := h.hotel.DeleteHotelRooms(c.Request.Context(), &entity.DeleteHotelRoomsReq{
		HotelID: hotelId,
		RoomID:  roomId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status.Status)
}

// UpdateStatusBook godoc
// @Summary UpdateStatusBook
// @Description UpdateStatusBook hotel
// @Tags room
// @Accept json
// @Produce json
// @Param hotel_id path string true "hotel_id"
// @Param room_id path string true "room_id"
// @Security ApiKeyAuth
// @Param body body entity.BookingRoomReq true " "
// @Success 201 {object} entity.Room
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /hotel/room/book/{hotel_id}/{room_id} [put]
func (h *HotelService) UpdateStatusBook(c *gin.Context) {
	hotelId := c.Param("hotel_id")
	if hotelId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id is required"})
		return
	}
	roomId := c.Param("room_id")
	if roomId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id is required"})
		return
	}

	req := &entity.BookingRoomReq{ // Инициализация структуры
		HotelID: hotelId,
		RoomID:  roomId,
	}

	if err := c.ShouldBindJSON(req); err != nil { // Передаем указатель на структуру
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.hotel.BookingRoom(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room.Room)
}
