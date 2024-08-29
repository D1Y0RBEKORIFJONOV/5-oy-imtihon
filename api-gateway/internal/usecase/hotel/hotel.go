package hotelusecase

import (
	"apigateway/internal/entity"
	"context"
)

type hotelServiceUseCase interface {
	CreateHotel(ctx context.Context, req *entity.CreateHotelServiceReq) (*entity.CreateHotelServiceRes, error)
	GetHotels(ctx context.Context, req *entity.GetAllHotelsReq) (*entity.GetAllHotelRes, error)
	GetHotelsRooms(ctx context.Context, req *entity.GetHotelsRoomsReq) (*entity.GetHotelsRoomsRes, error)
	AddRoomToHotels(ctx context.Context, req *entity.AddRoomToHotelsReq) (*entity.AddRoomToHotelsRes, error)
	UpdateHotel(ctx context.Context, req *entity.UpdateHotelReq) (*entity.UpdateHotelRes, error)
	UpdateHotelsRoom(ctx context.Context, req *entity.UpdateHotelsRoomReq) (*entity.UpdateHotelsRoomRes, error)
	BookingRoom(ctx context.Context, req *entity.BookingRoomReq) (*entity.BookingRoomRes, error)
	DeleteHotel(ctx context.Context, req *entity.DeleteHotelReq) (*entity.DeleteHotelRes, error)
	DeleteHotelRooms(ctx context.Context, req *entity.DeleteHotelRoomsReq) (*entity.DeleteHotelRoomsRes, error)
}

type HotelServiceUseCaseImpl struct {
	hotel hotelServiceUseCase
}

func NewHotelServiceUseCase(h hotelServiceUseCase) *HotelServiceUseCaseImpl {
	return &HotelServiceUseCaseImpl{h}
}

func (h *HotelServiceUseCaseImpl) CreateHotel(ctx context.Context, req *entity.CreateHotelServiceReq) (*entity.CreateHotelServiceRes, error) {
	return h.hotel.CreateHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) GetHotels(ctx context.Context, req *entity.GetAllHotelsReq) (*entity.GetAllHotelRes, error) {
	return h.hotel.GetHotels(ctx, req)
}

func (h *HotelServiceUseCaseImpl) GetHotelsRooms(ctx context.Context, req *entity.GetHotelsRoomsReq) (*entity.GetHotelsRoomsRes, error) {
	return h.hotel.GetHotelsRooms(ctx, req)
}

func (h *HotelServiceUseCaseImpl) AddRoomToHotels(ctx context.Context, req *entity.AddRoomToHotelsReq) (*entity.AddRoomToHotelsRes, error) {
	return h.hotel.AddRoomToHotels(ctx, req)
}

func (h *HotelServiceUseCaseImpl) UpdateHotel(ctx context.Context, req *entity.UpdateHotelReq) (*entity.UpdateHotelRes, error) {
	return h.hotel.UpdateHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) UpdateHotelsRoom(ctx context.Context, req *entity.UpdateHotelsRoomReq) (*entity.UpdateHotelsRoomRes, error) {
	return h.hotel.UpdateHotelsRoom(ctx, req)
}

func (h *HotelServiceUseCaseImpl) BookingRoom(ctx context.Context, req *entity.BookingRoomReq) (*entity.BookingRoomRes, error) {
	return h.hotel.BookingRoom(ctx, req)
}
func (h *HotelServiceUseCaseImpl) DeleteHotel(ctx context.Context, req *entity.DeleteHotelReq) (*entity.DeleteHotelRes, error) {
	return h.hotel.DeleteHotel(ctx, req)
}

func (h *HotelServiceUseCaseImpl) DeleteHotelRooms(ctx context.Context, req *entity.DeleteHotelRoomsReq) (*entity.DeleteHotelRoomsRes, error) {
	return h.hotel.DeleteHotelRooms(ctx, req)
}
