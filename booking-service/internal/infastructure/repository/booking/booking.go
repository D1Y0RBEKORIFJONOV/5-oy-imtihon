package repositorybooking

import (
	"context"
	bookingentity "ekzamen_5/booking-service/internal/entity/booking"
	"ekzamen_5/booking-service/internal/postgres"
	"errors"
	"fmt"
	"log/slog"
)

type Booking struct {
	db            *postgres.PostgresDB
	tableBooking  string
	tableWaitList string
	log           *slog.Logger
}

func NewSolderRepository(db *postgres.PostgresDB, log *slog.Logger) *Booking {
	return &Booking{
		db:            db,
		tableBooking:  "booking",
		tableWaitList: "wait_list",
		log:           log,
	}
}

func (b *Booking) selectQueryBooking() string {
	return `
    booking_id,
    user_id,
    hotel_id,
    room_type,
    check_in_data,
    check_out_data,
    total_amount,
    status
    `
}

func (w *Booking) selectQueryWaitList() string {
	return `
    wait_group_id,
    user_id,
    hotel_id,
    room_type,
    time_to_start_wait
    `
}

func (h *Booking) returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (h *Booking) SaveBooking(ctx context.Context, req *bookingentity.Booking) error {
	const op = "RepositoryBooking.SaveBooking"
	log := h.log.With("method", op)
	log.Info("Starting Save Booking")
	defer log.Info("Finishing Save Booking")
	data := map[string]interface{}{
		"booking_id":     req.BookingID,
		"user_id":        req.UserID,
		"hotel_id":       req.HotelID,
		"room_type":      req.RoomType,
		"check_in_data":  req.CheckInDate,
		"check_out_data": req.CheckOutDate,
		"total_amount":   req.TotalAmount,
		"status":         req.Status,
	}
	query, args, err := h.db.Sq.Builder.Insert(h.tableBooking).SetMap(data).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Info("query", query)
		log.Error("err", err)
		return err
	}

	return nil
}

func (h *Booking) AddUserToWaitingGroup(ctx context.Context, req *bookingentity.WaitingGroup) error {
	const op = "RepositoryBooking.AddUserToWaitingGroup"
	log := h.log.With("method", op)
	log.Info("Starting AddUserToWaitingGroup")
	defer log.Info("Finishing AddUserToWaitingGroup")
	data := map[string]interface{}{
		"user_id":            req.UserID,
		"hotel_id":           req.HotelID,
		"room_type":          req.RoomType,
		"wait_group_id":      req.WaitGroupId,
		"time_to_start_wait": req.TimeToStartWait,
	}
	query, args, err := h.db.Sq.Builder.Insert(h.tableWaitList).SetMap(data).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	return nil
}

func (h *Booking) GetUserToWaitingGroup(ctx context.Context, req *bookingentity.GetUserToWaitingGroupReq) (*bookingentity.GetUserToWaitingGroupRes, error) {
	const op = "RepositoryBooking.GetUserToWaitingGroup"
	log := h.log.With("method", op)
	log.Info("Starting GetUserToWaitingGroup")
	defer log.Info("Finishing GetUserToWaitingGroup")
	var wait bookingentity.WaitingGroup
	query, args, err := h.db.Sq.Builder.Select(h.selectQueryWaitList()).
		From(h.tableWaitList).
		Where(h.db.Sq.Equal("user_id", req.UserID)).ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	err = h.db.QueryRow(ctx, query, args...).Scan(
		&wait.WaitGroupId,
		&wait.UserID,
		&wait.HotelID,
		&wait.RoomType,
		&wait.TimeToStartWait)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	return &bookingentity.GetUserToWaitingGroupRes{
		WaitingGroup: wait,
	}, nil
}

func (h *Booking) GetAllBooking(ctx context.Context, req *bookingentity.GetBookingAllReq) (*bookingentity.GetBookingAllRes, error) {
	const op = "RepositoryBooking.GetAllBooking"
	log := h.log.With("method", op)
	log.Info("Starting GetAllBooking")
	defer log.Info("Finishing GetAllBooking")
	toSql := h.db.Sq.Builder.Select(h.selectQueryBooking()).From(h.tableBooking)
	if req.Value != "" && req.Field != "" {
		toSql = toSql.Where(h.db.Sq.Equal(req.Field, req.Value))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Offset != 0 {
		toSql = toSql.Offset(uint64(req.Offset))
	}
	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	rows, err := h.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	defer rows.Close()
	var bookings []bookingentity.Booking
	for rows.Next() {
		var booking bookingentity.Booking
		err = rows.Scan(
			&booking.BookingID,
			&booking.UserID,
			&booking.HotelID,
			&booking.RoomType,
			&booking.CheckInDate,
			&booking.CheckOutDate,
			&booking.TotalAmount,
			&booking.Status)
		if err != nil {
			log.Error("err", err)
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	if err := rows.Err(); err != nil {
		log.Error("err", err)
		return nil, err
	}
	if len(bookings) == 0 {
		return nil, nil
	}
	return &bookingentity.GetBookingAllRes{
		Bookings: bookings,
		Count:    uint64(len(bookings)),
	}, nil
}

func (h *Booking) GetAllWaitingGroup(ctx context.Context, req *bookingentity.GetAllWaitingGroupReq) (bookingentity.GetAllWaitingGroupRes, error) {
	const op = "RepositoryBooking.GetAllWaitingGroup"
	log := h.log.With("method", op)
	log.Info("Starting GetAllWaitingGroup")
	defer log.Info("Finishing GetAllWaitingGroup")
	toSql := h.db.Sq.Builder.Select(h.selectQueryWaitList()).From(h.tableWaitList)
	if req.Value != "" && req.Field != "" {
		toSql = toSql.Where(h.db.Sq.Equal(req.Field, req.Value))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Offset != 0 {
		toSql = toSql.Offset(uint64(req.Offset))
	}
	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error("err", err)
		return bookingentity.GetAllWaitingGroupRes{}, err
	}
	rows, err := h.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return bookingentity.GetAllWaitingGroupRes{}, err
	}
	defer rows.Close()

	var waitGroups []*bookingentity.WaitingGroup
	for rows.Next() {
		var waitGroup bookingentity.WaitingGroup
		err = rows.Scan(&waitGroup.WaitGroupId,
			&waitGroup.UserID,
			&waitGroup.HotelID,
			&waitGroup.RoomType,
			&waitGroup.TimeToStartWait)
		if err != nil {
			log.Error("err", err)
			return bookingentity.GetAllWaitingGroupRes{}, err
		}

		waitGroups = append(waitGroups, &waitGroup)
	}
	if err := rows.Err(); err != nil {
		log.Error("err", err)
		return bookingentity.GetAllWaitingGroupRes{}, err
	}
	if len(waitGroups) == 0 {
		return bookingentity.GetAllWaitingGroupRes{}, nil
	}
	return bookingentity.GetAllWaitingGroupRes{
		WaitingGroups: waitGroups,
	}, nil
}

func (h *Booking) UpdateBooking(ctx context.Context, req *bookingentity.UpdateBookingReq) error {
	const op = "RepositoryBooking.UpdateBooking"
	log := h.log.With("method", op)
	log.Info("Starting UpdateBooking")
	defer log.Info("Finishing UpdateBooking")
	data := map[string]interface{}{}
	if req.RoomType != "" {
		data["room_type"] = req.RoomType
	}
	if req.CheckInDate != "" {
		data["check_in_data"] = req.CheckInDate
	}
	if req.CheckOutDate != "" {
		data["check_out_data"] = req.CheckOutDate
	}
	if req.TotalAmount != "" {
		data["total_amount"] = req.TotalAmount
	}

	query, args, err := h.db.Sq.Builder.Update(h.tableBooking).SetMap(data).
		Where(h.db.Sq.Equal("booking_id", req.BookingID)).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	result, err := h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	if result.RowsAffected() == 0 {
		log.Error("err", errors.New("booking not found"))
		return errors.New("booking not found")
	}
	return nil
}

func (h *Booking) DeleteBooking(ctx context.Context, req *bookingentity.DeleteBookingReq) error {
	const op = "RepositoryBooking.DeleteBooking"
	log := h.log.With("method", op)
	log.Info("Starting DeleteBooking")
	defer log.Info("Finishing DeleteBooking")
	query, args, err := h.db.Sq.Builder.Delete(h.tableBooking).
		Where(h.db.Sq.Equal("booking_id", req.BookingID)).ToSql()
	if err != nil {
		log.Error("err", err)
		return err
	}
	result, err := h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("err", err)
		return err
	}
	if result.RowsAffected() == 0 {
		log.Error("err", errors.New("booking not found"))
		return errors.New("booking not found")
	}
	return nil
}
