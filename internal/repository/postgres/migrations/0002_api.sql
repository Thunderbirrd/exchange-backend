-- +migrate Up
CREATE TABLE airports
(
    code             varchar(8) PRIMARY KEY,
    country          varchar(64),
    city             varchar(64),
    machine_location varchar(512)
);

CREATE TABLE currencies
(
    code        varchar(8) PRIMARY KEY,
    rate_to_usd real
);

CREATE TABLE requests
(
    id            serial PRIMARY KEY,
    author_id     int references users (id) on delete cascade               not null,
    from_currency varchar(8) references currencies (code) on delete cascade not null,
    to_currency   varchar(8) references currencies (code) on delete cascade not null,
    value_from    real                                                      not null,
    value_to      real                                                      not null,
    date_time     timestamp                                                 not null,
    airport       varchar(8) references airports (code) on delete cascade   not null
);

-- +migrate Down
DROP TABLE currencies;
DROP TABLE airports;
DROP TABLE requests;