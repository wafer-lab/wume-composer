-- DROP USER IF EXISTS wume-composer;
CREATE USER wume-composer WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    CONNECTION LIMIT -1
    PASSWORD 'wume-composer';

-- DROP DATABASE IF EXISTS wume-composer;
CREATE DATABASE wume-composer
WITH OWNER = wume-composer
    TEMPLATE = template0
    ENCODING = 'UTF8'
    -- LC_COLLATE = 'ru_RU.UTF-8'
    -- LC_CTYPE = 'ru_RU.UTF-8'
    CONNECTION LIMIT = -1;
