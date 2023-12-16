-- init tables

CREATE TABLE news (
    id serial not null,
    title varchar(255) not null,
    content text not null,
    primary key(id)
);

CREATE TABLE news_categories (
    news_id int not null,
    category_id int not null,
    foreign key (news_id) references news (id) on delete cascade on update cascade
);

CREATE TABLE users (
    id serial not null,
    login varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);

-- Fill database with example data

INSERT INTO news (title, content) values ('Test', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.');
INSERT INTO news (title, content) values ('Test2', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.');
INSERT INTO news (title, content) values ('Test3', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.');


INSERT INTO news_categories (news_id, category_id) values (1, 1);
INSERT INTO news_categories (news_id, category_id) values (1, 2);
INSERT INTO news_categories (news_id, category_id) values (2, 5);
INSERT INTO news_categories (news_id, category_id) values (3, 4);


