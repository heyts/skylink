class Post < ApplicationRecord
    belongs_to :actor
    has_and_belongs_to_many :mentions, class_name: 'Actor', join_table: "posts_mentions"
    has_and_belongs_to_many :links, join_table: "posts_links"
    has_and_belongs_to_many :tags, join_table: "posts_tags"
    has_and_belongs_to_many :languages, join_table: "posts_languages"
end
