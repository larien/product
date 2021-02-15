CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products
(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    price_in_cents INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    created_at TIMESTAMP DEFAULT(now()),
    updated_at TIMESTAMP DEFAULT(now()),
    deleted_at TIMESTAMP
);

INSERT INTO products (price_in_cents, title, description)
VALUES
(1000, 'Produto 1', 'Descrição do produto 1'),
(2000, 'Produto 2', 'Descrição do produto 2'),
(3000, 'Produto 3', 'Descrição do produto 3'),
(4000, 'Produto 4', 'Descrição do produto 4');