package hotelusecase

import (
	"context"
	hotelentity "ekzamen_5/hotel-service/internal/entity/hotel"
)

type (
	saver interface {
		SaveHotel(ctx context.Context, req *hotelentity.Hotel) error
		AddRoomToHotels(ctx context.Context, req *hotelentity.Room) error
	}
	provider interface {
		GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) ([]*hotelentity.Hotel, error)
		GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) ([]*hotelentity.Room, error)
		IsRoomTypeExists(ctx context.Context, hotelId, roomType string) (bool, error)
	}
	updater interface {
		UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.Hotel, error)
		UpdateHotelRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.Room, error)
		UpdateBooking(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.Room, error)
	}
	deleter interface {
		DeleteHotel(ctx context.Context, id string) error
		DeleteHotelRoom(ctx context.Context, hotelId, roomID string) error
	}
)

type hotelRepoUseCase interface {
	provider
	updater
	deleter
	saver
}

type HotelRepoUseCase struct {
	hotelRepo hotelRepoUseCase
}

func NewHotelRepoUseCase(h hotelRepoUseCase) *HotelRepoUseCase {
	return &HotelRepoUseCase{hotelRepo: h}
}

func (h *HotelRepoUseCase) SaveHotel(ctx context.Context, req *hotelentity.Hotel) error {
	return h.hotelRepo.SaveHotel(ctx, req)
}

func (h *HotelRepoUseCase) AddRoomToHotels(ctx context.Context, req *hotelentity.Room) error {
	return h.hotelRepo.AddRoomToHotels(ctx, req)
}

func (h *HotelRepoUseCase) GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) ([]*hotelentity.Hotel, error) {
	return h.hotelRepo.GetHotels(ctx, req)
}

func (h *HotelRepoUseCase) GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) ([]*hotelentity.Room, error) {
	return h.hotelRepo.GetHotelsRooms(ctx, req)
}

func (h *HotelRepoUseCase) UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.Hotel, error) {
	return h.hotelRepo.UpdateHotel(ctx, req)
}

func (h *HotelRepoUseCase) UpdateHotelRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.Room, error) {
	return h.hotelRepo.UpdateHotelRoom(ctx, req)
}

func (h *HotelRepoUseCase) UpdateBooking(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.Room, error) {
	return h.hotelRepo.UpdateBooking(ctx, req)
}

func (h *HotelRepoUseCase) DeleteHotel(ctx context.Context, id string) error {
	return h.hotelRepo.DeleteHotel(ctx, id)
}

func (h *HotelRepoUseCase) DeleteHotelRoom(ctx context.Context, hotelId, roomID string) error {
	return h.hotelRepo.DeleteHotelRoom(ctx, hotelId, roomID)
}

func (h *HotelRepoUseCase) IsRoomTypeExists(ctx context.Context, hotelId, roomType string) (bool, error) {
	return h.hotelRepo.IsRoomTypeExists(ctx, hotelId, roomType)
}
