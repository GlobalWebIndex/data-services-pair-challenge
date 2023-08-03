-- Create the database user
CREATE USER testuser WITH PASSWORD 'passwd';

-- Create the database
CREATE DATABASE testdb OWNER testuser;
