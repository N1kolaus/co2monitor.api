CREATE TABLE IF NOT EXISTS locations (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE locations TO web;
GRANT USAGE, SELECT ON SEQUENCE locations_id_seq TO web;