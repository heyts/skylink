class CreateMonthlyStats < ActiveRecord::Migration[8.0]
  def change
    create_table :monthly_stats do |t|
      t.timestamps
    end
  end
end
