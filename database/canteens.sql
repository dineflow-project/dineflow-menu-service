CREATE TABLE `canteens`
(
    id   bigint auto_increment,
    name varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `canteens` (`name`) VALUES ('iCanteen'), ('Commarts');
