CREATE TABLE images (
    image_id UUID PRIMARY KEY,
    name text,
    format text,
    size int,
    content blob
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name text,
    email text,
    password text
);
