BEGIN;
CREATE TABLE posts_links (
    post_id varchar,
    link_id varchar
);
CREATE UNIQUE INDEX posts_links_posts_idx ON posts_links (post_id);
CREATE UNIQUE INDEX posts_links_links_idx ON posts_links (link_id);
COMMIT;
