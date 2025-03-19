class CreateTopPastWeeks < ActiveRecord::Migration[8.0]
  def change
    create_view :top_past_weeks, materialized: true
  end
end
