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

ActiveRecord::Schema[8.0].define(version: 0) do
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

  create_table "languages", id: false, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "id"
    t.string "country", null: false
    t.string "language"
    t.index ["country"], name: "language_language_idx"
    t.index ["id"], name: "language_pk_idx", unique: true
  end

  create_table "links", id: :string, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "original_url"
    t.string "url"
    t.integer "count"
    t.index ["id"], name: "links_pk_idx", unique: true
    t.index ["url"], name: "urls_idx", unique: true
  end

  create_table "posts", id: :string, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "collection"
    t.string "record_key"
    t.text "text"
    t.string "actor_id"
    t.index ["actor_id"], name: "actors_idx"
    t.index ["collection", "record_key"], name: "collections_record_keys_idx", unique: true
    t.index ["id"], name: "posts_pk_idx", unique: true
  end

  create_table "posts_languages", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "language_id"
    t.index ["post_id", "language_id"], name: "posts_languages_idx", unique: true
  end

  create_table "posts_links", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "link_id"
    t.index ["post_id", "link_id"], name: "posts_links_idx", unique: true
  end

  create_table "posts_mentions", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "actor_id"
    t.index ["post_id", "actor_id"], name: "posts_mentions_idx", unique: true
  end

  create_table "posts_tags", id: false, force: :cascade do |t|
    t.string "post_id"
    t.string "tag_id"
    t.index ["post_id", "tag_id"], name: "posts_tags_idx", unique: true
  end

  create_table "tags", id: false, force: :cascade do |t|
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.string "id"
    t.string "label"
    t.index ["id"], name: "tags_idx", unique: true
    t.index ["label"], name: "labels_idx", unique: true
  end
end
