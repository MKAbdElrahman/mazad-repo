CREATE TABLE IF NOT EXISTS auctions (
    id SERIAL PRIMARY KEY,
    item_name VARCHAR(255) NOT NULL,
    item_description TEXT,
    item_category VARCHAR(255),
    item_manufacturer VARCHAR(255),
    item_condition VARCHAR(255),
    item_images TEXT ARRAY,
    reserve_price INT,
    current_high_bid INT,
    seller VARCHAR(255),
    winner VARCHAR(255),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp,
    end_at TIMESTAMP,
    status INT
);

