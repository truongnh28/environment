DROP TABLE IF EXISTS `reports`;

CREATE TABLE `reports` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(500) NOT NULL,
    `description` text NOT NULL,
    `created_at` date DEFAULT NULL,
    `updated_at` date DEFAULT NULL,
    `deleted_at` date DEFAULT NULL,
    `status` varchar(20) NOT NULL,
    `author` varchar(100) NOT NULL,
    `lag` float(20,15) NOT NULL,
    `lng` float(20,15) NOT NULL,
    PRIMARY KEY (`id`),
    `resolver_id` int
)

DROP TABLE IF EXISTS `images`;

CREATE TABLE `images` (
    `id` varchar(100) NOT NULL,
    `report_id` int,
    `url` text NOT NULL,
    `status` int,
    PRIMARY KEY (`id`),
)

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_name` varchar(50) NOT NULL,
    `pass_word` varchar(50) NOT NULL,
    `is_resolver` bool,
    `email` varchar(50) NOT NULL,
    `phone` varchar(10),
    PRIMARY KEY (`id`),   
)