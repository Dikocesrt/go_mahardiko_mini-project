package response

type Activity struct {
	Id             int            `json:"id"`
	Title          string         `json:"title"`
	ActivityStart  string         `json:"activity_start"`
	ActivityFinish string         `json:"activity_finish"`
	UserId         int            `json:"user_id"`
	ActivityDetail ActivityDetail `json:"activity_detail"`
	ActivityType   ActivityType   `json:"activity_type"`
}

type ActivityType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ActivityDetail struct {
	HeartRate      int     `json:"heart_rate"`
	Intensity      string  `json:"intensity"`
	CaloriesBurned float64 `json:"calories_burned"`
	FoodDetails    string  `json:"food_details"`
	ImageUrl       string  `json:"image_url"`
}