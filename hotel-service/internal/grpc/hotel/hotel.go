package serverhotel

import (
	"context"
	hotelentity "ekzamen_5/hotel-service/internal/entity/hotel"
	hotelusecase "ekzamen_5/hotel-service/internal/usecase/hotel"
	hotelpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/hotel"
	"google.golang.org/grpc"
)

type server struct {
	hotelpb.UnimplementedHotelServiceServer
	hotel *hotelusecase.HotelServiceUseCaseImpl
}

func RegisterHotelServer(grpcServer *grpc.Server, hotel *hotelusecase.HotelServiceUseCaseImpl) {
	hotelpb.RegisterHotelServiceServer(grpcServer, &server{
		hotel: hotel,
	})
}

func (s *server) CreateHotelService(ctx context.Context, req *hotelpb.CreateHotelServiceReq) (*hotelpb.CreateHotelServiceRes, error) {
	hotel, err := s.hotel.CreateHotel(ctx, &hotelentity.CreateHotelServiceReq{
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
		Address:  req.Address,
	})
	if err != nil {
		return nil, err
	}

	return &hotelpb.CreateHotelServiceRes{
		Hotel: &hotelpb.Hotel{
			HotelId:  hotel.Hotel.HotelID,
			Name:     hotel.Hotel.Name,
			Location: hotel.Hotel.Location,
			Rating:   hotel.Hotel.Rating,
			Address:  hotel.Hotel.Address,
		},
	}, nil
}

func (s *server) GetHotels(ctx context.Context, req *hotelpb.GetAllHotelsReq) (*hotelpb.GetAllHotelRes, error) {
	hotels, err := s.hotel.GetHotels(ctx, &hotelentity.GetAllHotelsReq{
		Field:  req.Field,
		Value:  req.Value,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}
	var res []*hotelpb.Hotel
	for _, h := range hotels.Hotels {
		hotel := &hotelpb.Hotel{
			HotelId:  h.HotelID,
			Name:     h.Name,
			Rating:   h.Rating,
			Address:  h.Address,
			Location: h.Location,
		}
		for _, r := range h.Rooms {
			hotel.Rooms = append(hotel.Rooms, &hotelpb.Room{
				RoomId:        r.RoomID,
				RoomType:      r.RoomType,
				PricePerNight: r.PricePerNight,
				Availability:  r.Availability,
				RoomIsBusy:    r.RoomIsBusy,
			})
		}
		res = append(res, hotel)
	}

	return &hotelpb.GetAllHotelRes{
		Hotels: res,
		Count:  hotels.Count,
	}, nil
}

func (s *server) GetHotelsRooms(ctx context.Context, req *hotelpb.GetHotelsRoomsReq) (*hotelpb.GetHotelsRoomsRes, error) {
	rooms, err := s.hotel.GetHotelsRooms(ctx, &hotelentity.GetHotelsRoomsReq{
		Field:   req.Field,
		Value:   req.Value,
		Limit:   req.Limit,
		Offset:  req.Offset,
		HotelID: req.HotelId,
	})
	if err != nil {
		return nil, err
	}
	var res []*hotelpb.Room
	for _, r := range rooms.Rooms {
		res = append(res, &hotelpb.Room{
			RoomId:        r.RoomID,
			RoomType:      r.RoomType,
			PricePerNight: r.PricePerNight,
			Availability:  r.Availability,
			RoomIsBusy:    r.RoomIsBusy,
		})
	}
	return &hotelpb.GetHotelsRoomsRes{
		Rooms: res,
		Count: rooms.Count,
	}, nil
}

func (s *server) AddRoomToHotels(ctx context.Context, req *hotelpb.AddRoomToHotelsReq) (*hotelpb.AddRoomToHotelsRes, error) {
	res, err := s.hotel.AddRoomToHotels(ctx, &hotelentity.AddRoomToHotelsReq{
		HotelID:       req.HotelId,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Availability:  req.Availability,
	})
	if err != nil {
		return nil, err
	}

	return &hotelpb.AddRoomToHotelsRes{
		Room: &hotelpb.Room{
			RoomId:        res.Room.RoomID,
			RoomType:      res.Room.RoomType,
			PricePerNight: res.Room.PricePerNight,
			Availability:  res.Room.Availability,
			RoomIsBusy:    res.Room.RoomIsBusy,
		},
	}, nil
}

func (s *server) UpdateHotel(ctx context.Context, req *hotelpb.UpdateHotelReq) (*hotelpb.UpdateHotelRes, error) {
	res, err := s.hotel.UpdateHotel(ctx, &hotelentity.UpdateHotelReq{
		HotelID:  req.HotelId,
		Name:     req.Name,
		Location: req.Location,
		Rating:   req.Rating,
		Address:  req.Address,
	})
	if err != nil {
		return nil, err
	}
	return &hotelpb.UpdateHotelRes{
		Hotel: &hotelpb.Hotel{
			HotelId:  res.Hotel.HotelID,
			Name:     res.Hotel.Name,
			Location: res.Hotel.Location,
			Rating:   res.Hotel.Rating,
			Address:  res.Hotel.Address,
		},
	}, nil
}

func (s *server) UpdateHotelsRoom(ctx context.Context, req *hotelpb.UpdateHotelsRoomReq) (*hotelpb.UpdateHotelsRoomRes, error) {
	res, err := s.hotel.UpdateHotelsRoom(ctx, &hotelentity.UpdateHotelsRoomReq{
		HotelID:       req.HotelId,
		RoomType:      req.RoomType,
		PricePerNight: req.PricePerNight,
		Availability:  req.Availability,
		RoomID:        req.RoomId,
	})
	if err != nil {
		return nil, err
	}
	return &hotelpb.UpdateHotelsRoomRes{
		Room: &hotelpb.Room{
			RoomId:        res.Room.RoomID,
			RoomType:      res.Room.RoomType,
			PricePerNight: res.Room.PricePerNight,
			Availability:  res.Room.Availability,
			RoomIsBusy:    res.Room.RoomIsBusy,
		},
	}, nil

}

func (s *server) BookingRoom(ctx context.Context, req *hotelpb.BookingRoomReq) (*hotelpb.BookingRoomRes, error) {
	res, err := s.hotel.BookingRoom(ctx, &hotelentity.BookingRoomReq{
		HotelID: req.HotelId,
		RoomID:  req.RoomId,
		Book:    req.Book,
	})
	if err != nil {
		return nil, err
	}
	return &hotelpb.BookingRoomRes{
		Room: &hotelpb.Room{
			RoomId:        res.Room.RoomID,
			RoomType:      res.Room.RoomType,
			PricePerNight: res.Room.PricePerNight,
			Availability:  res.Room.Availability,
			RoomIsBusy:    res.Room.RoomIsBusy,
		},
	}, nil
}

func (s *server) DeleteHotel(ctx context.Context, req *hotelpb.DeleteHotelReq) (*hotelpb.DeleteHotelRes, error) {
	res, err := s.hotel.DeleteHotel(ctx, &hotelentity.DeleteHotelReq{
		HotelID: req.HotelId,
	})
	if err != nil {
		return nil, err
	}
	return &hotelpb.DeleteHotelRes{
		Status: res.Status,
	}, nil
}

func (s *server) DeleteHotelRooms(ctx context.Context, req *hotelpb.DeleteHotelRoomsReq) (*hotelpb.DeleteHotelRoomsRes, error) {
	res, err := s.hotel.DeleteHotelRooms(ctx, &hotelentity.DeleteHotelRoomsReq{
		HotelID: req.HotelId,
		RoomID:  req.RoomId,
	})
	if err != nil {
		return nil, err
	}
	return &hotelpb.DeleteHotelRoomsRes{
		Status: res.Status,
	}, nil
}
