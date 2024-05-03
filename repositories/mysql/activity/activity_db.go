package activity

import (
	activityEntities "habit/entities/activity"

	"gorm.io/gorm"
)
type Activity struct {
	gorm.Model
	Id               int `gorm:"primaryKey;autoIncrement"`
	Title            string
	ActivityStart    string
	ActivityFinish   string
	ActivityTypeId   int `gorm:"index"`
	ActivityDetailId int `gorm:"index"`
	UserId           int `gorm:"index"`
}

type ActivityType struct {
	gorm.Model
	Id          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
}

type ActivityDetail struct {
	gorm.Model
	Id             int `gorm:"primaryKey;autoIncrement"`
	HeartRate      int
	Intensity      string
	CaloriesBurned float64
	FoodDetails    string
	ImageUrl       string
}

func FromActivityEntitiesToActivityDb(activity activityEntities.Activity) *Activity {
	return &Activity{
		Title:            activity.Title,
		ActivityStart:    activity.ActivityStart,
		ActivityFinish:   activity.ActivityFinish,
		ActivityTypeId:   activity.ActivityTypeId,
		UserId:           activity.UserId,
	}
}

func FromActivityEntitiesToActivityDetailDb(activity activityEntities.Activity) *ActivityDetail {
	return &ActivityDetail{
		HeartRate:      activity.ActivityDetail.HeartRate,
		Intensity:      activity.ActivityDetail.Intensity,
		CaloriesBurned: activity.ActivityDetail.CaloriesBurned,
		FoodDetails:    activity.ActivityDetail.FoodDetails,
		ImageUrl:       activity.ActivityDetail.ImageUrl,
	}
}

func (activity *Activity) FromActivityDbToActivityEntities() *activityEntities.Activity {
	return &activityEntities.Activity{
		Id:             activity.Id,
		Title:          activity.Title,
		ActivityStart:  activity.ActivityStart,
		ActivityFinish: activity.ActivityFinish,
		UserId:         activity.UserId,
		ActivityTypeId: activity.ActivityTypeId,
	}
}

func (activity *ActivityDetail) FromActivityDbToActivityDetailEntities() *activityEntities.ActivityDetail {
	return &activityEntities.ActivityDetail{
		Id:             activity.Id,
		HeartRate:      activity.HeartRate,
		Intensity:      activity.Intensity,
		CaloriesBurned: activity.CaloriesBurned,
		FoodDetails:    activity.FoodDetails,
		ImageUrl:       activity.ImageUrl,
	}
}

func (activity *ActivityType) FromActivityDbToActivityTypeEntities() *activityEntities.ActivityType {
	return &activityEntities.ActivityType{
		Id:          activity.Id,
		Name:        activity.Name,
		Description: activity.Description,
	}
}