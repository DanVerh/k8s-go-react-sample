CREATE DATABASE db;

USE db;

CREATE TABLE messages
( 
    id INT AUTO_INCREMENT PRIMARY KEY,
    lang VARCHAR(50) NOT NULL,
    hello VARCHAR(50) NOT NULL
);

INSERT INTO messages (lang, hello) 
VALUES ('Go', 'Hello, Go');

INSERT INTO messages (lang, hello) 
VALUES ('JS', 'Hello, JavaScript');
