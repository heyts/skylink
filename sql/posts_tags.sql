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
