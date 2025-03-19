class UpdateTopPastHoursToVersion3 < ActiveRecord::Migration[8.0]
  def change
    update_view :top_past_hours,
      version: 3,
      revert_to_version: 2,
      materialized: true
  end
end
