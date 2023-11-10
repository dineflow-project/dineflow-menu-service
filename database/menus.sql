CREATE TABLE
    menus (
        id bigint AUTO_INCREMENT PRIMARY KEY,
        vendor_id bigint NOT NULL,
        name VARCHAR(255) NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        image_path VARCHAR(255) NOT NULL,
        description VARCHAR(255),
        is_available INT NOT NULL
    );

-- INSERT INTO

--     menus (

--         vendor_id,

--         name,

--         price,

--         image_path,

--         description,

--         is_available

--     )

-- VALUES (

--         1,

--         'Example Menu Item01',

--         9.99,

--         '/images/menu_item01.jpg',

--         'A delicious example menu item01.',

--         0

--     ), (

--         2,

--         'Example Menu Item02',

--         9999.99,

--         '/images/menu_item02.jpg',

--         'A delicious example menu item02.',

--         1

--     );