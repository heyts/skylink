BEGIN;
CREATE TABLE posts (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar PRIMARY KEY,
    collection varchar,
    record_key varchar,
    text TEXT,
    actor_id varchar
);
CREATE UNIQUE INDEX posts_pk_idx ON posts (id);
CREATE UNIQUE INDEX actors_idx ON posts (actor_id);
CREATE UNIQUE INDEX record_keys_idx ON posts (record_key);
COMMIT;
