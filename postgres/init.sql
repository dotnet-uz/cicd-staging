CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO users (name) VALUES
('Ali Valiyev'),
('Sardor Karimov'),
('Jasur Abdullayev'),
('Bekzod Tursunov'),
('Dilshod Rahimov'),
('Aziza Islomova'),
('Madina Rasulova'),
('Nodira Yusupova');