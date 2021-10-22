-- Write your migrate up statements here

CREATE TABLE users (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(200) NOT NULL,
    city VARCHAR(100) NOT NULL,
    zip_code VARCHAR(5) NOT NULL,
    date_of_birth TIMESTAMP WITH TIME ZONE NOT NULL,
    status VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);


---- create above / drop below ----


DROP TABLE users CASCADE;
