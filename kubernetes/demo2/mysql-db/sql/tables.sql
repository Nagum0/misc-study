CREATE TABLE Albums (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    artist VARCHAR(255) UNIQUE,
    name VARCHAR(255) UNIQUE,
    release_date DATE
);

CREATE TABLE Reviews (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    album_id INT NOT NULL,
    user VARCHAR(255),
    review VARCHAR(1024),
    FOREIGN KEY (album_id) REFERENCES Albums(id)
);