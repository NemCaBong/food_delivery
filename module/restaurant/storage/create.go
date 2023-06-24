package restaurantstorage

import (
	"context"
	restaurantmodel "github.com/nemcabong/food_delivery/module/restaurant/model"
)

// Tất cả những hàm có IO nên truyền context vào để tracing truy gốc, cancel.

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
