CREATE TABLE categories (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    description text
);

CREATE TABLE courses (
    id int PRIMARY KEY AUTO_INCREMENT,
    category_id int NOT NULL,
    name varchar(255) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories(id));