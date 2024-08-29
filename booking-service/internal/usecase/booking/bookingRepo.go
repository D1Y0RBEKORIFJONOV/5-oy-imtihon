package bookingusecase

import (
	"context"
	bookingentity "ekzamen_5/booking-service/internal/entity/booking"
)

type (
	saver interface {
		SaveBooking(ctx context.Context, req *bookingentity.Booking) error
		AddUserToWaitingGroup(ctx context.Context, req *bookingentity.WaitingGroup) error
	}
	provider interface {
		GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error)
		GetAllBooking(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error)
		GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error)
	}
	updater interface {
		UpdateBooking(context.Context, *bookingentity.UpdateBookingReq) error
	}
	deleter interface {
		DeleteBooking(context.Context, *bookingentity.DeleteBookingReq) error
	}
)

type bookingRepoUseCase interface {
	saver
	provider
	updater
	deleter
}

type BookingRepoUseCase struct {
	booking bookingRepoUseCase
}

func NewBookingRepoUseCase(booking bookingRepoUseCase) *BookingRepoUseCase {
	return &BookingRepoUseCase{
		booking: booking,
	}
}

func (v *BookingRepoUseCase) SaveBooking(ctx context.Context, req *bookingentity.Booking) error {
	return v.booking.SaveBooking(ctx, req)
}

func (v *BookingRepoUseCase) AddUserToWaitingGroup(ctx context.Context, req *bookingentity.WaitingGroup) error {
	return v.booking.AddUserToWaitingGroup(ctx, req)
}

func (v *BookingRepoUseCase) GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error) {
	return v.booking.GetUserToWaitingGroup(ctx, req)
}

func (v *BookingRepoUseCase) GetAllBooking(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error) {
	return v.booking.GetAllBooking(ctx, req)
}

func (v *BookingRepoUseCase) GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error) {
	return v.booking.GetAllWaitingGroup(ctx, req)
}

func (v *BookingRepoUseCase) UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) error {
	return v.booking.UpdateBooking(ctx, req)
}

func (v *BookingRepoUseCase) DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) error {
	return v.booking.DeleteBooking(ctx, req)
}
