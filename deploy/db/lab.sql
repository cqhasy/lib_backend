create database if not exists lab;
use lab;

CREATE TABLE if not exists user (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(50) UNIQUE,
            password VARCHAR(255),
            description TEXT
);

create table if not exists document  (
                id INT PRIMARY KEY AUTO_INCREMENT,
                block VARCHAR(255) NOT NULL,
                group_name VARCHAR(255) NOT NULL,
                title VARCHAR(255) NOT NULL,
                create_at BIGINT NOT NULL,
                content TEXT NOT NULL,
                avatar VARCHAR(255)
);