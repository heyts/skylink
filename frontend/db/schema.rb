# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[8.0].define(version: 2025_03_12_174016) do
  # These are extensions that must be enabled in order to support this database
  enable_extension "pg_catalog.plpgsql"

  create_table "actors", id: :string, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "display_name"
    t.string "handle"
    t.string "avatar"
    t.string "banner"
    t.integer "followers_count"
    t.integer "follows_count"
    t.integer "posts_count"
    t.index ["handle"], name: "handles_idx", unique: true
    t.index ["id"], name: "actors_pk_idx", unique: true
  end

  create_table "daily_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count"
    t.integer "reposts_count"
    t.integer "quotes_count"
    t.virtual "score", type: :integer, as: "((reposts_count * 10) + (COALESCE((quotes_count * 10), 0) * likes_count))", stored: true
    t.index ["post_id"], name: "posts_daily_stats_post_idx"
    t.index ["ymdh", "post_id"], name: "daily_stats_pk_idx", unique: true
    t.index ["ymdh"], name: "daily_posts_stats_ts_idx"
  end

  create_table "hourly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count"
    t.integer "reposts_count"
    t.integer "quotes_count"
    t.virtual "score", type: :integer, as: "((reposts_count * 10) + (COALESCE((quotes_count * 10), 0) * likes_count))", stored: true
    t.index ["post_id", "ymdh"], name: "posts_stats_pk_reverse_idx", unique: true
    t.index ["post_id"], name: "posts_hourly_stats_post_idx"
    t.index ["ymdh", "post_id"], name: "posts_stats_pk_idx", unique: true
    t.index ["ymdh"], name: "posts_stats_ts_idx"
  end

  create_table "links", id: :string, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "original_url"
    t.string "url"
    t.integer "count"
    t.string "title"
    t.string "og_title"
    t.string "og_description"
    t.string "og_site_name"
    t.string "og_image"
    t.jsonb "og_image_options"
    t.jsonb "og_optional"
    t.index ["count"], name: "link_count_idx"
    t.index ["url"], name: "urls_idx", unique: true
  end

  create_table "links_posts", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "link_id"
    t.datetime "created_at", precision: nil
    t.index ["link_id"], name: "links_link_id_idx"
    t.index ["post_id", "link_id"], name: "posts_links_idx", unique: true
    t.index ["post_id"], name: "links_post_id_posts_idx"
  end

  create_table "mentions_posts", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "actor_id"
    t.index ["post_id", "actor_id"], name: "posts_mentions_idx", unique: true
  end

  create_table "monthly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count"
    t.integer "reposts_count"
    t.integer "quotes_count"
    t.virtual "score", type: :integer, as: "((reposts_count * 10) + (COALESCE((quotes_count * 10), 0) * likes_count))", stored: true
    t.index ["ymdh", "post_id"], name: "monthly_stats_pk_idx", unique: true
  end

  create_table "posts", id: :string, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "collection"
    t.string "record_key"
    t.text "text"
    t.string "actor_id"
    t.string "language"
    t.string "tags", array: true
    t.datetime "published_at", precision: nil
    t.virtual "text_fts", type: :tsvector, as: "to_tsvector('simple'::regconfig, text)", stored: true
    t.string "country"
    t.string "locale"
    t.index ["actor_id"], name: "actors_idx"
    t.index ["collection", "record_key"], name: "collections_record_keys_idx", unique: true
    t.index ["country"], name: "posts_country_col_idx"
    t.index ["created_at", "id"], name: "posts_created_at_id_idx_tmp"
    t.index ["created_at"], name: "posts_created_at_idx"
    t.index ["language"], name: "posts_language_col_idx"
    t.index ["locale"], name: "posts_locale_col_idx"
    t.index ["published_at"], name: "published_at_idx"
    t.index ["tags"], name: "posts_tags_col_idx", using: :gin
    t.index ["text_fts"], name: "posts_fts_idx", using: :gin
  end

  create_table "posts_counts", id: false, force: :cascade do |t|
    t.datetime "period", precision: nil
    t.string "link_id"
    t.integer "posts_count"
    t.index ["period", "link_id"], name: "tmp_posts_counts_ymdh"
  end

  create_table "weekly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count"
    t.integer "reposts_count"
    t.integer "quotes_count"
    t.virtual "score", type: :integer, as: "((reposts_count * 10) + (COALESCE((quotes_count * 10), 0) * likes_count))", stored: true
    t.index ["ymdh", "post_id"], name: "weekly_stats_pk_idx", unique: true
  end

  create_view "top_links_1d", materialized: true, sql_definition: <<-SQL
      SELECT now() AS generated_at,
      l.id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      count(DISTINCT p.*) AS count_from_agg,
      sum(st.likes_count) AS likes,
      sum(st.reposts_count) AS reposts,
      sum(st.quotes_count) AS quotes,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score_from_agg
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN hourly_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE ((p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'P1D'::interval)) AND ((p.language)::text = 'en'::text))
    GROUP BY (now()), l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at))
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST
   LIMIT 30;
  SQL
  create_view "top_links_1d_fr", materialized: true, sql_definition: <<-SQL
      SELECT now() AS generated_at,
      l.id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      count(DISTINCT p.*) AS count_from_agg,
      sum(st.likes_count) AS likes,
      sum(st.reposts_count) AS reposts,
      sum(st.quotes_count) AS quotes,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score_from_agg
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN hourly_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE ((p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'P1D'::interval)) AND ((p.language)::text = 'fr'::text))
    GROUP BY (now()), l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at))
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST
   LIMIT 30;
  SQL
  create_view "top_last_day", materialized: true, sql_definition: <<-SQL
      SELECT l.id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      p.language,
      count(DISTINCT p.*) AS count_from_agg,
      sum(st.likes_count) AS likes,
      sum(st.reposts_count) AS reposts,
      sum(st.quotes_count) AS quotes,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN daily_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE (p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'P1D'::interval))
    GROUP BY l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at)), p.language
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST;
  SQL
  create_view "top_past_days", materialized: true, sql_definition: <<-SQL
      SELECT now() AS generated_at,
      l.id AS link_id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      p.language,
      count(DISTINCT p.*) AS posts_count,
      sum(st.likes_count) AS likes_count,
      sum(st.reposts_count) AS reposts_count,
      sum(st.quotes_count) AS quotes_count,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN daily_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE (p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'P1D'::interval))
    GROUP BY (now()), l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at)), p.language
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST;
  SQL
  create_view "top_past_hours", materialized: true, sql_definition: <<-SQL
      SELECT now() AS generated_at,
      l.id AS link_id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      p.language,
      count(DISTINCT p.*) AS posts_count,
      sum(st.likes_count) AS likes_count,
      sum(st.reposts_count) AS reposts_count,
      sum(st.quotes_count) AS quotes_count,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN hourly_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE (p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'PT1H'::interval))
    GROUP BY (now()), l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at)), p.language
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST;
  SQL
  create_view "top_past_weeks", materialized: true, sql_definition: <<-SQL
      SELECT now() AS generated_at,
      l.id AS link_id,
      age((now() AT TIME ZONE 'utc'::text), l.created_at) AS age,
      p.language,
      count(DISTINCT p.*) AS posts_count,
      sum(st.likes_count) AS likes_count,
      sum(st.reposts_count) AS reposts_count,
      sum(st.quotes_count) AS quotes_count,
      ((count(DISTINCT p.*) * 10) + sum(st.score)) AS score
     FROM (((links l
       JOIN links_posts lp ON (((l.id)::text = (lp.link_id)::text)))
       JOIN posts p ON (((p.id)::text = (lp.post_id)::text)))
       LEFT JOIN weekly_stats st ON (((st.post_id)::text = (p.id)::text)))
    WHERE (p.published_at > (date_trunc('hour'::text, (now() AT TIME ZONE 'utc'::text)) - 'P7D'::interval))
    GROUP BY (now()), l.id, (age((now() AT TIME ZONE 'utc'::text), l.created_at)), p.language
    ORDER BY ((count(DISTINCT p.*) * 10) + sum(st.score)) DESC NULLS LAST;
  SQL
end
