package bookingserver

import (
	"context"
	bookingentity "ekzamen_5/booking-service/internal/entity/booking"
	bookingusecase "ekzamen_5/booking-service/internal/usecase/booking"
	bookingpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/booking"
	"google.golang.org/grpc"
)

type server struct {
	bookingpb.UnimplementedBookingServiceServer
	booking *bookingusecase.BookingUseCase
}

func RegisterBookingServer(grpcServer *grpc.Server, bookingUseCase *bookingusecase.BookingUseCase) {
	bookingpb.RegisterBookingServiceServer(grpcServer, &server{booking: bookingUseCase})
}

func (s *server) CreateBooking(ctx context.Context, req *bookingpb.CreateBookingReq) (*bookingpb.CreateBookingRes, error) {
	status, err := s.booking.CreateBooking(ctx, &bookingentity.CreateBookingReq{
		UserID:       req.UserID,
		HotelID:      req.HotelID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
	})
	if err != nil {
		return nil, err
	}

	return &bookingpb.CreateBookingRes{
		Status:  status.Status,
		IsError: status.IsError,
	}, nil
}

func (s *server) GetBookingAll(ctx context.Context, req *bookingpb.GetBookingAllReq) (*bookingpb.GetBookingAllRes, error) {
	books, err := s.booking.GetBookingAll(ctx, &bookingentity.GetBookingAllReq{
		Field:  req.Field,
		Limit:  req.Limit,
		Offset: req.Offset,
		Value:  req.Value,
	})
	if books == nil {
		return &bookingpb.GetBookingAllRes{}, nil
	}
	if err != nil {
		return nil, err
	}
	var res bookingpb.GetBookingAllRes
	for _, b := range books.Bookings {
		res.Bookings = append(res.Bookings, &bookingpb.Booking{
			BookingID:    b.BookingID,
			UserID:       b.UserID,
			HotelID:      b.HotelID,
			RoomType:     b.RoomType,
			CheckInDate:  b.CheckInDate,
			CheckOutDate: b.CheckOutDate,
			TotalAmount:  b.TotalAmount,
			Status:       b.Status,
		})
	}
	res.Count = books.Count
	return &res, nil
}

func (s *server) UpdateBooking(ctx context.Context, req *bookingpb.UpdateBookingReq) (*bookingpb.UpdateBookingRes, error) {
	book, err := s.booking.UpdateBooking(ctx, &bookingentity.UpdateBookingReq{
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       req.Status,
		BookingID:    req.BookingId,
	})
	if err != nil {
		return nil, err
	}

	return &bookingpb.UpdateBookingRes{
		Status: book.Status,
	}, nil
}

func (s *server) DeleteBooking(ctx context.Context, req *bookingpb.DeleteBookingReq) (*bookingpb.DeleteBookingRes, error) {
	status, err := s.booking.DeleteBooking(ctx, &bookingentity.DeleteBookingReq{
		BookingID: req.BookingId,
	})
	if err != nil {
		return nil, err
	}
	return &bookingpb.DeleteBookingRes{
		Status: status.Status,
	}, nil
}

func (s *server) AddUserToWaitingGroup(ctx context.Context, req *bookingpb.AddUserToWaitingGroupReq) (*bookingpb.AddUserToWaitingGroupRes, error) {
	status, err := s.booking.AddUserToWaitingGroup(ctx, &bookingentity.AddUserToWaitingGroupReq{
		UserID:   req.UserId,
		HotelID:  req.HotelID,
		RoomType: req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	return &bookingpb.AddUserToWaitingGroupRes{
		Status: status.Status,
	}, nil
}

func (s *server) GetUserTOWaitingGroup(ctx context.Context, req *bookingpb.GetUserTOWaitingGroupReq) (*bookingpb.WaitingGroup, error) {
	weitGroup, err := s.booking.GetUserToWaitingGroup(ctx, &bookingentity.GetUserToWaitingGroupReq{
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &bookingpb.WaitingGroup{
		WaitId:             weitGroup.WaitingGroup.WaitGroupId,
		UserId:             weitGroup.WaitingGroup.UserID,
		HotelId:            weitGroup.WaitingGroup.HotelID,
		RoomType:           weitGroup.WaitingGroup.RoomType,
		TimeToStartWaiting: weitGroup.WaitingGroup.TimeToStartWait,
	}, nil
}

func (s *server) GetAllWaitingGroup(ctx context.Context, req *bookingpb.GetAllWaitingGroupReq) (*bookingpb.GetAllWaitingGroupRes, error) {
	waitsGroups, err := s.booking.GetAllWaitingGroup(ctx, &bookingentity.GetAllWaitingGroupReq{
		Field:  req.Field,
		Limit:  req.Limit,
		Offset: req.Offset,
		Value:  req.Value,
	})

	if err != nil {
		return nil, err
	}
	var res bookingpb.GetAllWaitingGroupRes
	for _, w := range waitsGroups.WaitingGroups {
		res.WaitingGroups = append(res.WaitingGroups, &bookingpb.WaitingGroup{
			WaitId:             w.WaitGroupId,
			UserId:             w.UserID,
			HotelId:            w.HotelID,
			RoomType:           w.RoomType,
			TimeToStartWaiting: w.TimeToStartWait,
		})
	}
	return &res, nil
}
