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
