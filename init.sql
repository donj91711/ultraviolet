CREATE DATABASE IF NOT EXISTS UltraBooks;

USE UltraBooks;

CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    author VARCHAR(255),
    publication_date DATE,
    ISBN VARCHAR(13),
    page_count INT
);

/*
insert into books (title, author, publication_date, ISBN, page_count)
select "The Cat in the Hat", "Dr. Seuss", '1963-01-02', "123-456", 23;
*/
