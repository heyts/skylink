class Link < ApplicationRecord
    has_and_belongs_to_many :posts

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
    Actor.joins(posts: [ :mentions, :links ]).where("links.id = ?", self.id).distinct
end
end
