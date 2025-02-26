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
CREATE UNIQUE INDEX handles_idx ON actors (handle);

CREATE TABLE languages (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    
    -- format:  <country>[-<language>]
    id VARCHAR,
    
    -- ISO 3166 Country Code (Unused for now)
    country VARCHAR NOT NULL,

    -- ISO 639 Language Code (Unused for now)
    language VARCHAR 
);
CREATE UNIQUE INDEX language_pk_idx ON languages (id);
CREATE INDEX language_language_idx ON languages (country);

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

CREATE TABLE posts_links (
    post_id varchar,
    link_id varchar
);
CREATE UNIQUE INDEX posts_links_idx ON posts_links (post_id, link_id);

CREATE TABLE posts_languages (
    post_id varchar,
    language_id varchar
);
CREATE UNIQUE INDEX posts_languages_idx ON posts_languages (post_id, language_id);


CREATE TABLE posts_mentions (
    post_id varchar,
    actor_id varchar
);
CREATE UNIQUE INDEX posts_mentions_idx ON posts_mentions (post_id, actor_id);

CREATE TABLE posts_tags (
    post_id varchar,
    tag_id varchar
);
CREATE UNIQUE INDEX posts_tags_idx ON posts_tags (post_id, tag_id);

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
CREATE INDEX actors_idx ON posts (actor_id);
CREATE UNIQUE INDEX collections_record_keys_idx ON posts (collection, record_key);

CREATE TABLE tags (
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    id varchar,
    label varchar
);
CREATE UNIQUE INDEX tags_idx ON tags (id);
CREATE UNIQUE INDEX labels_idx ON tags (label);
COMMIT;