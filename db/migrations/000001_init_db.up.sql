CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    username  VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    password  VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books
(
    id     SERIAL PRIMARY KEY,
    author varchar(255) NOT NULL,
    title  varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    order_id SERIAL PRIMARY KEY,
    user_id  INTEGER   NOT NULL,
    book_id  INTEGER   NOT NULL,
    take     TIMESTAMP NOT NULL,
    return   TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (book_id) REFERENCES books (id)
);