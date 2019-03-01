/*docker run --name file_data_storage -e POSTGRES_PASSWORD=password1 -d postgres
docker exec -it file_data_storage psql -U postgres -c "CREATE DATABASE file_db"
docker exec -it file_data_storage psql -U postgres
*/
create database files_data;
\connect files_data;
/* CREATE TABLE files_data(
   id SERIAL PRIMARY KEY,
   name VARCHAR,
   fsize INT,
   fdate VARCHAR
);

INSERT INTO files_data(name,fsize)VALUES('test',20) */