class UpdateTopPastHoursToVersion5 < ActiveRecord::Migration[8.0]
  def change
    update_view :top_past_hours,
      version: 5,
      revert_to_version: 4,
      materialized: true
  end
end
