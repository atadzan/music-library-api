CREATE TABLE IF NOT EXISTS song_groups (
    id serial primary key,
    title varchar(255) not null
);

CREATE TABLE IF NOT EXISTS songs (
    id bigserial primary key,
    group_id int references song_groups(id),
    title varchar(255) not null,
    release_date date not null,
    text text default '',
    link varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp
);