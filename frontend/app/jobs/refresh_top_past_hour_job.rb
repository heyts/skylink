class RefreshTopPastHourJob < ApplicationJob
  queue_as :default

  def perform(*args)
    TopPastHour.refresh
  end
end
