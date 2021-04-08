create table tg_user (
    id bigint not null primary key,
    first_name varchar,
    last_name varchar,
    username varchar,
    language_code varchar,
    is_bot boolean,
    can_join_groups boolean,
    can_read_messages boolean,
    supports_inline boolean
);