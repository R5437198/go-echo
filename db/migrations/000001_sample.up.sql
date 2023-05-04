CREATE TABLE "user" (
    "id" uuid default gen_random_uuid(),
    first_name varchar(10) not null,
    last_name varchar(10) not null,
    birthday date not null
)