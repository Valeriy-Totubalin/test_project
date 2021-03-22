
-- +goose Up
CREATE TABLE `users` (
	`id`       	 INT(11) NOT NULL AUTO_INCREMENT,
	`login`    	 VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`password` 	 VARCHAR(255) NOT NULL COLLATE 'utf8mb4_general_ci',

    PRIMARY KEY (`id`),
    UNIQUE INDEX `login` (`login`)
);

-- +goose Down
DROP TABLE `users`;