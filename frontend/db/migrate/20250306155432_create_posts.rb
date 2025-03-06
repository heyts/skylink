class CreatePosts < ActiveRecord::Migration[8.0]
  def change
    create_table :posts, id: :string do |t|
      t.timestamps
      t.timestamp :published_at
      t.string :collection
      t.string :record_key
      t.text :text
      t.string :language
      t.string :tags, array: true 
      t.references :actor, type: :string

      t.index ["language"], name: "posts_language_col_idx"
    end
    add_index :posts, :tags, name: "posts_tags_col_idx", using: "gin"
  end
end
