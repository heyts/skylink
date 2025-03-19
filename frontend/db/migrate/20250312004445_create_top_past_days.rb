class CreateTopPastDays < ActiveRecord::Migration[8.0]
  def change
    create_view :top_past_days, materialized: true
  end
end
