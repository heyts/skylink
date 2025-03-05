class Post < ApplicationRecord
    belongs_to :actor
    has_and_belongs_to_many :mentions, class_name: "Actor", join_table: "mentions_posts"
    has_and_belongs_to_many :links
    has_and_belongs_to_many :tags
    has_and_belongs_to_many :languages

    has_many :hourly_stats
    has_many :daily_stats
    has_many :weekly_stats
    has_many :monthly_stats

    def url
        "#{Rails.application.config.x.skylink.bsky_web_url}/profile/#{self.actor.handle}/post/#{self.record_key}"
    end
end
