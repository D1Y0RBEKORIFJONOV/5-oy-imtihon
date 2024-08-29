package postgresql

import (
	"context"
	hotelentity "ekzamen_5/hotel-service/internal/entity/hotel"
	"ekzamen_5/hotel-service/internal/postgres"
	"fmt"
	"log/slog"
)

type Hotel struct {
	db            *postgres.PostgresDB
	tableName     string
	tableRoomName string
	log           *slog.Logger
}

func NewSolderRepository(db *postgres.PostgresDB, log *slog.Logger) *Hotel {
	return &Hotel{
		db:            db,
		tableName:     "hotels",
		tableRoomName: "rooms",
		log:           log,
	}
}

func (h *Hotel) selectQueryRoom() string {
	return `
	room_id,
	price_per_night,
	avail_lability,
	room_is_busy,
	room_type
`
}

func (h Hotel) selectQueryHotel() string {
	return `
	hotel_id,
	name,
	rating,
	address,
	location
`
}

func (h *Hotel) returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (h *Hotel) SaveHotel(ctx context.Context, req *hotelentity.Hotel) error {
	const op = "Repository.SaveHotel"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting Save hotel")
	defer log.Info("Ending Save hotel")

	data := map[string]interface{}{
		"name":     req.Name,
		"rating":   req.Rating,
		"address":  req.Address,
		"hotel_id": req.HotelID,
		"location": req.Location,
	}
	query, args, err := h.db.Sq.Builder.Insert(h.tableName).
		SetMap(data).ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (h *Hotel) AddRoomToHotels(ctx context.Context, req *hotelentity.Room) error {
	const op = "Repository.AddRoomToHotel"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting AddRoomToHotel")
	defer log.Info("Ending AddRoomToHotel")
	data := map[string]interface{}{
		"room_id":         req.RoomID,
		"hotel_id":        req.HotelID,
		"price_per_night": req.PricePerNight,
		"avail_lability":  req.Availability,
		"room_is_busy":    req.RoomIsBusy,
		"room_type":       req.RoomType,
	}
	query, args, err := h.db.Sq.Builder.Insert(h.tableRoomName).SetMap(data).ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (h *Hotel) GetHotels(ctx context.Context, req *hotelentity.GetAllHotelsReq) ([]*hotelentity.Hotel, error) {
	const op = "Repository.GetHotels"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting GetHotels")
	defer log.Info("Ending GetHotels")
	toSql := h.db.Sq.Builder.Select(h.selectQueryHotel()).From(h.tableName)
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
		log.Error(err.Error())
		return nil, err
	}
	rows, err := h.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	hotels := make([]*hotelentity.Hotel, 0)
	for rows.Next() {
		hotel := &hotelentity.Hotel{}
		err = rows.Scan(
			&hotel.HotelID,
			&hotel.Name,
			&hotel.Rating,
			&hotel.Address,
			&hotel.Location)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		hotels = append(hotels, hotel)
	}
	if err := rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return hotels, nil
}

func (h *Hotel) GetHotelsRooms(ctx context.Context, req *hotelentity.GetHotelsRoomsReq) ([]*hotelentity.Room, error) {
	const op = "Repository.GetHotelsRooms"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting GetHotelsRooms")
	defer log.Info("Ending GetHotelsRooms")
	toSql := h.db.Sq.Builder.Select(h.selectQueryRoom()).From(h.tableRoomName)
	if req.Value != "" && req.Field != "" {
		toSql = toSql.Where(h.db.Sq.And(h.db.Sq.Equal(req.Field, req.Value), h.db.Sq.Equal("hotel_id", req.HotelID)))
	} else {
		toSql = toSql.Where(h.db.Sq.Equal("hotel_id", req.HotelID))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Offset != 0 {
		toSql = toSql.Offset(uint64(req.Offset))
	}
	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	row, err := h.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
	}
	defer row.Close()
	rooms := make([]*hotelentity.Room, 0)
	for row.Next() {
		room := &hotelentity.Room{}
		err = row.Scan(
			&room.RoomID,
			&room.PricePerNight,
			&room.Availability,
			&room.RoomIsBusy,
			&room.RoomType)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err := row.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return rooms, nil
}

func (h *Hotel) UpdateHotel(ctx context.Context, req *hotelentity.UpdateHotelReq) (*hotelentity.Hotel, error) {
	const op = "Repository.UpdateHotel"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting UpdateHotel")
	defer log.Info("Ending UpdateHotel")
	data := map[string]interface{}{}
	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Location != "" {
		data["location"] = req.Location
	}
	if req.Rating != "" {
		data["rating"] = req.Rating
	}
	if req.Address != "" {
		data["address"] = req.Address
	}
	query, args, err := h.db.Sq.Builder.Update(h.tableName).
		SetMap(data).
		Where(h.db.Sq.Equal("hotel_id", req.HotelID)).
		Suffix(h.returning(h.selectQueryHotel())).ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var hotel hotelentity.Hotel
	err = h.db.QueryRow(ctx, query, args...).Scan(&hotel.HotelID,
		&hotel.Name,
		&hotel.Rating,
		&hotel.Address,
		&hotel.Location)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &hotel, nil
}

func (h *Hotel) UpdateHotelRoom(ctx context.Context, req *hotelentity.UpdateHotelsRoomReq) (*hotelentity.Room, error) {
	const op = "Repository.UpdateHotelRoom"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting UpdateHotelRoom")
	defer log.Info("Ending UpdateHotelRoom")
	data := map[string]interface{}{}

	if req.RoomType != "" {
		data["room_type"] = req.RoomType
	}
	if req.Availability != "" {
		data["avail_lability"] = req.Availability
	}
	if req.PricePerNight != "" {
		data["price_per_night"] = req.PricePerNight
	}
	query, args, err := h.db.Sq.Builder.Update(h.tableRoomName).SetMap(data).
		Where(h.db.Sq.And(h.db.Sq.Equal("room_id", req.RoomID),
			h.db.Sq.Equal("hotel_id", req.HotelID))).
		Suffix(h.returning(h.selectQueryRoom())).ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var room hotelentity.Room
	err = h.db.QueryRow(ctx, query, args...).Scan(&room.RoomID,
		&room.PricePerNight,
		&room.Availability,
		&room.RoomIsBusy,
		&room.RoomType)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &room, nil
}

func (h *Hotel) UpdateBooking(ctx context.Context, req *hotelentity.BookingRoomReq) (*hotelentity.Room, error) {
	const op = "Repository.UpdateBooking"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting UpdateBooking")
	defer log.Info("Ending UpdateBooking")
	data := map[string]interface{}{
		"room_is_busy": req.Book,
	}
	query, args, err := h.db.Sq.Builder.Update(h.tableRoomName).SetMap(data).
		Where(h.db.Sq.And(h.db.Sq.Equal("room_id", req.RoomID),
			h.db.Sq.Equal("hotel_id", req.HotelID))).
		Suffix(h.returning(h.selectQueryRoom())).ToSql()
	if err != nil {
		log.Error("err", err)
		return nil, err
	}
	var room hotelentity.Room
	err = h.db.QueryRow(ctx, query, args...).Scan(&room.RoomID,
		&room.PricePerNight,
		&room.Availability,
		&room.RoomIsBusy,
		&room.RoomType)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &room, nil
}

func (h *Hotel) DeleteHotel(ctx context.Context, id string) error {
	const op = "Repository.DeleteHotel"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting DeleteHotel")
	defer log.Info("Ending DeleteHotel")

	query, args, err := h.db.Sq.Builder.Delete(h.tableName).
		Where(h.db.Sq.Equal("hotel_id", id)).ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (h *Hotel) DeleteHotelRoom(ctx context.Context, hotelId, roomID string) error {
	const op = "Repository.DeleteHotelRoom"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting DeleteHotelRoom")
	defer log.Info("Ending DeleteHotelRoom")

	query, args, err := h.db.Sq.Builder.Delete(h.tableRoomName).
		Where(h.db.Sq.And(
			h.db.Sq.Equal("room_id", roomID),
			h.db.Sq.Equal("hotel_id", hotelId),
		)).
		ToSql()

	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = h.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (h *Hotel) IsRoomTypeExists(ctx context.Context, hotelId, roomType string) (bool, error) {
	const op = "Repository.IsRoomTypeExists"
	log := h.log.With(slog.String("operation", op))
	log.Info("Starting IsRoomTypeExists")
	var exists bool

	err := h.db.QueryRow(ctx, `
	SELECT EXISTS(
		SELECT 1 
		FROM rooms 
		WHERE hotel_id = $1 AND room_type = $2
	)`, hotelId, roomType).Scan(&exists)

	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	return exists, nil
}
