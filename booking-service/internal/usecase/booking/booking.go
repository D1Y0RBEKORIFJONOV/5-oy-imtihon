package bookingusecase

import (
	"context"
	bookingentity "ekzamen_5/booking-service/internal/entity/booking"
)

type bookingService interface {
	CreateBooking(ctx context.Context, req *bookingentity.CreateBookingReq) (*bookingentity.CreateBookingRes, error)
	GetBookingAll(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error)
	UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) (*bookingentity.UpdateBookingRes, error)
	DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) (*bookingentity.DeleteBookingRes, error)
	AddUserToWaitingGroup(ctx context.Context, req *bookingentity.AddUserToWaitingGroupReq) (*bookingentity.AddUserToWaitingGroupRes, error)
	GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error)
	GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error)
}

type BookingUseCase struct {
	bookingService bookingService
}

func NewBookingUseCase(bookingService bookingService) *BookingUseCase {
	return &BookingUseCase{
		bookingService: bookingService,
	}
}

func (b *BookingUseCase) CreateBooking(ctx context.Context, req *bookingentity.CreateBookingReq) (*bookingentity.CreateBookingRes, error) {
	return b.bookingService.CreateBooking(ctx, req)
}

func (b *BookingUseCase) GetBookingAll(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error) {
	return b.bookingService.GetBookingAll(ctx, req)
}

func (b *BookingUseCase) UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) (*bookingentity.UpdateBookingRes, error) {
	return b.bookingService.UpdateBooking(ctx, req)
}

func (b *BookingUseCase) DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) (*bookingentity.DeleteBookingRes, error) {
	return b.bookingService.DeleteBooking(ctx, req)
}

func (b *BookingUseCase) AddUserToWaitingGroup(ctx context.Context, req *bookingentity.AddUserToWaitingGroupReq) (*bookingentity.AddUserToWaitingGroupRes, error) {
	return b.bookingService.AddUserToWaitingGroup(ctx, req)
}

func (b *BookingUseCase) GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error) {
	return b.bookingService.GetUserToWaitingGroup(ctx, req)
}

func (b *BookingUseCase) GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error) {
	return b.bookingService.GetAllWaitingGroup(ctx, req)
}
