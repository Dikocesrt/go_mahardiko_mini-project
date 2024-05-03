package activity

import (
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