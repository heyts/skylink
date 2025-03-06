class Link < ApplicationRecord
    has_and_belongs_to_many :posts
    
    has_many :actors, -> { distinct }, through: :posts
    has_many :mentions, -> { distinct }, through: :posts

    has_many :hourly_stats, through: :posts
    has_many :daily_stats, through: :posts
    has_many :weekly_stats, through: :posts
    has_many :monthly_stats, through: :posts

def pretty_title
    !og_title.blank? ? og_title : title
end

def locale
    og_optional["og:optional"] ? og_optional["og:optional"] : ""
end

def self.top(lang: "en", limit: 20, tags: [], since: 7.days.ago)
    top = self
        .select(
            :id,
            :title,
            :og_title,
            :og_description, 
            :url, 
            "COUNT(distinct posts.id) posts_count", 
            "COUNT(distinct actor_id) actor_count"
        )
        .joins(posts:[:actor, :languages, :tags])
        .where("posts.created_at > ? and languages.id = ?", since, lang)
        .group(:id, :url)
        .distinct
        .order(Arel.sql("COUNT(distinct actor_id) DESC"))
        .limit(limit)
    
    if !tags.blank?
        top = top.where("tags.label IN (?)", tags)
    end
    top
end
end


