CREATE TABLE IF NOT EXISTS co2_data (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    location_id bigserial NOT NULL,
    co2 integer NOT NULL,
    temp numeric(3, 1) NOT NULL,
    humidity integer NOT NULL,
    FOREIGN KEY (location_id) REFERENCES locations (id)
);

GRANT SELECT, INSERT ON TABLE co2_data TO web;
GRANT USAGE, SELECT ON SEQUENCE co2_data_id_seq TO web;