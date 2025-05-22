-- Cria o banco de dados (execute fora dele)
CREATE DATABASE booksdb;

-- Conecta ao banco de dados (no psql: \c booksdb)
\c booksdb;

-- Cria a tabela "books"
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);
