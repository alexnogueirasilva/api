CREATE
    DATABASE IF NOT EXISTS `devbook` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE
    `devbook`;

DROP TABLE IF EXISTS `publications`;
DROP TABLE IF EXISTS `followers`;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`       int(11)                                        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`     varchar(50) COLLATE utf8mb4_unicode_ci         NOT NULL,
    `nickname` varchar(50) UNIQUE COLLATE utf8mb4_unicode_ci  NOT NULL,
    `email`    varchar(100) UNIQUE COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` varchar(120) COLLATE utf8mb4_unicode_ci        NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP

) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `followers`;

CREATE TABLE `followers`
(
    `user_id`     int(11) NOT NULL,
    `follower_id` int(11) NOT NULL,
    created_at    timestamp DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
    PRIMARY KEY (`user_id`, `follower_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `publications`;

CREATE TABLE `publications`
(
    `id`         int(11)                                NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title`      varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
    `content`    text COLLATE utf8mb4_unicode_ci        NOT NULL,
    `author_id`  int(11)                                NOT NULL,
    `likes`      int(11)                                NOT NULL DEFAULT 0,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
