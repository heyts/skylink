BEGIN;
CREATE TABLE likes (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar,
    actor_id varchar
);
CREATE INDEX likes_pk_idx ON likes (id, actor_id);
COMMIT;