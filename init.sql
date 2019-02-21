/*docker run --name file_data_storage -e POSTGRES_PASSWORD=password1 -d postgres
docker exec -it file_data_storage psql -U postgres -c "CREATE DATABASE file_db"
docker exec -it file_data_storage psql -U postgres
*/
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
