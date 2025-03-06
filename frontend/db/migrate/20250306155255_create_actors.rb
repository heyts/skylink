class CreateActors < ActiveRecord::Migration[8.0]
  def change
    create_table :actors, id: :string, force: :cascade do |t|
      t.timestamps

      t.string :display_name
      t.string :handle 
      t.string :avatar
      t.string :banner

      t.integer :followers_count
      t.integer :follows_count
      t.integer :posts_count

      t.index ["handle"], name: "actor_handle_idx"

    end
  end
end
