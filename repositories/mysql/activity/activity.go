package activity

import (
	"fmt"
	activityEntities "habit/entities/activity"

	"gorm.io/gorm"
)

type ActivityRepo struct {
	DB *gorm.DB
}

func NewActivityRepo(db *gorm.DB) *ActivityRepo {
	return &ActivityRepo{
		DB: db,
	}
}

func (activityRepo *ActivityRepo) CreateActivity(activity activityEntities.Activity) (activityEntities.Activity, error) {
	activityDb := FromActivityEntitiesToActivityDb(activity)
	activityDetailDb := FromActivityEntitiesToActivityDetailDb(activity)

	err := activityRepo.DB.Create(&activityDetailDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	err = activityRepo.DB.First(&activityDetailDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	activityDb.ActivityDetailId = activityDetailDb.Id

	err = activityRepo.DB.Create(&activityDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	var activityTypeTemp activityEntities.ActivityType
	err = activityRepo.DB.Where("id = ?", activity.ActivityTypeId).First(&activityTypeTemp).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	newActivityEnt := activityDb.FromActivityDbToActivityEntities()
	newActivityDetailEnt := activityDetailDb.FromActivityDbToActivityDetailEntities()

	newActivityDetailEnt.Id = activityDb.ActivityDetailId
	newActivityEnt.ActivityDetail = *newActivityDetailEnt
	newActivityEnt.ActivityType = activityTypeTemp

	return *newActivityEnt, nil
}

func (activityRepo *ActivityRepo) GetActivityByUserId(userId int) ([]activityEntities.Activity, error) {
	var activities []Activity
	err := activityRepo.DB.Where("user_id = ?", userId).Find(&activities).Error
	if err != nil {
		fmt.Println(err)
		return []activityEntities.Activity{}, err
	}

	activityDetails := make([]ActivityDetail, len(activities))

	for i:=0;i<len(activities);i++ {
		err = activityRepo.DB.Where("id = ?", activities[i].ActivityDetailId).First(&activityDetails[i]).Error
		if err != nil {
			fmt.Println(err)
			return []activityEntities.Activity{}, err
		}
	}

	activityTypes := make([]ActivityType, len(activities))

	for i:=0;i<len(activities);i++ {
		err = activityRepo.DB.Where("id = ?", activities[i].ActivityTypeId).First(&activityTypes[i]).Error
		if err != nil {
			fmt.Println(err)
			return []activityEntities.Activity{}, err
		}
	}

	activitiesEntities := make([]activityEntities.Activity, len(activities))

	for i := 0;i<len(activities);i++ {
		activitiesEntities[i] = *activities[i].FromActivityDbToActivityEntities()
		activitiesEntities[i].ActivityDetail = *activityDetails[i].FromActivityDbToActivityDetailEntities()
		activitiesEntities[i].ActivityType = *activityTypes[i].FromActivityDbToActivityTypeEntities()
	}

	return activitiesEntities, nil
}