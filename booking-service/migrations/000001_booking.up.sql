CREATE  TABLE IF NOT EXISTS booking(
    booking_id UUID PRIMARY KEY,
    user_id VARCHAR(255),
    hotel_id UUID,
    room_type VARCHAR(255),
    check_in_data VARCHAR(255),
    check_out_data VARCHAR(255),
    total_amount VARCHAR(255),
    status VARCHAR(255)
);


