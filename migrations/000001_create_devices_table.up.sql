CREATE TABLE `devices` (
    `id` CHAR(36) PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `last_ping` TIMESTAMP
);
