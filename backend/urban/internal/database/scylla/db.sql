CREATE TABLE images (
    id UUID PRIMARY KEY,
    name text,
    date timestamp,
    content blob
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name text,
    email text,
    password text
);
