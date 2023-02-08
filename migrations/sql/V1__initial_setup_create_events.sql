create table if not exists usersauth(
    id uuid primary key not null,
    name varchar(255) not null,
    age int not null,
    regular bool not null,
    password varchar(255) not null,
    refreshtoken varchar(255)
);

create table if not exists books(
    id uuid primary key not null,
    "name" varchar(255) not null,
    "year" int not null,
    "new" bool not null
);

create table if not exists images(
    id uuid primary key not null,
    route varchar(255) not null
);

create table if not exists books_images(
    id serial primary key not null,
    book_id uuid not null,
    image_id uuid not null,
    FOREIGN KEY (book_id) REFERENCES books(id)
                                       ON DELETE CASCADE,
    FOREIGN KEY (image_id) REFERENCES images(id)
                                       ON DELETE CASCADE
)