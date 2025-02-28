class Link < ApplicationRecord
    has_and_belongs_to_many :posts, join_table: "posts_links"
    # has_many :languages, through: :posts
    # has_many :tags, through: :posts
    # has_many :mentions, through: :posts

def _actors
    @actors = []
    self.posts.includes(:actor).each do |post| 
        @actors << post.actor
    end
    @actors.uniq
end

def actors
    Actor.joins(posts: :links).where("links.id = ?", self.id).distinct
end

def tags
    Tag.joins(posts: :links).where("links.id = ?", self.id).distinct
end

def languages
    Language.joins(posts: :links).where("links.id = ?", self.id).distinct
end

def mentions
    Actor.joins(posts: [:mentions, :links]).where("links.id = ?", self.id).distinct
end

end
