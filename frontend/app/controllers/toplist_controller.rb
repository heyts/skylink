class ToplistController < ApplicationController
    def index
        @period = :day
        if params[:period].present?
            @period = params[:period].to_sym
        end
        @top = Link.top(period: @period, lang: "en", limit: 20)
    end

    def show
        id = params[:id]
        @period = params[:period] || :day
        @link = Link.find(id)
        @posts = @link.posts.all
      end
end
