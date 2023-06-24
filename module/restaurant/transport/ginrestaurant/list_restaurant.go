package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcabong/food_delivery/common"
	"github.com/nemcabong/food_delivery/component/appctx"
	restaurantbiz "github.com/nemcabong/food_delivery/module/restaurant/biz"
	restaurantmodel "github.com/nemcabong/food_delivery/module/restaurant/model"
	restaurantstorage "github.com/nemcabong/food_delivery/module/restaurant/storage"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := ctx.ShouldBind(&pagingData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		pagingData.Fulfill() // nhiệm vụ của transport

		var filter restaurantmodel.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var result []restaurantmodel.Restaurant

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(ctx.Request.Context(), &filter, &pagingData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.NewSucessResponse(result, filter, pagingData))
	}
}
