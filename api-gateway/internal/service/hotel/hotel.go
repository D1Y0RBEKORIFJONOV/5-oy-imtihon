package hotelservice

import (
	"apigateway/internal/entity"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"context"
	"errors"
	"fmt"
	user1 "github.com/D1Y0RBEKORIFJONOV/SmartHome_Protos/gen/go/user"
	bookingpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/booking"
	hotelpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/hotel"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"log/slog"
	"sync"
)

type Hotel struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
}

func NewHotel(logger *slog.Logger, client clientgrpcserver.ServiceClient) *Hotel {
	return &Hotel{logger: logger, client: client}
}

func (h *Hotel) CreateHotel(ctx context.Context, req *entity.CreateHotelServiceReq) (*entity.CreateHotelServiceRes, error) {
	const op = "Service.CreateHotel"
	log := h.logger.With(
		slog.String("method", op))
	log.Info("Start Create hotel")
	defer log.Info("End Create hotel")
	res, err := h.client.HotelServiceClient().CreateHotelService(ctx, &hotelpb.CreateHotelServiceReq{
		Name:     req.Name,
		Address:  req.Address,
		Location: req.Location,
		Rating:   req.Rating,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.CreateHotelServiceRes{
		Hotel: entity.Hotel{
			HotelID:  res.Hotel.HotelId,
			Name:     res.Hotel.Name,
			Address:  res.Hotel.Address,
			Location: res.Hotel.Location,
			Rating:   res.Hotel.Rating,
		},
	}, nil
}

func (h *Hotel) GetHotels(ctx context.Context, req *entity.GetAllHotelsReq) (*entity.GetAllHotelRes, error) {
	const op = "Hotel.GetHotels"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start Get hotels")
	defer log.Info("End Get hotels")
	hotels, err := h.client.HotelServiceClient().GetHotels(ctx, &hotelpb.GetAllHotelsReq{
		Field:  req.Field,
		Value:  req.Value,
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var res entity.GetAllHotelRes
	for _, ho := range hotels.Hotels {
		res.Hotels = append(res.Hotels, &entity.Hotel{
			HotelID:  ho.HotelId,
			Name:     ho.Name,
			Address:  ho.Address,
			Location: ho.Location,
			Rating:   ho.Rating,
		})
	}
	res.Count = hotels.Count
	return &res, nil
}

func (h *Hotel) GetHotelsRooms(ctx context.Context, req *entity.GetHotelsRoomsReq) (*entity.GetHotelsRoomsRes, error) {
	const op = "Hotel.GetHotelsRooms"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start Get hotels rooms")
	defer log.Info("End Get hotels rooms")
	res, err := h.client.HotelServiceClient().GetHotelsRooms(ctx, &hotelpb.GetHotelsRoomsReq{
		Field:   req.Field,
		Value:   req.Value,
		Offset:  req.Offset,
		Limit:   req.Limit,
		HotelId: req.HotelID,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var response entity.GetHotelsRoomsRes
	response.Count = res.Count
	for _, r := range res.Rooms {
		response.Rooms = append(response.Rooms, &entity.Room{
			RoomID:        r.RoomId,
			RoomType:      r.RoomType,
			RoomIsBusy:    r.RoomIsBusy,
			Availability:  r.Availability,
			PricePerNight: r.PricePerNight,
		})
	}
	return &response, nil
}

func (h *Hotel) AddRoomToHotels(ctx context.Context, req *entity.AddRoomToHotelsReq) (*entity.AddRoomToHotelsRes, error) {
	const op = "Service.AddRoomToHotels"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start AddRoomToHotels")
	defer log.Info("End AddRoomToHotels")
	room, err := h.client.HotelServiceClient().AddRoomToHotels(ctx, &hotelpb.AddRoomToHotelsReq{
		HotelId:       req.HotelID,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Availability:  req.Availability,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.AddRoomToHotelsRes{
		Room: entity.Room{
			RoomType:      room.Room.RoomType,
			RoomID:        room.Room.RoomId,
			PricePerNight: room.Room.PricePerNight,
			Availability:  room.Room.Availability,
			RoomIsBusy:    room.Room.RoomIsBusy,
		},
	}, nil
}

func (h *Hotel) UpdateHotel(ctx context.Context, req *entity.UpdateHotelReq) (*entity.UpdateHotelRes, error) {
	const op = "Hotel.UpdateHotel"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start Update hotel")
	defer log.Info("End Update hotel")
	hotel, err := h.client.HotelServiceClient().UpdateHotel(ctx, &hotelpb.UpdateHotelReq{
		HotelId: req.HotelID,
		Name:    req.Name,
		Address: req.Address,
		Rating:  req.Rating,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.UpdateHotelRes{
		Hotel: entity.Hotel{
			HotelID:  hotel.Hotel.HotelId,
			Name:     hotel.Hotel.Name,
			Address:  hotel.Hotel.Address,
			Rating:   hotel.Hotel.Rating,
			Location: hotel.Hotel.Location,
		},
	}, nil
}

func (h *Hotel) UpdateHotelsRoom(ctx context.Context, req *entity.UpdateHotelsRoomReq) (*entity.UpdateHotelsRoomRes, error) {
	const op = "Hotel.UpdateHotelsRoom"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start Update hotels")
	defer log.Info("End Update hotels")

	res, err := h.client.HotelServiceClient().UpdateHotelsRoom(ctx, &hotelpb.UpdateHotelsRoomReq{
		HotelId:       req.HotelID,
		RoomId:        req.RoomID,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Availability:  req.Availability,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.UpdateHotelsRoomRes{
		Room: entity.Room{
			RoomID:        res.Room.RoomId,
			RoomType:      res.Room.RoomType,
			RoomIsBusy:    res.Room.RoomIsBusy,
			Availability:  res.Room.Availability,
			PricePerNight: res.Room.PricePerNight,
		},
	}, err
}

func (h *Hotel) BookingRoom(ctx context.Context, req *entity.BookingRoomReq) (*entity.BookingRoomRes, error) {
	const op = "Service.BookingRoom"
	log := h.logger.With(
		slog.With("method", op))
	log.Info("Starting Booking Room")
	defer log.Info("Ending Booking Room")
	if !req.Book {
		rooms, err := h.client.HotelServiceClient().GetHotelsRooms(ctx, &hotelpb.GetHotelsRoomsReq{
			HotelId: req.HotelID,
			Field:   "room_id",
			Value:   req.RoomID,
		})
		if err != nil {
			log.Error("err", err)
		}
		if rooms.Count == 0 {
			return &entity.BookingRoomRes{}, errors.New("no room found")
		}
		room := rooms.Rooms[0]
		waitListUsers, err := h.client.BookingServiceClient().GetAllWaitingGroup(ctx, &bookingpb.GetAllWaitingGroupReq{
			Field: "room_type",
			Value: room.RoomType,
		})
		if err != nil {
			log.Error("err", err)
		}
		if len(waitListUsers.WaitingGroups) != 0 {
			for _, w := range waitListUsers.WaitingGroups {
				waitGroup := sync.WaitGroup{}
				waitGroup.Add(1)
				go func() {
					_, err = h.client.NotificationServiceClient().AddNotification(ctx, &notificationpb.AddNotificationReq{
						UserId: w.UserId,
						Messages: &notificationpb.CreateMessage{
							SenderName: "BOOKING-SERVICE",
							Status: fmt.Sprintf("[booking]:Siz band qilish hohlagan hona  boshadi!\nhonani tezroq band qiling\nroom_id:%s\nhotel_id:%s",
								room.RoomId, req.HotelID),
						},
					})
					if err != nil {
						log.Error("err", err)
					}
					waitGroup.Done()
				}()
				waitGroup.Add(1)
				go func() {
					user, err := h.client.UserServiceClient().GetUser(ctx, &user1.GetUserReq{
						Filed: "id",
						Value: w.UserId,
					})
					if err != nil {
						log.Error("err", err)
					}
					_, err = h.client.NotificationServiceClient().SendEmailNotification(ctx, &notificationpb.SendEmailNotificationReq{
						Email:      user.Email,
						SenderName: "BOOKING-SERVICE",
						Notification: fmt.Sprintf("[booking]:Siz band qilish hohlagan hona  boshadi!\nhonani tezroq band qiling\nroom_id:%s\nhotel_id:%s",
							room.RoomId, req.HotelID),
					})
					if err != nil {
						log.Error("err", err)
					}
					waitGroup.Done()
				}()
				waitGroup.Wait()
			}
		}
	}
	res, err := h.client.HotelServiceClient().BookingRoom(ctx, &hotelpb.BookingRoomReq{
		HotelId: req.HotelID,
		RoomId:  req.RoomID,
		Book:    req.Book,
	})
	if err != nil {
		log.Error("err", err)
		return &entity.BookingRoomRes{}, err
	}
	return &entity.BookingRoomRes{
		Room: entity.Room{
			RoomID:        res.Room.RoomId,
			RoomType:      res.Room.RoomType,
			RoomIsBusy:    res.Room.RoomIsBusy,
			Availability:  res.Room.Availability,
			PricePerNight: res.Room.PricePerNight,
		},
	}, nil
}

func (h *Hotel) DeleteHotel(ctx context.Context, req *entity.DeleteHotelReq) (*entity.DeleteHotelRes, error) {
	const op = "Hotel.DeleteHotel"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start Delete hotel")
	defer log.Info("End Delete hotel")
	status, err := h.client.HotelServiceClient().DeleteHotel(ctx, &hotelpb.DeleteHotelReq{
		HotelId: req.HotelID,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.DeleteHotelRes{
		Status: status.Status,
	}, nil
}

func (h *Hotel) DeleteHotelRooms(ctx context.Context, req *entity.DeleteHotelRoomsReq) (*entity.DeleteHotelRoomsRes, error) {
	const op = "Service.DeleteHotelRooms"
	log := h.logger.With(slog.String("method", op))
	log.Info("Start DeleteHotel rooms")
	defer log.Info("End delete Hotel rooms")

	res, err := h.client.HotelServiceClient().DeleteHotelRooms(ctx, &hotelpb.DeleteHotelRoomsReq{
		HotelId: req.HotelID,
		RoomId:  req.RoomID,
	})
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &entity.DeleteHotelRoomsRes{
		Status: res.Status,
	}, nil
}
