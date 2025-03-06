class CreateDailyStats < ActiveRecord::Migration[8.0]
  def change
    create_table :daily_stats do |t|
      t.timestamps
    end
  end
end
