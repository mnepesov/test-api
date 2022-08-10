create table if not exists photos
(
    id         serial primary key,
    apod       jsonb not null,
    created_at timestamp default now()
);
