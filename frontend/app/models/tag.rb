class Tag < ApplicationRecord
    self.primary_key = "id"
    has_and_belongs_to_many :posts

    def links
        Link.joins(posts: :tags).where("tags.id = ?", self.id)
    end
end
