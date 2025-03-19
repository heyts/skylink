class RefreshTopPastWeekJob < ApplicationJob
  queue_as :default

  def perform(*args)
    TopPastWeek.refresh
  end
end
