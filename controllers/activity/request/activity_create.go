package request

type ActivityCreateRequest struct {
	Title          string         `json:"title" form:"title"`
	ActivityStart  string         `json:"activity_start" form:"activity_start"`
	ActivityFinish string         `json:"activity_finish" form:"activity_finish"`
	UserId         int            `json:"user_id" form:"user_id"`
	ActivityTypeId int            `json:"activity_type_id" form:"activity_type_id"`
	ActivityDetail ActivityDetail `json:"activity_detail"`
}

type ActivityDetail struct {
	Id             int     `json:"id" form:"id"`
	HeartRate      int     `json:"heart_rate" form:"heart_rate"`
	Intensity      string  `json:"intensity" form:"intensity"`
	CaloriesBurned float64 `json:"calories_burned" form:"calories_burned"`
	FoodDetails    string  `json:"food_details" form:"food_details"`
	ImageUrl       string  `json:"food_images" form:"food_images"`
}