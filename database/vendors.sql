CREATE TABLE
    vendors (
        id bigint AUTO_INCREMENT PRIMARY KEY,
        canteen_id bigint NOT NULL,
        name VARCHAR(255) NOT NULL,
        owner_id VARCHAR(255) NOT NULL,
        opening_timestamp VARCHAR(255) NOT NULL,
        closing_timestamp VARCHAR(255) NOT NULL,
        status ENUM('Open', 'Close') NOT NULL,
        image_path VARCHAR(255) NOT NULL
    );

INSERT INTO
    vendors (
        canteen_id,
        name,
        owner_id,
        opening_timestamp,
        closing_timestamp,
        status,
        image_path
    )
VALUES (
        1,
        'VendorOne',
        'asdgfasdg31a5efg6e5',
        '08:00',
        '18:00',
        'Open',
        '/images/vendor_item01.jpg'
    ), (
        2,
        'VendorTwo',
        'asdgfasdg31a5efg6e6',
        '09:30',
        '19:30',
        'Close',
        '/images/vendor_item02.jpg'
    );