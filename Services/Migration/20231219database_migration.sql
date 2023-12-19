CREATE DATABASE IF NOT EXISTS akasia;

USE akasia;

CREATE TABLE IF NOT EXISTS t_product(
    id varchar(36) UNIQUE PRIMARY KEY,
    title VARCHAR(30),
    description TEXT,
    rating FLOAT,
    image VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME
)