# Golang web app project

Description
A website that shows your book recommendations

usage:
Launch Docker postgres container and create database

$ docker pull postgres
$ docker run –-name reading-list-db-container -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 postgres
$ apt install postgresql-client
$ psql -h localhost -p 5432 -U postgres

CREATE DATABASE readinglist;

CREATE ROLE readinglist WITH LOGIN PASSWORD 'pa$$w0rd';

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,  
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    published integer NOT NULL,
    pages integer NOT NULL,
    genres text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);

GRANT SELECT, INSERT, UPDATE, DELETE ON books TO readinglist;
GRANT USAGE, SELECT ON SEQUENCE books_id_seq TO readinglist;

$ export READINGLIST_DB_DSN='postgres://readinglist:pa55w0rd@localhost/readinglist?sslmode=disable'
$ go run ./cmd/api
$ go run ./cmd/web

Open web browser and navigate to localhost:8081