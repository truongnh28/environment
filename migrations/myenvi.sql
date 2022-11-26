DROP TABLE IF EXISTS `albums`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `albums` (
                          `album_ID` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(40) NOT NULL,
                          `artist_ID` int NOT NULL,
                          `created_at` date DEFAULT NULL,
                          `upload_at` date DEFAULT NULL,
                          `cover_img` varchar(400) DEFAULT NULL,
                          PRIMARY KEY (`album_ID`),
                          KEY `artist_ID` (`artist_ID`),
                          CONSTRAINT `albums_ibfk_1` FOREIGN KEY (`artist_ID`) REFERENCES `artists` (`artist_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `reports`;

CREATE TABLE `reports` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(500) NOT NULL,
    `description` text NOT NULL,
    `created_at` date DEFAULT NULL,
    `update_at` date DEFAULT NULL,
    `status` varchar(20) NOT NULL,
    `author` varchar(100) NOT NULL,
    `lag` float(20,15) NOT NULL,
    `lng` float(20,15) NOT NULL,
    PRIMARY KEY (`id`),
    `resolver_id` int
)

DROP TABLE IF EXISTS `images`;

CREATE TABLE `images` (
    `id` int NOT NULL AUTO_INCREMENT,
    `report_id` int,
    `url` varchar(200) NOT NULL,
    PRIMARY KEY (`id`),
)

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_name` varchar(50) NOT NULL,
    `pass_word` varchar(50) NOT NULL,
    `is_resolver` bool,
    PRIMARY KEY (`id`),
    
)