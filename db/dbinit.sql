CREATE DATABASE db;

\c db

CREATE TABLE messages
( 
    id serial primary key,
    lang varchar(50) NOT NULL,
    hello varchar(50) NOT NULL
);

INSERT INTO messages (lang, hello) 
VALUES ('Go', 'Hello, Go!');

INSERT INTO messages (lang, hello) 
VALUES ('JS', 'Hello, JavaScript!')