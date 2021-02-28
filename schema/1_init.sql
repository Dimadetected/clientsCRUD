CREATE TABLE sales
(
    id    bigint PRIMARY KEY NOT NULL UNIQUE,
    count bigint             not null,
    total float              not null
);

CREATE TABLE clients
(
    id    bigint primary key not null unique,
    name  varchar(255)       not null,
    email varchar(255)       not null unique,
    age   int                not null
);

CREATE TABLE users
(
    id       bigint primary key not null unique,
    name     varchar(255)       not null,
    username varchar(255)       not null,
    password text               not null
);

CREATE TABLE clients_sales
(
    id        bigint primary key                            not null unique,
    client_id int references clients (id) on delete cascade not null,
    sale_id   int references sales (id) on delete cascade   not null
)


