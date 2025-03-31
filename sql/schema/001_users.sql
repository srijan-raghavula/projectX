-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    firstName TEXT NOT NULL,
    lastName TEXT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    password TEXT NOT NULL,
    region TEXT NOT NULL,
    about TEXT,
    pfpURL TEXT,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    gender CHAR(1) CHECK (gender IN ('M', 'F', 'O')),
    isPrivate BOOLEAN DEFAULT FALSE,
    followers INT DEFAULT 0,
    following INT DEFAULT 0,
    posts INT DEFAULT 0,
    site TEXT DEFAULT NULL,
    isDeleted BOOLEAN DEFAULT FALSE,
    dob DATE CHECK (dob <= CURRENT_DATE - INTERVAL '18 years')
);

-- +goose Down
DROP TABLE users;
