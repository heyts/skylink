BEGIN;
CREATE TABLE hyperlinks (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar PRIMARY KEY,
    url varchar
);
CREATE UNIQUE INDEX hyperlinks_pk_idx ON hyperlinks (id);
CREATE UNIQUE INDEX urls_idx ON hyperlinks (url);
COMMIT;

BEGIN;
CREATE TABLE actors (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar PRIMARY KEY,
    display_name varchar,
    handle varchar,

    profile_url varchar
);
CREATE UNIQUE INDEX actors_pk_idx ON actors (id);
CREATE UNIQUE INDEX handles_idx ON actors (handle);
COMMIT;

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

BEGIN;
CREATE TABLE posts_hyperlinks (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    post_id varchar,
    hyperlink_id varchar
);
CREATE UNIQUE INDEX posts_yperlinks_idx ON posts_hyperlinks (post_id);
CREATE UNIQUE INDEX hyperlink_idx ON posts_hyperlinks (hyperlink_id);
COMMIT;

BEGIN;
CREATE TABLE posts_tags (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    post_id varchar,
    tag_id varchar
);
CREATE UNIQUE INDEX posts_tags_idx ON posts_tags (post_id);
CREATE UNIQUE INDEX tag_idx ON posts_tags (tag_id);
COMMIT;

BEGIN;
CREATE TABLE posts_mentions (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    post_id varchar,
    actor_id varchar
);
CREATE UNIQUE INDEX posts_mentions_idx ON posts_mentions (post_id);
CREATE UNIQUE INDEX actors_mentions_idx ON posts_mentions (actor_id);
COMMIT;

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




