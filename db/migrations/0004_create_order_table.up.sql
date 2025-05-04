CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER NOT NULL,
                        film_id INTEGER NOT NULL,
                        seat_row INTEGER NOT NULL,
                        seat_number INTEGER NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
