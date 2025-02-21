BEGIN;
CREATE TABLE tags (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar,
    label varchar
);
CREATE UNIQUE INDEX tags_idx ON tags (id);
CREATE UNIQUE INDEX labels_idx ON tags (label);
COMMIT;