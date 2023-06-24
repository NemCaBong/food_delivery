package restaurantstorage

import (
	"context"
	restaurantmodel "github.com/nemcabong/food_delivery/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	/* nếu không về con trỏ thì sẽ không trả về giá trị nil được
	mà phải về restaurantmodel.Restaurant{} => 1 struct rỗng thì các field
	của nó sẽ có default value => tốn mem */
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
