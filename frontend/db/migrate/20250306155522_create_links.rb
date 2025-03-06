class CreateLinks < ActiveRecord::Migration[8.0]
  def change
    create_table :links, id: :string do |t|
      t.timestamps
      t.string :original_url
      t.string :url
      t.integer :count

      t.index ["url"], name: "links_urls_col_idx"
    end

    create_table :links_posts, id: false do |t|
      t.belongs_to :link, type: :string
      t.belongs_to :post, type: :string
      t.index [:link_id, :post_id], name: "links_posts_idx", unique: true 
      t.index [:post_id], name: "links_post_id_posts_idx" 
    end
  end
end
