class UpdateTopPastHoursToVersion4 < ActiveRecord::Migration[8.0]
  def change
    update_view :top_past_hours,
      version: 4,
      revert_to_version: 3,
      materialized: true
  end
end
