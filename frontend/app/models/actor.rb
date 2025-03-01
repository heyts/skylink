class Actor < ApplicationRecord
    has_many :posts
    def url
        "#{Rails.application.config.x.skylink.bsky_web_url}/profile/#{self.handle}"
    end
end
