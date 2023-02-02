create table if not exists usersauth(
    id uuid primary key not null,
    name varchar(255) not null,
    age int not null,
    regular bool not null,
    password varchar(255) not null,
    refreshtoken varchar(255)
);