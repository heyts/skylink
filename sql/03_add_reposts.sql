BEGIN;
CREATE TABLE reposts (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar,
    actor_id varchar
);
CREATE INDEX reposts_pk_idx ON reposts (id, actor_id);
COMMIT;