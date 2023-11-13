CREATE TABLE
    menus (
        id bigint AUTO_INCREMENT PRIMARY KEY,
        vendor_id bigint NOT NULL,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        image_path VARCHAR(255) NOT NULL,
        description VARCHAR(255),
        is_available VARCHAR(255) NOT NULL
    );

INSERT INTO
    menus (
        vendor_id,
        name,
        price,
        image_path,
        description,
        is_available
    )
VALUES (
        1,
        'Salad',
        80,
        'https://media.discordapp.net/attachments/1171496314807267378/1171496383795167232/photo-1546069901-ba9599a7e63c.png?ex=655ce407&is=654a6f07&hm=0b8e4beddeb4e4b703c80f8a36a470f4d0b77c400ba93390e91f56a047ad3e71&=&width=662&height=662',
        'A delicious salad.',
        'yes'
    ), (
        1,
        'Pizza',
        200,
        'https://media.discordapp.net/attachments/1171496314807267378/1171496484542361640/photo-1565299624946-b28f40a0ae38.png?ex=655ce41f&is=654a6f1f&hm=e7f4a7caf3ca680889ad63611fdc49e2ab35aba3724218f94619f1d4eb913a19&=&width=548&height=548',
        'A cheese pizza.',
        'no'
    );