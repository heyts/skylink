class Language < ApplicationRecord
    self.primary_key = "id"
    has_and_belongs_to_many :posts, join_table: "posts_languages"
end
