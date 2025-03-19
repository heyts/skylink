class UpdateTopPastHoursToVersion6 < ActiveRecord::Migration[8.0]
  def change
    update_view :top_past_hours,
      version: 6,
      revert_to_version: 5,
      materialized: true
  end
end
