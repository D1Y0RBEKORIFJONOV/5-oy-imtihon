CREATE TABLE IF NOT EXISTS wait_list(
    wait_group_id UUID PRIMARY KEY,
    user_id VARCHAR(255),
    hotel_id UUID,
    room_type VARCHAR(255),
    time_to_start_wait  VARCHAR(255)
);


