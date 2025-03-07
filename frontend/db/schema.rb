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

ActiveRecord::Schema[8.0].define(version: 2025_03_06_232025) do
  # These are extensions that must be enabled in order to support this database
  enable_extension "pg_catalog.plpgsql"

  create_table "actors", id: :string, force: :cascade do |t|
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string "display_name"
    t.string "handle"
    t.string "avatar"
    t.string "banner"
    t.integer "followers_count"
    t.integer "follows_count"
    t.integer "posts_count"
    t.index ["handle"], name: "actor_handle_idx"
  end

  create_table "daily_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count", default: 0
    t.integer "reposts_count", default: 0
    t.integer "quotes_count", default: 0
    t.index ["post_id"], name: "index_daily_stats_on_post_id"
    t.index ["ymdh", "post_id"], name: "daily_stats_pk_col_idx", unique: true
  end

  create_table "hourly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count", default: 0
    t.integer "reposts_count", default: 0
    t.integer "quotes_count", default: 0
    t.index ["post_id"], name: "index_hourly_stats_on_post_id"
    t.index ["ymdh", "post_id"], name: "hourly_stats_pk_col_idx", unique: true
  end

  create_table "links", id: :string, force: :cascade do |t|
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string "original_url"
    t.string "url"
    t.integer "count"
    t.string "title"
    t.string "og_title"
    t.text "og_description"
    t.string "og_site_name"
    t.string "og_image"
    t.json "og_image_options"
    t.json "og_optional"
    t.index ["url"], name: "links_urls_col_idx"
  end

  create_table "links_posts", id: false, force: :cascade do |t|
    t.string "link_id"
    t.string "post_id"
    t.index ["link_id", "post_id"], name: "links_posts_idx", unique: true
    t.index ["link_id"], name: "index_links_posts_on_link_id"
    t.index ["post_id"], name: "index_links_posts_on_post_id"
    t.index ["post_id"], name: "links_post_id_posts_idx"
  end

  create_table "mentions_posts", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "actor_id"
    t.index ["actor_id"], name: "index_mentions_posts_on_actor_id"
    t.index ["post_id", "actor_id"], name: "mentions_posts_idx", unique: true
    t.index ["post_id"], name: "index_mentions_posts_on_post_id"
  end

  create_table "monthly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count", default: 0
    t.integer "reposts_count", default: 0
    t.integer "quotes_count", default: 0
    t.index ["post_id"], name: "index_monthly_stats_on_post_id"
    t.index ["ymdh", "post_id"], name: "monthly_stats_pk_col_idx", unique: true
  end

  create_table "posts", id: :string, force: :cascade do |t|
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.datetime "published_at", precision: nil
    t.string "collection"
    t.string "record_key"
    t.text "text"
    t.string "language"
    t.string "tags", array: true
    t.string "actor_id"
    t.index ["actor_id"], name: "index_posts_on_actor_id"
    t.index ["language"], name: "posts_language_col_idx"
    t.index ["tags"], name: "posts_tags_col_idx", using: :gin
  end

  create_table "weekly_stats", id: false, force: :cascade do |t|
    t.datetime "ymdh", precision: nil
    t.string "post_id"
    t.integer "likes_count", default: 0
    t.integer "reposts_count", default: 0
    t.integer "quotes_count", default: 0
    t.index ["post_id"], name: "index_weekly_stats_on_post_id"
    t.index ["ymdh", "post_id"], name: "weekly_stats_pk_col_idx", unique: true
  end
end
