BEGIN;
CREATE TABLE actors (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    -- did identifier
    id varchar PRIMARY KEY,
    display_name varchar,
    handle varchar,
    avatar varchar,
    banner varchar,

    followers_count INTEGER,
    follows_count INTEGER,
    posts_count INTEGER
);
CREATE UNIQUE INDEX actors_pk_idx ON actors (id);
CREATE INDEX handles_idx ON actors (handle);

CREATE TABLE links (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar PRIMARY KEY,
    original_url varchar,
    url varchar,
    count integer
);
CREATE UNIQUE INDEX links_pk_idx ON links (id);
CREATE UNIQUE INDEX urls_idx ON links (url);

CREATE TABLE posts (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    published_at TIMESTAMP,

    id varchar PRIMARY KEY,
    collection varchar,
    record_key varchar,
    text TEXT,
    language varchar,
    tags varchar[],
    actor_id varchar
);
CREATE UNIQUE INDEX posts_pk_idx ON posts (id);
CREATE INDEX actors_idx ON posts (actor_id);
CREATE UNIQUE INDEX collections_record_keys_idx ON posts (collection, record_key);
CREATE INDEX posts_language_col_idx ON posts (language);
CREATE INDEX posts_tags_col_idx ON posts USING gin (tags);

CREATE TABLE links_posts (
    post_id varchar,
    link_id varchar
);
CREATE UNIQUE INDEX links_posts_idx ON links_posts(post_id, link_id);
CREATE INDEX links_post_id_posts_idx ON links_posts(post_id);
CREATE INDEX links_link_id_posts_idx ON links_posts(post_id);

CREATE TABLE mentions_posts (
    post_id varchar,
    actor_id varchar
);
CREATE UNIQUE INDEX mentions_posts_idx ON posts_mentions (post_id, actor_id);
COMMIT;
