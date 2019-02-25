# Docker-Golang-Postgres-app

init the DB  
`PATH in my case it's C:/Users/USER/Docker`  
mount init.sql script for DB init  
`docker run --name file_data_storage -v PATH/init.sql:/docker-entrypoint-initdb.d/init.sql -e POSTGRES_PASSWORD=password1 -d -p 5432:5432 postgres`

example  
`docker run --name file_data_storage -v C:/Users/USER/Docker/init.sql:/docker-entrypoint-initdb.d/init.sql -e POSTGRES_PASSWORD=password1 -d -p 5432:5432 postgres`

run app 
go run main.go -path="path to be traced"

connect to the DB via Docker  
`docker exec -it file_data_storage psql -U postgres`

the init.sql script is  
create database files_data;
\connect files_data;
CREATE TABLE files_data(
   id SERIAL PRIMARY KEY,
   name VARCHAR,
   fsize INT,
   fdate VARCHAR
);

INSERT INTO files_data(name,fsize)VALUES('test',20)