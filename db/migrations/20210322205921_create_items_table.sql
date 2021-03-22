
-- +goose Up
CREATE TABLE `items` (
	`id`       	 INT(11) NOT NULL AUTO_INCREMENT,
	`name`    	 VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`user_id` 	 INT(11) NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)    
);

-- +goose Down
DROP TABLE `items`;