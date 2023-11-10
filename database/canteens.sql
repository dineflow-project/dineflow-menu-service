CREATE TABLE
    `canteens` (
        id bigint auto_increment PRIMARY KEY,
        name varchar(255) NOT NULL,
        image_path VARCHAR(255) NOT NULL
    );

INSERT INTO
    `canteens` (name, image_path)
VALUES (
        'iCanteen',
        '/images/canteen_item01.jpg'
    ), (
        'Commarts',
        '/images/canteen_item02.jpg'
    );