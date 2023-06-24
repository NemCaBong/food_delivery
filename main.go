package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nemcabong/food_delivery/component/appctx"
	"github.com/nemcabong/food_delivery/module/restaurant/transport/ginrestaurant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	ID   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

// TableName let the GORM know what table is this struct represent
func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	ID   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	// dsn := "nemcabong:thobeogalaxy257@tcp(0.0.0.0:3308)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug()

	var myRestaurant Restaurant
	if err := db.Where("id = ?", 5).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println(myRestaurant)

	appContext := appctx.NewAppContext(db)
	r := gin.Default()
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	// Get list restaurant
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))

	// Get restaurant by id
	restaurants.GET("/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		var data Restaurant

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Where("id = ?", id).First(&data)

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// Update table
	restaurants.PATCH("/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var data RestaurantUpdate

		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		var result Restaurant
		db.Where("id = ?", id).First(&result)

		ctx.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	err = r.Run(":8081")

	if err != nil {
		log.Println(err)
	}
}
