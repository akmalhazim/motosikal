CREATE TABLE `records` (
    `id` CHAR(36) PRIMARY KEY,
    `lat` POINT NOT NULL,
    `lng` POINT NOT NULL,
    `timestamp` TIMESTAMP NOT NULL,
    `device_id` CHAR(36) NOT NULL,
    `survey_id` CHAR(36),

    FOREIGN KEY (`device_id`) REFERENCES `devices`(`id`),
    FOREIGN KEY (`survey_id`) REFERENCES `surveys`(`id`)
);
