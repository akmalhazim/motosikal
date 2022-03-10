CREATE TABLE `surveys` (
    `id` CHAR(36) PRIMARY KEY,
    `respondent_name` VARCHAR(255) NOT NULL,
    `respondent_email` VARCHAR(255),
    `respondent_phone` VARCHAR(255),
    `completed_at` TIMESTAMP,
    `result_percentage` TINYINT,
    `device_id` CHAR(36) NOT NULL,

    FOREIGN KEY (`device_id`) REFERENCES `devices`(`id`)
);
