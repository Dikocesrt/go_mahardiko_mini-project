package activity

type Activity struct {
	Id             int
	Title          string
	ActivityStart  string
	ActivityFinish string
	ActivityTypeId int
	ActivityType   ActivityType
	ActivityDetail ActivityDetail
	UserId         int
}

type ActivityType struct {
	Id          int
	Name        string
	Description string
}

type ActivityDetail struct {
	Id             int
	HeartRate      int
	Intensity      string
	CaloriesBurned float64
	FoodDetails    string
	ImageUrl       string
}

type RepositoryInterface interface {
	CreateActivity(activity Activity) (Activity, error)
	GetActivityByUserId(userId int) ([]Activity, error)
	UpdateActivityById(activity Activity) (Activity, error)
}

type UseCaseInterface interface {
	CreateActivity(activity Activity) (Activity, error)
	GetActivityByUserId(userId int) ([]Activity, error)
	UpdateActivityById(activity Activity) (Activity, error)
}