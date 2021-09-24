-- Book Manage System Tables
-- book table
CREATE TABLE book(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT,
    title varchar(20) NOT NULL,
    price double(16,2) NOT NULL
)engine=InnoDB default charset=utf8mb4;
-- user table 
CREATE TABLE user(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT,
    username varchar(20) NOT NULL,
    password varchar(20) NOT NULL
)engine=InnoDB default charset=utf8mb4;