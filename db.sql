create table users (
    id serial primary key,
    name varchar,
    lastname varchar,
    created_at timestamp default current_timestamp
)