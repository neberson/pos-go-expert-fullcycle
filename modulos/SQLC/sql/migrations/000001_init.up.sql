CREATE TABLE Categories (
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    description text
);

CREATE TABLE Courses (
    id varchar(36) NOT NULL PRIMARY KEY,
    category_id varchar(36),
    name text NOT NULL,
    description text,
    price decimal(10, 2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories(id) ON DELETE SET NULL
);