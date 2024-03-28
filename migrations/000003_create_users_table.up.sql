CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text UNIQUE NOT NULL,
    token text NOT NULL,
    active bool NOT NULL
);

CREATE INDEX IF NOT EXISTS users_name_idx ON users USING GIN (to_tsvector('simple', name));

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users TO web;
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO web;