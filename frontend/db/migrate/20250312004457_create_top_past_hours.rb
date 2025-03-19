class CreateTopPastHours < ActiveRecord::Migration[8.0]
  def change
    create_view :top_past_hours, materialized: true
  end
end
