CREATE TABLE
    IF NOT EXISTS `users`
(
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` TEXT NOT NULL ,
    `password` TEXT NOT NULL
) ENGINE = InnoDB DEFAULT CHARSET = utf8
;
CREATE TABLE
    IF NOT EXISTS `sessions`
(
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `userId` INT NOT NULL,
    `expirationTime` DATETIME NOT NULL,
    `uuid` TEXT NOT NULL

) ENGINE = InnoDB DEFAULT CHARSET = utf8
;