CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    genre VARCHAR(50) NOT NULL,
    budget INT NOT NULL
);

INSERT INTO movies (title, genre, budget) VALUES 
('SAW', 'Horror', 500000),
('TEST', 'Romance', 1000000);