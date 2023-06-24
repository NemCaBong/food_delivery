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

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data restaurantmodel.RestaurantCreate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
