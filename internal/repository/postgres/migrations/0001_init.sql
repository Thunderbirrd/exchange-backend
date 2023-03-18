-- +migrate Up
CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    surname       varchar(255) not null,
    login         varchar(255) not null unique,
    password_hash varchar(255) not null,
    rating        real
);

-- +migrate Down
DROP TABLE users;