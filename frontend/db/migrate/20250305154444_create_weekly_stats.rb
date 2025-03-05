class CreateWeeklyStats < ActiveRecord::Migration[8.0]
  def change
    create_table :weekly_stats do |t|
      t.timestamps
    end
  end
end
