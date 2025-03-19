class CreatePosts < ActiveRecord::Migration[8.0]
  def change
    create_table :posts, id: :string do |t|
      t.timestamps
      t.timestamp :published_at
      t.string :collection
      t.string :record_key
      t.text :text
      t.string :language
      t.string :country
      t.string :locale
      t.string :tags, array: true 
      t.references :actor, type: :string
      t.virtual :text_fts, type: :tsvector, as: "to_tsvector('simple', text)", stored: true

      t.index ["published_at"], name: "published_at_col_idx"
      t.index ["language"], name: "posts_language_col_idx"
      t.index ["country"], name: "posts_country_col_idx"
      t.index ["record_key"], name: "posts_record_key_col_idx", unique: true
      t.index ["text_fts"], using: :gin, name: "text_fts_idx"
      t.index ["tags"], name: "posts_tags_col_idx", using: :gin
    end
    
    create_table :mentions_posts, id: false do |t|
      t.belongs_to :post, type: :string
      t.belongs_to :actor, type: :string
      t.index [:post_id, :actor_id], name: "mentions_posts_idx", unique: true
    end
  
  end
end
