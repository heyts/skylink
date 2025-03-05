class HourlyStat < ApplicationRecord
    belongs_to :post
end


# HourlyStat hourly_stats (request for < 24h)
# DailyStat daily_stats (request for < 7d)
# WeeklyStat weekly_stats (request for < 30d)
# MonthlyStat monthly_stats



