CREATE TABLE users (
                       id bigserial not null primary key,
                       name varchar(32) not null,
                       surname varchar(32) not null,
                       email varchar not null,
                       encrypted_password varchar not null,
                       card_id int not null,
                       is_worker bool not null
);