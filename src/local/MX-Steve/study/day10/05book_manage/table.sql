-- Book Manage System Tables
-- book table
CREATE TABLE book(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT,
    title varchar(20) NOT NULL,
    price double(16,2) NOT NULL,
    publisher_id bigint(20) default 0
)engine=InnoDB default charset=utf8mb4;
-- publisher table
CREATE TABLE publisher(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT,
    province varchar(20) NOT NULL,
    city varchar(20) NOT NULL,
    name varchar(20) NOT NULL
)engine=InnoDB default charset=utf8mb4;
-- sql
-- select 
--     book.id,book.title,book.price,publisher.province,publisher.city,publisher.name
-- from book join publisher 
-- on book.publisher_id = publisher.id and book.id=10;

-- user table 
CREATE TABLE user(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT,
    username varchar(20) NOT NULL,
    password varchar(20) NOT NULL
)engine=InnoDB default charset=utf8mb4;
