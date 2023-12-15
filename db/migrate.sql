CREATE TABLE news (
    id serial not null,
    title varchar(255) not null,
    content text not null,
    primary key(id)
);

CREATE TABLE news_categories (
    news_id int not null,
    category_id serial not null,
    primary key(category_id)
);

CREATE TABLE users (
    id serial not null,
    login varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);
