-- Make DB
CREATE DATABASE testdb;

\c testdb

-- Make shema
CREATE SCHEMA testschema;

-- role
CREATE ROLE test WITH LOGIN PASSWORD 'test';

--
GRANT ALL PRIVILEGES ON SCHEMA testschema TO test;