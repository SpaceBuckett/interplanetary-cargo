-- +goose up
-- +goose StatementBegin 
CREATE TABLE IF NOT EXISTS payload(
    id BIGSERIAL PRIMARY KEY, 
    fromStation VARCHAR(255) NOT NULL, 
    toStation VARCHAR(255) NOT NULL,  
    bio TEXT, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,  
) 
-- +goose StatementEnd

-- +goose down 
-- +goose StatementBegin
DROP TABLE payload;  
-- +goose statementEnd