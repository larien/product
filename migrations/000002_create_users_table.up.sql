CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    created_at TIMESTAMP DEFAULT(now()),
    updated_at TIMESTAMP DEFAULT(now()),
    deleted_at TIMESTAMP
);

INSERT INTO users (first_name, last_name, date_of_birth)
VALUES
('Nome1', 'Sobrenome1', '2000-07-12'),
('Nome2', 'Sobrenome2', '1992-12-15'),
('Nome3', 'Sobrenome3', '1996-11-25'),
('Nome4', 'Sobrenome4', '1987-02-04');