class HomeController < ApplicationController
  def index
    redirect_to toplist_index_url(:day), status: 301
  end
end
