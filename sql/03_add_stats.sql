BEGIN;
CREATE TABLE hourly_stats (
    ymdh TIMESTAMP,
    post_id varchar,
    likes_count integer,
    reposts_count integer
);
-- TODO probably needs a GIST index on date ?
CREATE UNIQUE INDEX hourly_stats_pk_idx on hourly_stats (ymdh, post_id);

CREATE TABLE daily_stats (
    ymdh TIMESTAMP,
    post_id varchar,
    likes_count integer,
    reposts_count integer
);
-- TODO probably needs a GIST index on date ?
CREATE UNIQUE INDEX daily_stats_pk_idx on daily_stats (ymdh, post_id);

CREATE TABLE weekly_stats (
    ymdh TIMESTAMP,
    post_id varchar,
    likes_count integer,
    reposts_count integer
);
-- TODO probably needs a GIST index on date ?
CREATE UNIQUE INDEX weekly_stats_pk_idx on weekly_stats (ymdh, post_id);

CREATE TABLE monthly_stats (
    ymdh TIMESTAMP,
    post_id varchar,
    likes_count integer,
    reposts_count integer
);
-- TODO probably needs a GIST index on date ?
CREATE UNIQUE INDEX monthly_stats_pk_idx on monthly_stats (ymdh, post_id);
COMMIT;