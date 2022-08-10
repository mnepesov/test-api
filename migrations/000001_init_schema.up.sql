create table if not exists pictures
(
    id         serial primary key,
    apod       jsonb not null,
    created_at date default CURRENT_DATE
);
