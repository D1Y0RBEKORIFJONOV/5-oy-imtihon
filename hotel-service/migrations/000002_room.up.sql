CREATE TABLE IF NOT EXISTS rooms (
    room_id UUID PRIMARY KEY,
    hotel_id UUID REFERENCES hotels(hotel_id) ON DELETE CASCADE,
    price_per_night VARCHAR(255),
    avail_lability VARCHAR(255),
    room_type VARCHAR(255),
    room_is_busy BOOLEAN
);
