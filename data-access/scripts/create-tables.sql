
DROP TABLE IF EXISTS movie;
CREATE TABLE movie (
                       id         INT AUTO_INCREMENT NOT NULL,
                       title      VARCHAR(128) NOT NULL,
                       price      DECIMAL(5,2) NOT NULL,
                       PRIMARY KEY (`id`)
);

INSERT INTO movie (title, price)
VALUES
    ('John Wick II', 56.99),
    ('Black Panther', 63.99),
    ('Aquaman', 17.99),
    ('Iron man', 56.99),
    ('Batman Origin', 63.99),
    ('Spiderman in the Spiderverse', 17.99),
    ('Antman I', 56.99),
    ('Hitman Bodyguard', 63.99),
    ('The Recruit', 17.99),
    ('Equalizer I', 34.98),
    ('Equalizer II', 34.98);

CREATE TABLE client (
                        id         INT AUTO_INCREMENT NOT NULL,
                        first_name      VARCHAR(128) NOT NULL,
                        last_name      VARCHAR(128) NOT NULL,
                        email      VARCHAR(128) NOT NULL,
                        PRIMARY KEY (`id`)
);

INSERT INTO client (first_name, last_name, email)
VALUES
    ('Jon', 'Wick', 'jon@wick.com'),
    ('Chadwick', 'Boseman', 'chadwick@boseman.com'),
    ('Denzel', 'Washington', 'denzel@washington.com');

DROP TABLE IF EXISTS favorite_movie;
CREATE TABLE favorite_movie (
                                id         INT AUTO_INCREMENT NOT NULL,
                                client_id  INT NOT NULL,
                                movie_id   INT NOT NULL,
                                PRIMARY KEY (`id`),
                                FOREIGN KEY (client_id) REFERENCES client(id),
                                FOREIGN KEY (movie_id) REFERENCES movie(id)
);

