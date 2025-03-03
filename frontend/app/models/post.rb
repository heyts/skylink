class Post < ApplicationRecord
    belongs_to :actor
    has_and_belongs_to_many :mentions, class_name: "Actor", join_table: "mentions_posts", counter_cache: true
    has_and_belongs_to_many :links, counter_cache: true
    has_and_belongs_to_many :tags, counter_cache: true
    has_and_belongs_to_many :languages, counter_cache: true

    def url
        "#{Rails.application.config.x.skylink.bsky_web_url}/profile/#{self.actor.handle}/post/#{self.record_key}"
    end
end
