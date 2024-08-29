package hotelservices

import (
	"context"
	hotelentity "ekzamen_5/hotel-service/internal/entity/hotel"
	hotelusecase "ekzamen_5/hotel-service/internal/usecase/hotel"
	"errors"
	"github.com/google/uuid"
	"log/slog"
)

type Hotel struct {
	logger *slog.Logger
	hotel  *hotelusecase.HotelRepoUseCase
}

func NewHotelService(logger *slog.Logger, hotel *hotelusecase.HotelRepoUseCase) *Hotel {
	return &Hotel{
		logger: logger,
		hotel:  hotel,
	}
}

func (h *Hotel) CreateHotel(ctx context.Context, req *hotelentity.CreateHotelServiceReq) (*hotelentity.CreateHotelServiceRes, error) {
	const op = "Service.CreateHotel"
	log := h.logger.With(
		slog.String("method", op))
	log.Info("Start Create hotel")
	defer log.Info("End Create hotel")
	hotel := &hotelentity.Hotel{
		HotelID:  uuid.NewString(),
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
		Address:  req.Address,
		Rooms:    []hotelentity.Room{},
	}
	log.Info("Call SaverHotel ")
	err := h.hotel.SaveHotel(ctx, hotel)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.CreateHotelServiceRes{}, err
	}

	return &hotelentity.CreateHotelServiceRes{
		Hotel: *hotel,
	}, nil
}

func (h *Hotel) GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) (*hotelentity.GetAllHotelRes, error) {
	const op = "Service.GetHotels"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start get hotel service")
	defer log.Info("End get hotel service")
	hotels, err := h.hotel.GetHotels(ctx, req)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.GetAllHotelRes{}, err
	}
	return &hotelentity.GetAllHotelRes{
		Hotels: hotels,
		Count:  int64(len(hotels)),
	}, nil
}

func (h *Hotel) GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) (*hotelentity.GetHotelsRoomsRes, error) {
	const op = "Service.GetHotels"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start get room service")
	defer log.Info("End get room service")
	if req.Field == "room_type" {
		exists, err := h.hotel.IsRoomTypeExists(ctx, req.HotelID, req.Value)
		if err != nil {
			log.Error("err", err)
			return &hotelentity.GetHotelsRoomsRes{}, err
		}
		if !exists {
			log.Error("room type  does not exist")
			return &hotelentity.GetHotelsRoomsRes{}, errors.New("room type does not exist")
		}
	}
	rooms, err := h.hotel.GetHotelsRooms(ctx, req)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.GetHotelsRoomsRes{}, err
	}
	return &hotelentity.GetHotelsRoomsRes{
		Rooms: rooms,
		Count: int64(len(rooms)),
	}, nil
}

func (h *Hotel) AddRoomToHotels(ctx context.Context, req *hotelentity.AddRoomToHotelsReq) (*hotelentity.AddRoomToHotelsRes, error) {
	const op = "Service.AddRoomToHotels"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start add room to hotel service")
	defer log.Info("End add room to hotel service")
	room := hotelentity.Room{
		HotelID:       req.HotelID,
		RoomID:        uuid.NewString(),
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Availability:  req.Availability,
		RoomIsBusy:    false,
	}
	err := h.hotel.AddRoomToHotels(ctx, &room)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.AddRoomToHotelsRes{}, err
	}

	return &hotelentity.AddRoomToHotelsRes{
		Room: room,
	}, nil
}

func (h *Hotel) UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.UpdateHotelRes, error) {
	const op = "Service.UpdateHotel"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start update hotel service")
	defer log.Info("End update hotel service")
	hotel, err := h.hotel.UpdateHotel(ctx, &hotelentity.UpdateHotelReq{
		HotelID:  req.HotelID,
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
		Address:  req.Address,
	})
	if err != nil {
		log.Error("err", err)
		return &hotelentity.UpdateHotelRes{}, err
	}
	return &hotelentity.UpdateHotelRes{
		Hotel: *hotel,
	}, nil
}

func (h *Hotel) UpdateHotelsRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.UpdateHotelsRoomRes, error) {
	const op = "Service.UpdateHotel"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start update hotel room service")
	defer log.Info("End update hotel room service")
	room, err := h.hotel.UpdateHotelRoom(ctx, req)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.UpdateHotelsRoomRes{}, err
	}
	return &hotelentity.UpdateHotelsRoomRes{
		Room: *room,
	}, nil
}

func (h *Hotel) BookingRoom(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.BookingRoomRes, error) {
	const op = "Service.BookingRoom"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start booking room service")
	defer log.Info("End booking room service")
	room, err := h.hotel.UpdateBooking(ctx, req)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.BookingRoomRes{}, err
	}
	return &hotelentity.BookingRoomRes{
		Room: *room,
	}, nil
}

func (h *Hotel) DeleteHotel(ctx context.Context, req *hotelentity.DeleteHotelReq) (*hotelentity.DeleteHotelRes, error) {
	const op = "Service.DeleteHotel"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start delete hotel service")
	defer log.Info("End delete hotel service")
	err := h.hotel.DeleteHotel(ctx, req.HotelID)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.DeleteHotelRes{}, err
	}
	return &hotelentity.DeleteHotelRes{
		Status: "hotel is deleted",
	}, nil
}

func (h *Hotel) DeleteHotelRooms(ctx context.Context, req *hotelentity.DeleteHotelRoomsReq) (*hotelentity.DeleteHotelRoomsRes, error) {
	const op = "Service.DeleteHotelRooms"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start delete hotel room service")
	defer log.Info("End delete hotel room service")
	err := h.hotel.DeleteHotelRoom(ctx, req.HotelID, req.RoomID)
	if err != nil {
		log.Error("err", err)
		return &hotelentity.DeleteHotelRoomsRes{}, err
	}

	return &hotelentity.DeleteHotelRoomsRes{
		Status: "room is deleted",
	}, nil
}
