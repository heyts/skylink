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