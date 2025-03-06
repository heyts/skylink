class CreateLinks < ActiveRecord::Migration[8.0]
  def change
    create_table :links do |t|
      t.timestamps
    end
  end
end
