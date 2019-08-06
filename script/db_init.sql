-- DROP USER IF EXISTS wume_composer;
CREATE USER wume_composer WITH
    LOGIN
    NOSUPERUSER
    INHERIT
    NOCREATEDB
    NOCREATEROLE
    NOREPLICATION
    CONNECTION LIMIT -1
    PASSWORD 'wume_composer';

-- DROP DATABASE IF EXISTS wume_composer;
CREATE DATABASE wume_composer
    WITH
    OWNER = wume_composer
    TEMPLATE = template0
    ENCODING = 'UTF8'
    -- LC_COLLATE = 'ru_RU.UTF-8'
    -- LC_CTYPE = 'ru_RU.UTF-8'
    CONNECTION LIMIT = -1;
