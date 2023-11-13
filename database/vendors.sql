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
        'https://media.discordapp.net/attachments/1171496314807267378/1173528248215416952/chu5602.png?ex=65644859&is=6551d359&hm=dd5f2b2e1c0c519a0a674c97102f69d9ccc490483a1d2c0a7d6742a823c20153'
    ), (
        2,
        'VendorTwo',
        'asdgfasdg3185fg6e6',
        '09:30',
        '19:30',
        'Close',
        'https://media.discordapp.net/attachments/1171496314807267378/1173528664038707240/chu5604.png?ex=656448bc&is=6551d3bc&hm=9ad29e61367299bc0a38bc07b2378e4e9f0fd1648f157d6e453fcfb107250aca'
    );