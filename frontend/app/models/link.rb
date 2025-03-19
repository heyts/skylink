class Link < ApplicationRecord
    has_and_belongs_to_many :posts
    
    has_many :actors, -> { distinct }, through: :posts
    has_many :mentions, -> { distinct }, through: :posts

    has_many :hourly_stats, through: :posts
    has_many :daily_stats, through: :posts
    has_many :weekly_stats, through: :posts
    has_many :monthly_stats, through: :posts

    has_one :top_past_hour
    has_one :top_past_day
    has_one :top_past_week

def pretty_title
    !og_title.blank? ? og_title : title
end

def locale
    og_optional["og:optional"] ? og_optional["og:optional"] : ""
end

def domain
    og_site_name.blank? ? url.split("/")[2].split('.').last(2).join('.') : og_site_name 
end

def self.top(period: :hour, lang: "en", limit: 20, tags: [], since: 7.days.ago)
    periods = {
        hour: TopPastHour,
        day: TopPastDay,
        week: TopPastWeek,
    }
    
    if not periods.key? period
        period = :hour
    end 

    rel = periods[period].name.underscore
    self
        .select("links.*", "#{rel}.score")
        .includes(:top_past_week)
        .joins(rel.to_sym)
        .where("#{rel}.language = ?", lang)
        .order(Arel.sql("#{rel}.score DESC NULLS LAST"))
        .limit(limit)
end
end


