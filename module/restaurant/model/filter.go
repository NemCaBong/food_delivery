package restaurantmodel

// Quy định những tham số client muốn lọc data thêm

type Filter struct {
	OwnerId int `json:"owner_id,omitempty" form:"owner_id"`
}
