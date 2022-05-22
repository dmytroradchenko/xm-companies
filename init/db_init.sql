CREATE DATABASE "xm-companies"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1251'
    LC_CTYPE = 'English_United States.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

GRANT ALL ON DATABASE "xm-companies" TO postgres;

CREATE USER companies_admin WITH PASSWORD '******';

GRANT ALL ON DATABASE "xm-companies" TO companies_admin;

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username text UNIQUE not null,
    password text not null
);

CREATE TABLE companies (
    code text PRIMARY KEY not null,
    name text not null,
    country text,
    phone text,
    website text
)
