class CreateHourlyStats < ActiveRecord::Migration[8.0]
  def change
    create_table :hourly_stats, id: false do |t|
      t.timestamp :ymdh
      t.references :post, type: :string
      t.integer :likes_count, default: 0
      t.integer :reposts_count, default: 0
      t.integer :quotes_count, default: 0

      t.index [:ymdh, :post_id], name: "hourly_stats_pk_col_idx", unique: true
    end
  end
end
