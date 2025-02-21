BEGIN;
CREATE TABLE links (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar PRIMARY KEY,
    url varchar,
);
CREATE UNIQUE INDEX links_pk_idx ON links (id);
CREATE UNIQUE INDEX urls_idx ON links (url);
COMMIT;
