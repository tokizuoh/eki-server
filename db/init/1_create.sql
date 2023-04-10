CREATE DATABASE station_db;
USE station_db;

CREATE TABLE station
(
    `id` INT(7) NOT NULL PRIMARY KEY,
    `name` VARCHAR(30) NOT NULL,
    `is_liked` BOOLEAN NOT NULL DEFAULT 0
);
