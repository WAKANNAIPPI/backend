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

create table conste_stars(
    id int AUTO_INCREMENT,
    user_constellation_id varchar(60) not null,
    conste_stored_star json,
    foreign key(user_constellation_id)
    references user_constellation_id(user_constellation_id),
    primary key(id)
);
create table conste_lines(
    id int AUTO_INCREMENT,
    user_constellation_id varchar(60) not null,
    conste_stored_line json,
    foreign key(user_constellation_id)
    references user_constellation_id(user_constellation_id),
    primary key(id)
);


