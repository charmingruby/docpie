CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    
    name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    avatar_url varchar NOT NULL,
    password varchar NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL   
);

CREATE TABLE IF NOT EXISTS subscribers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,

    name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL   
);

CREATE TABLE IF NOT EXISTS newsletters_tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    title varchar NOT NULL,
    description varchar NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS newsletters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    
    title varchar NOT NULL,
    subject varchar NOT NULL,
    content text NOT NULL,
    status varchar NOT NULL,
    
    tag_id UUID REFERENCES newsletters_tags (id) NOT NULL,
    author_id UUID REFERENCES accounts (id) NOT NULL,

    published_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS mailing_lists (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    status varchar NOT NULL,
    
    recipient_id UUID REFERENCES subscribers (id) NOT NULL,
    newsletters_id UUID REFERENCES newsletters (id) NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS subscribers_email_uindex
    ON subscribers (email);

CREATE UNIQUE INDEX IF NOT EXISTS accounts_email_uindex
    ON accounts (email);