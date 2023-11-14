create database lab;
use lab;

CREATE TABLE user (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) UNIQUE,
            password VARCHAR(255)
);

insert into user (username, password) values ('root11', 'root11');

create table if not exists document  (
                id INT PRIMARY KEY AUTO_INCREMENT,
                `key` VARCHAR(255),
                `value` TEXT,
                `created_at` BIGINT DEFAULT 0
);

DROP TABLE if exists document;