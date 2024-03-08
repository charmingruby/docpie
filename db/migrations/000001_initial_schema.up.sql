CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY NOT NULL,
    
    name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    avatar_url varchar,
    role varchar NOT NULL,
    password varchar NOT NULL,

    collections_created_quantity int NOT NULL,
    collections_member_quantity integer NOT NULL,
    upload_quantity integer NOT NULL,

    deleted_by UUID REFERENCES accounts (id),

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS collection_tags (
    id UUID PRIMARY KEY NOT NULL,

    name varchar NOT NULL,
    description varchar NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp 
);

CREATE TABLE IF NOT EXISTS collections (
    id UUID PRIMARY KEY NOT NULL,
    
    name varchar NOT NULL,
    description varchar,
    secret varchar NOT NULL,
    tag varchar NOT NULL,
    
    uploads_quantity integer NOT NULL,
    members_quantity integer NOT NULL,

    tag_id UUID REFERENCES collection_tags (id) NOT NULL,
    creator_id UUID REFERENCES accounts (id) NOT NULL,
    deleted_by UUID REFERENCES accounts (id),

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS collection_members (
    id UUID PRIMARY KEY NOT NULL,
    
    role varchar NOT NULL,
    upload_quantity integer NOT NULL,

    account_id UUID REFERENCES accounts (id) NOT NULL,
    collection_id UUID REFERENCES collections (id) NOT NULL,

    joined_at timestamp DEFAULT now() NOT NULL,
    left_at timestamp,
    updated_at timestamp    
);

CREATE TABLE IF NOT EXISTS uploads (
    id UUID PRIMARY KEY NOT NULL,
    
    name varchar NOT NULL,
    url varchar NOT NULL,

    file_size varchar NOT NULL,
    file_mimetype varchar NOT NULL,

    collection_id UUID REFERENCES collections (id) NOT NULL,
    uploader_id UUID REFERENCES accounts (id) NOT NULL,
    
    uploaded_at timestamp DEFAULT now() NOT NULL,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY NOT NULL,

    context varchar NOT NULL,
    message_id varchar NOT NULL,

    recipient_id UUID REFERENCES accounts (id) NOT NULL,

    sent_at timestamp DEFAULT now() NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS accounts_email_uindex
    ON accounts (email);