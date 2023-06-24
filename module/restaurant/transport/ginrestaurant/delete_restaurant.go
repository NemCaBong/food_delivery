package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcabong/food_delivery/common"
	"github.com/nemcabong/food_delivery/component/appctx"
	restaurantbiz "github.com/nemcabong/food_delivery/module/restaurant/biz"
	restaurantstorage "github.com/nemcabong/food_delivery/module/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
