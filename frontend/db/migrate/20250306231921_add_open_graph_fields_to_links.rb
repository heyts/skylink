class AddOpenGraphFieldsToLinks < ActiveRecord::Migration[8.0]
  def change
    add_column :links, :title, :string
    add_column :links, :og_title, :string
    add_column :links, :og_description, :text
    add_column :links, :og_site_name, :string
    add_column :links, :og_image, :string
    add_column :links, :og_image_options, :json
    add_column :links, :og_optional, :json
  end
end
