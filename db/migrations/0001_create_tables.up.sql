-- Director кестесі
CREATE TABLE directors (
                           id SERIAL PRIMARY KEY,
                           full_name TEXT NOT NULL,
                           created_at TEXT -- Назар аудар: Бұл жерде уақыт емес, жай STRING, қажет болса TIMESTAMP түріне өзгерту керек
);

-- Genre кестесі
CREATE TABLE genres (
                        id SERIAL PRIMARY KEY,
                        name TEXT UNIQUE NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Film кестесі
CREATE TABLE films (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       description TEXT,
                       genre_id INT REFERENCES genres(id) ON DELETE SET NULL,
                       director_id INT REFERENCES directors(id) ON DELETE SET NULL,
                       duration_minutes INT,
                       release_year INT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User кестесі
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL,  -- gorm:"column=name"
                       email TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,  -- gorm:"column=password_hash"
                       role TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
