CREATE TABLE cinemas (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL,
                         address TEXT NOT NULL,
                         city TEXT NOT NULL
);
CREATE TABLE sessions (
                          id SERIAL PRIMARY KEY,
                          film_id INTEGER REFERENCES films(id) ON DELETE CASCADE,
                          cinema_id INTEGER REFERENCES cinemas(id) ON DELETE CASCADE,
                          start_time TIMESTAMP NOT NULL,
                          hall_name TEXT NOT NULL
);
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                        session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
                        seat_row INTEGER NOT NULL,
                        seat_number INTEGER NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
