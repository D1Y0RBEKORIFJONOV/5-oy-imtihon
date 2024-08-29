package bookingservice

import (
	"apigateway/internal/config"
	bookingentity "apigateway/internal/entity/booking"
	clientgrpcserver "apigateway/internal/infastructure/client_grpc_server"
	"apigateway/internal/infastructure/producer"
	"context"
	"encoding/json"
	bookingpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/booking"
	"log/slog"
)

type Booking struct {
	logger *slog.Logger
	client clientgrpcserver.ServiceClient
	cfg    *config.Config
}

func NewBooking(logger *slog.Logger, client clientgrpcserver.ServiceClient) *Booking {
	cfg := config.New()
	return &Booking{
		logger: logger,
		client: client,
		cfg:    cfg,
	}
}

func (b *Booking) CreateBooking(ctx context.Context, req *bookingentity.CreateBookingReq) (*bookingentity.CreateBookingRes, error) {
	const op = "Service.CreateBooking"
	log := b.logger.With(
		slog.String("method", op))
	log.Info("Create booking started")
	defer log.Info("Create booking ended")
	req1 := bookingpb.CreateBookingReq{
		UserID:       req.UserID,
		HotelID:      req.HotelID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
	}
	reqBytes, err := json.Marshal(&req1)
	if err != nil {
		log.Error("err", err.Error())
		return nil, err
	}
	err = producer.Producer(b.cfg.MessageBrokerUses.KeysBooking.CreateOrder, reqBytes, b.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Error("err", err.Error())
		return nil, err
	}
	return &bookingentity.CreateBookingRes{
		Status: "creating order processing",
	}, nil
}

func (b *Booking) GetBookingAll(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error) {
	panic("implement me")
}

func (b *Booking) UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) (*bookingentity.UpdateBookingRes, error) {
	panic("implement me")
}

func (b *Booking) DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) (*bookingentity.DeleteBookingRes, error) {
	panic("implement me")
}

func (b *Booking) AddUserToWaitingGroup(ctx context.Context, req *bookingentity.AddUserToWaitingGroupReq) (*bookingentity.AddUserToWaitingGroupRes, error) {
	const op = "Booking.AddUserToWaitingGroup"
	log := b.logger.With(
		slog.String("method", op))
	log.Info("AddUserToWaitingGroup started")
	defer log.Info("AddUserToWaitingGroup ended")
	req1 := bookingpb.AddUserToWaitingGroupReq{
		UserId:   req.UserID,
		HotelID:  req.HotelID,
		RoomType: req.RoomType,
	}
	reqBytes, err := json.Marshal(&req1)
	if err != nil {
		log.Error("err", err.Error())
		return nil, err
	}
	err = producer.Producer(b.cfg.MessageBrokerUses.KeysBooking.AddWaitGroup, reqBytes, b.cfg.MessageBrokerUses.Topic)
	if err != nil {
		log.Error("err", err.Error())
		return nil, err
	}
	return &bookingentity.AddUserToWaitingGroupRes{
		Status: "adding user to waiting group processing",
	}, nil
}

func (b *Booking) GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error) {
	panic("implement me")
}

func (b *Booking) GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error) {
	panic("implement me")
}
