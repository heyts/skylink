class UpdateTopPastWeeksToVersion2 < ActiveRecord::Migration[8.0]
  def change
    update_view :top_past_weeks,
      version: 2,
      revert_to_version: 1,
      materialized: true
  end
end
