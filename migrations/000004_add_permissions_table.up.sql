CREATE TABLE IF NOT EXISTS permissions (
    id bigserial PRIMARY KEY,
    code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions_locations (
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE,
    location_id bigint NOT NULL REFERENCES locations ON DELETE CASCADE,
    PRIMARY KEY (user_id, permission_id, location_id)
);

INSERT INTO permissions (code)
VALUES 
    ('co2:read'),
    ('co2:write'),
    ('location:read'),
    ('location:write'),
    ('location:admin');   

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE permissions TO web;
GRANT USAGE, SELECT ON SEQUENCE permissions_id_seq TO web;
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE users_permissions_locations TO web;