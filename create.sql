
CREATE SCHEMA IF NOT EXISTS quortle;

CREATE TABLE IF NOT EXISTS quortle.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



--Tracks user overall scores, 
CREATE TABLE IF NOT EXISTS quortle.scores (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total_guesses INTEGER DEFAULT 0,
    score INTEGER NOT NULL,
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

--Contains all of the 4 letter words that can be used in the game
CREATE TABLE IF NOT EXISTS quortle.words(
    id SERIAL PRIMARY KEY,
    word VARCHAR(4) NOT NULL UNIQUE
)

--Contains the days word, and the date it is active for
CREATE TABLE IF NOT EXISTS quortle.game (
    id SERIAL PRIMARY KEY,
    word_id INTEGER NOT NULL,
    date_from DATE NOT NULL UNIQUE,
    date_to DATE NOT NULL UNIQUE,

    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE
);