use wakannaippi;
create table users(
    user_id varchar(60) not null,
    user_name varchar(60) not null,
    user_password varchar(100) not null,
    is_deleted boolean default 0,
    primary key(user_id) 
);

create table user_constellations(
    user_constellation_id varchar(60) not null,
    user_constellation_name varchar(60),
    user_id varchar(60) not null,
    user_constellation_data int,
    primary key(user_constellation_id),
    foreign key(user_id)
    references users(user_id)
);

create table user_items(
    user_item_id varchar(60) not null,
    user_id varchar(60) not null,
    user_item_quantity int not null default 0,
    primary key(user_item_id, user_id),
    foreign key(user_id)
    references users(user_id)
);


