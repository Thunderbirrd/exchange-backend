-- +migrate Up
CREATE TABLE exchanges
(
    id               serial PRIMARY KEY,
    author_id        int references users (id) on delete cascade    not null,
    acceptor_id      int references users (id) on delete cascade    not null,
    request_id       int references requests (id) on delete cascade not null,
    author_code      varchar(16),
    acceptor_code    varchar(16),
    author_approve   bool,
    acceptor_approve bool,
    status           varchar(16)                                    not null,
    expired_time     timestamp                                      not null
);

-- +migrate Down
DROP TABLE exchanges;