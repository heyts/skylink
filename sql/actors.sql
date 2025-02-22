BEGIN;
CREATE TABLE actors (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    -- did identifier
    id varchar PRIMARY KEY,
    display_name varchar,
    handle varchar,

    profile_url varchar
);
CREATE UNIQUE INDEX actors_pk_idx ON actors (id);
CREATE UNIQUE INDEX handles_idx ON actors (handle);
COMMIT;