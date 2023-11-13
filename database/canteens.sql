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
        'https://media.discordapp.net/attachments/1171496314807267378/1173531957066485771/2iCatneen-P2.jpg?ex=65644bcd&is=6551d6cd&hm=031a7afd179862b56a37ba2b231a54f86c732d8370967cdeabf7bfc32c836c8d&=&width=600&height=600'
    ), (
        'Commarts',
        'https://media.discordapp.net/attachments/1171496314807267378/1173531956596703303/Aksorn.jpg?ex=65644bcd&is=6551d6cd&hm=b8a52e6b3bd7f3424d399c23981d64533f948e0811811139f59a5f7425ce9e42&=&width=662&height=662'
    ),(
        'PolSci',
        'https://media.discordapp.net/attachments/1171496314807267378/1173531955879481374/PolSci.jpg?ex=65644bcd&is=6551d6cd&hm=d60b2cb7f380086e9f0c598397b0ec37c33c57f1e757f39ef5ec07bc93242b14&=&width=662&height=662'
    ), (
        'Medicine',
        'https://media.discordapp.net/attachments/1171496314807267378/1173531954939973712/DoctorCanteen.jpg?ex=65644bcd&is=6551d6cd&hm=25a1749b95a1785047e746f7853d8dbf25bbe0fa26c926c2df979bcc78955f36&=&width=662&height=662'
    ), (
        'FoodHubb',
        'https://media.discordapp.net/attachments/1171496314807267378/1173531954382114836/FoodHubb.jpg?ex=65644bcd&is=6551d6cd&hm=32ff1868e5062f125e32e72d4b0cf015740436b6604fe4037e7ec661ab7cbb5a&=&width=662&height=662'
    ),(
        'FoodWorld',
        'https://media.discordapp.net/attachments/1171496314807267378/1173531955363577976/FoodWorld.jpg?ex=65644bcd&is=6551d6cd&hm=3ad7cd3aadfd038cdd2e4a9ecbb21e622167bfd01e1a5c7ce0fc7a741dda30b5&=&width=662&height=662'
    );
