
DROP TABLE IF EXISTS collection_members;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS uploads;
DROP TABLE IF EXISTS collections;
DROP TABLE IF EXISTS collection_tags;
DROP TABLE IF EXISTS accounts;

DROP INDEX IF EXISTS accounts_email_uindex;
DROP EXTENSION IF EXISTS "uuid-ossp";