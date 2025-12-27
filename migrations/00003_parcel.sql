-- +goose up
-- +goose StatementBegin 
CREATE TABLE IF NOT EXISTS parcel(
    id BIGSERIAL PRIMARY KEY, 
    from_station VARCHAR(255) NOT NULL, 
    to_station VARCHAR(255) NOT NULL,  
    pickup_address TEXT NOT NULL, 
    delivery_address TEXT NOT NULL, 
    contact TEXT,   
    special_notes TEXT,  
    weight NUMERIC, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
) 
-- +goose StatementEnd

-- +goose down 
-- +goose StatementBegin
DROP TABLE parcel;  
-- +goose statementEnd