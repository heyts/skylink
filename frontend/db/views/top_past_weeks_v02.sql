SELECT
    now() generated_at,
    l.id link_id,
    age(now() at time zone 'utc', l.created_at) age,
    p.language language,
    count(distinct p.*) posts_count,
    sum(st.likes_count) likes_count,
    sum(st.reposts_count) reposts_count,
    sum(st.quotes_count) quotes_count,
    (count(distinct p.*) * 10 + sum(st.score)) score
FROM
    links l
    join links_posts lp on l.id = lp.link_id
    join posts p on p.id = lp.post_id
    left outer join weekly_stats st on st.post_id = p.id
WHERE
    p.published_at > date_trunc ('hour', now () at time zone 'utc') - interval '1 week'
GROUP BY
    1,
    2,
    3,
    4
ORDER BY
    9 DESC NULLS LAST;