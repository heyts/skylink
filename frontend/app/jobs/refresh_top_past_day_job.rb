class RefreshTopPastDayJob < ApplicationJob
  queue_as :default

  def perform(*args)
    TopPastDay.refresh
  end
end
