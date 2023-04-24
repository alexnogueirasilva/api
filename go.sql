CREATE
DATABASE IF NOT EXISTS `devbook` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE
`devbook`;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`       int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`     varchar(50) COLLATE utf8mb4_unicode_ci         NOT NULL,
    `nickname` varchar(50) UNIQUE COLLATE utf8mb4_unicode_ci  NOT NULL,
    `email`    varchar(100) UNIQUE COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(32) COLLATE utf8mb4_unicode_ci         NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP

) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;