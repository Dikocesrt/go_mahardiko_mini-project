package activity

import (
	activityEntities "habit/entities/activity"

	"gorm.io/gorm"
)
type Activity struct {
	gorm.Model
	Id               int `gorm:"primaryKey;autoIncrement"`
	Title            string
	ActivityStart    string `gorm:"type:time"`
	ActivityFinish   string `gorm:"type:time"`
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

func (activityDb *ActivityDetail) FromActivityDbToActivityDetailEntities() *activityEntities.ActivityDetail {
	return &activityEntities.ActivityDetail{
		Id:             activityDb.Id,
		HeartRate:      activityDb.HeartRate,
		Intensity:      activityDb.Intensity,
		CaloriesBurned: activityDb.CaloriesBurned,
		FoodDetails:    activityDb.FoodDetails,
		ImageUrl:       activityDb.ImageUrl,
	}
}

func (activityType *ActivityType) FromActivityDbToActivityTypeEntities() *activityEntities.ActivityType {
	return &activityEntities.ActivityType{
		Id:          activityType.Id,
		Name:        activityType.Name,
		Description: activityType.Description,
	}
}