CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name varchar NOT NULL,
    last_name varchar NOT NULL,
    email varchar NOT NULL,
    role varchar NOT NULL,
    avatar_url varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL   
);

CREATE UNIQUE INDEX IF NOT EXISTS accounts_email_uindex
    ON accounts (email);


CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name varchar NOT NULL,
    description varchar NOT NULL,
    logo_url varchar NOT NULL,

    account_id UUID REFERENCES accounts (id) NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS product_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    total_feedbacks integer NOT NULL,
    amount_of_stars integer NOT NULL,

    product_id UUID REFERENCES products (id) NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS feedback_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    name varchar NOT NULL,
    description varchar NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS feedbacks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() NOT NULL,
    rate integer NOT NULL,
    comment text NOT NULL,

    feedback_category_id UUID REFERENCES feedback_categories (id) NOT NULL,
    product_id UUID REFERENCES products (id) NOT NULL,
    account_id UUID REFERENCES accounts (id) NOT NULL,

    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);
