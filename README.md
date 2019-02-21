# Docker-Golang-Postgres-app


init the DB
PATH in my case it's C:/Users/USER/Docker
mount init.sql script for DB init
docker run --name file_data_storage -v PATH/init.sql:/docker-entrypoint-initdb.d/init.sql -e POSTGRES_PASSWORD=password1 -d postgres

example
docker run --name file_data_storage -v C:/Users/USER/Docker/init.sql:/docker-entrypoint-initdb.d/init.sql -e POSTGRES_PASSWORD=password1 -d postgres

connect to the DB
docker exec -it file_data_storage psql -U postgres

the init.sql script is
create database sometest;
\connect sometest;
CREATE TABLE files_data(
   id SERIAL PRIMARY KEY,
   name VARCHAR NOT NULL,
   ext VARCHAR ,
   fsize INT,
   fdate TIMESTAMPTZ
);

INSERT INTO files_data(name,ext,fsize)VALUES('test','txt',20)
