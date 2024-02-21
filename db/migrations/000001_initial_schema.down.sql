DROP INDEX IF EXISTS subscribers_email_uindex;
DROP INDEX IF EXISTS accounts_email_uindex;

DROP TABLE IF EXISTS subscribers;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS mailing_lists;
DROP TABLE IF EXISTS newsletters;
DROP TABLE IF EXISTS newsletters_tags;

DROP EXTENSION IF EXISTS "uuid-ossp";