class HomeController < ApplicationController
  def index
    @links = Link.top(lang: "en", limit: 20, since:24.hour.ago)
  end

  def show
    id = params[:id]
    @link = Link.find(id)
    @posts = @link.posts.all
  end
end
