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
