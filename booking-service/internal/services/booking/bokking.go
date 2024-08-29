package bookingservices

import (
	"context"
	bookingentity "ekzamen_5/booking-service/internal/entity/booking"
	clientgrpcserver "ekzamen_5/booking-service/internal/infastructure/client_grpc_server"
	bookingusecase "ekzamen_5/booking-service/internal/usecase/booking"
	hotelpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/hotel"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type BookingService struct {
	logger      *slog.Logger
	grpcClients clientgrpcserver.ServiceClient
	booking     *bookingusecase.BookingRepoUseCase
}

func NewBookingService(logger *slog.Logger,
	grpcClient clientgrpcserver.ServiceClient,
	booking *bookingusecase.BookingRepoUseCase) *BookingService {
	return &BookingService{
		logger:      logger,
		grpcClients: grpcClient,
		booking:     booking,
	}
}

func (b *BookingService) CreateBooking(ctx context.Context, req *bookingentity.CreateBookingReq) (*bookingentity.CreateBookingRes, error) {
	const op = "booking.CreateBooking"
	log := b.logger.With("method", op)
	log.Info("Starting CreateBooking")
	defer log.Info("Ending CreateBooking")
	res, err := b.grpcClients.HotelServiceClient().GetHotelsRooms(ctx, &hotelpb.GetHotelsRoomsReq{
		HotelId: req.HotelID,
		Field:   "room_type",
		Value:   req.RoomType,
	})
	if err != nil {
		return nil, err
	}
	for _, r := range res.Rooms {
		if !r.RoomIsBusy {
			resBooking, err := b.grpcClients.HotelServiceClient().BookingRoom(ctx, &hotelpb.BookingRoomReq{
				HotelId: req.HotelID,
				RoomId:  r.RoomId,
				Book:    true,
			})
			if err != nil {
				log.Error("err", err)
				return nil, err
			}
			err = b.booking.SaveBooking(ctx, &bookingentity.Booking{
				BookingID:    uuid.NewString(),
				UserID:       req.UserID,
				HotelID:      req.HotelID,
				RoomType:     resBooking.Room.RoomType,
				CheckInDate:  time.Now().Format("2006-01-02 15:04:05"),
				CheckOutDate: req.CheckOutDate,
				TotalAmount:  req.TotalAmount,
				Status:       "busy",
			})
			if err != nil {
				log.Error("err", err)
				return nil, err
			}

			return &bookingentity.CreateBookingRes{
				Status: "Booking successfully created",
			}, nil
		}
	}
	return &bookingentity.CreateBookingRes{
		Status:  "",
		IsError: true,
	}, nil
}

func (b *BookingService) GetBookingAll(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error) {
	const op = "booking.GetBooking"
	log := b.logger.With("method", op)
	log.Info("Starting GetBooking")
	defer log.Info("Ending GetBooking")

	bookings, err := b.booking.GetAllBooking(ctx, req)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}

	return bookings, nil
}

func (b *BookingService) UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) (*bookingentity.UpdateBookingRes, error) {
	const op = "booking.UpdateBooking"
	log := b.logger.With("method", op)
	log.Info("Starting UpdateBooking")
	defer log.Info("Ending UpdateBooking")
	err := b.booking.UpdateBooking(ctx, req)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &bookingentity.UpdateBookingRes{
		Status: "update booking successfully",
	}, nil
}

func (b *BookingService) DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) (*bookingentity.DeleteBookingRes, error) {
	const op = "booking.DeleteBooking"
	log := b.logger.With("method", op)
	log.Info("Starting DeleteBooking")
	defer log.Info("Ending DeleteBooking")
	err := b.booking.DeleteBooking(ctx, req)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &bookingentity.DeleteBookingRes{
		Status: "delete booking successfully",
	}, nil
}

func (b *BookingService) AddUserToWaitingGroup(ctx context.Context, req *bookingentity.AddUserToWaitingGroupReq) (*bookingentity.AddUserToWaitingGroupRes, error) {
	const op = "booking.AddUserToWaitingGroup"
	log := b.logger.With("method", op)
	log.Info("Starting AddUserToWaitingGroup")
	defer log.Info("Ending AddUserToWaitingGroup")
	wait := bookingentity.WaitingGroup{
		WaitGroupId:     uuid.NewString(),
		UserID:          req.UserID,
		HotelID:         req.HotelID,
		RoomType:        req.RoomType,
		TimeToStartWait: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := b.booking.AddUserToWaitingGroup(ctx, &wait)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &bookingentity.AddUserToWaitingGroupRes{
		Status: "add user successfully",
	}, nil
}

func (b *BookingService) GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error) {
	const op = "booking.GetUserToWaitingGroup"
	log := b.logger.With("method", op)
	log.Info("Starting GetUserToWaitingGroup")
	defer log.Info("Ending GetUserToWaitingGroup")

	waitList, err := b.booking.GetUserToWaitingGroup(ctx, req)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return waitList, nil
}

func (b *BookingService) GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error) {
	const op = "booking.GetAllWaitingGroup"
	log := b.logger.With("method", op)
	log.Info("Starting GetAllWaitingGroup")
	defer log.Info("Ending GetAllWaitingGroup")
	waitList, err := b.booking.GetAllWaitingGroup(ctx, req)
	if err != nil {
		log.Error("err", err)
		return bookingentity.GetAllWaitingGroupRes{}, err
	}
	return waitList, nil
}
