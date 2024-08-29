package hotelusecase

import (
	"context"
	hotelentity "ekzamen_5/hotel-service/internal/entity/hotel"
)

type hotelServiceUseCase interface {
	CreateHotel(ctx context.Context, req *hotelentity.CreateHotelServiceReq) (*hotelentity.CreateHotelServiceRes, error)
	GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) (*hotelentity.GetAllHotelRes, error)
	GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) (*hotelentity.GetHotelsRoomsRes, error)
	AddRoomToHotels(ctx context.Context, req *hotelentity.AddRoomToHotelsReq) (*hotelentity.AddRoomToHotelsRes, error)
	UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.UpdateHotelRes, error)
	UpdateHotelsRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.UpdateHotelsRoomRes, error)
	BookingRoom(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.BookingRoomRes, error)
	DeleteHotel(ctx context.Context, req *hotelentity.DeleteHotelReq) (*hotelentity.DeleteHotelRes, error)
	DeleteHotelRooms(ctx context.Context, req *hotelentity.DeleteHotelRoomsReq) (*hotelentity.DeleteHotelRoomsRes, error)
}

type HotelServiceUseCaseImpl struct {
	hotel hotelServiceUseCase
}

func NewHotelServiceUseCase(h hotelServiceUseCase) *HotelServiceUseCaseImpl {
	return &HotelServiceUseCaseImpl{h}
}

func (h *HotelServiceUseCaseImpl) CreateHotel(ctx context.Context, req *hotelentity.CreateHotelServiceReq) (*hotelentity.CreateHotelServiceRes, error) {
	return h.hotel.CreateHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) (*hotelentity.GetAllHotelRes, error) {
	return h.hotel.GetHotels(ctx, req)
}

func (h *HotelServiceUseCaseImpl) GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) (*hotelentity.GetHotelsRoomsRes, error) {
	return h.hotel.GetHotelsRooms(ctx, req)
}

func (h *HotelServiceUseCaseImpl) AddRoomToHotels(ctx context.Context, req *hotelentity.AddRoomToHotelsReq) (*hotelentity.AddRoomToHotelsRes, error) {
	return h.hotel.AddRoomToHotels(ctx, req)
}

func (h *HotelServiceUseCaseImpl) UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.UpdateHotelRes, error) {
	return h.hotel.UpdateHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) UpdateHotelsRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.UpdateHotelsRoomRes, error) {
	return h.hotel.UpdateHotelsRoom(ctx, req)
}

func (h *HotelServiceUseCaseImpl) BookingRoom(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.BookingRoomRes, error) {
	return h.hotel.BookingRoom(ctx, req)
}
func (h *HotelServiceUseCaseImpl) DeleteHotel(ctx context.Context, req *hotelentity.DeleteHotelReq) (*hotelentity.DeleteHotelRes, error) {
	return h.hotel.DeleteHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) DeleteHotelRooms(ctx context.Context, req *hotelentity.DeleteHotelRoomsReq) (*hotelentity.DeleteHotelRoomsRes, error) {
	return h.hotel.DeleteHotelRooms(ctx, req)
}
