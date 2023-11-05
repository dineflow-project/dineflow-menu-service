CREATE TABLE vendors (
    id bigint AUTO_INCREMENT PRIMARY KEY,
    canteen_id bigint NOT NULL,
    name VARCHAR(255) NOT NULL,
    owner_id bigint,
    opening_timestamp TIMESTAMP NOT NULL,
    closing_timestamp TIMESTAMP NOT NULL,
    status ENUM('Open', 'Close') NOT NULL
);

INSERT INTO vendors (canteen_id, name, owner_id, opening_timestamp, closing_timestamp, status)
VALUES
    (1, 'Vendor 1', NULL, '2023-09-16 08:00:00', '2023-09-16 18:00:00', 'Open'),
    (2, 'Vendor 2', NULL, '2023-09-16 09:30:00', '2023-09-16 19:30:00', 'Close');
