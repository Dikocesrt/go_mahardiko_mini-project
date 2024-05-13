package activity

import "mime/multipart"

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
	GetActivityByUserId(userId int) ([]Activity, int64, error)
	GetActivityById(activity Activity) (Activity, error)
	UpdateActivityById(activity Activity) (Activity, error)
	DeleteActivityById(activity Activity) error
	CreateActivityType(activityType ActivityType) (ActivityType, error)
	GetAllActivityType() ([]ActivityType, error)
	GetActivityTypeById(activityType ActivityType) (ActivityType, error)
	UpdateActivityTypeById(activityType ActivityType) (ActivityType, int64, error)
	DeleteActivityTypeById(activityType ActivityType) (int64, error)
}

type UseCaseInterface interface {
	CreateActivity(activity Activity, file *multipart.FileHeader) (Activity, error)
	GetActivityByUserId(userId int) ([]Activity, error)
	GetActivityById(activity Activity) (Activity, error)
	UpdateActivityById(activity Activity) (Activity, error)
	DeleteActivityById(activity Activity) error
	CreateActivityType(activityType ActivityType) (ActivityType, error)
	GetAllActivityType() ([]ActivityType, error)
	GetActivityTypeById(activityType ActivityType) (ActivityType, error)
	UpdateActivityTypeById(activityType ActivityType) (ActivityType, error)
	DeleteActivityTypeById(activityType ActivityType) error
}