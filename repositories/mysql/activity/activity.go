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

	var activityTypeTemp ActivityType
	err = activityRepo.DB.Where("id = ?", activity.ActivityTypeId).First(&activityTypeTemp).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	newActivityEnt := activityDb.FromActivityDbToActivityEntities()
	newActivityDetailEnt := activityDetailDb.FromActivityDbToActivityDetailEntities()
	newActivityType := activityTypeTemp.FromActivityDbToActivityTypeEntities()

	newActivityDetailEnt.Id = activityDb.ActivityDetailId
	newActivityEnt.ActivityDetail = *newActivityDetailEnt
	newActivityEnt.ActivityType = *newActivityType

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

func (activityRepo *ActivityRepo) GetActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	var activityDb Activity
	
	err := activityRepo.DB.Where("id = ?", activity.Id).First(&activityDb).Error
	if err != nil {
		fmt.Println(err)
		return activityEntities.Activity{}, err
	}

	var activityDetailDb ActivityDetail
	err = activityRepo.DB.Where("id = ?", activityDb.ActivityDetailId).First(&activityDetailDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	var activityTypeDb ActivityType
	err = activityRepo.DB.Where("id = ?", activityDb.ActivityTypeId).First(&activityTypeDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	newActivityEnt := activityDb.FromActivityDbToActivityEntities()
	newActivityDetailEnt := activityDetailDb.FromActivityDbToActivityDetailEntities()
	newActivityType := activityTypeDb.FromActivityDbToActivityTypeEntities()

	newActivityEnt.ActivityDetail = *newActivityDetailEnt
	newActivityEnt.ActivityType = *newActivityType

	return *newActivityEnt, nil
}

func (activityRepo *ActivityRepo) UpdateActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	activityDb := FromActivityEntitiesToActivityDb(activity)
	activityDb.Id = activity.Id
	activityDetailDb := FromActivityEntitiesToActivityDetailDb(activity)

	var activityTemp Activity

	err := activityRepo.DB.Where("id = ?", activity.Id).First(&activityTemp).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	activityDb.CreatedAt = activityTemp.CreatedAt
	activityDb.UserId = activityTemp.UserId

	activityDb.ActivityDetailId = activityTemp.ActivityDetailId
	err = activityRepo.DB.Save(&activityDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	err = activityRepo.DB.First(&activityDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	activityDetailDb.Id = activityDb.ActivityDetailId
	var activityDetailTemp ActivityDetail

	err = activityRepo.DB.Where("id = ?", activityDetailDb.Id).First(&activityDetailTemp).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	if activityDetailDb.ImageUrl == "" {
		activityDetailDb.ImageUrl = activityDetailTemp.ImageUrl
	}

	err = activityRepo.DB.Save(&activityDetailDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	err = activityRepo.DB.First(&activityDetailDb).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	var activityTypeTemp ActivityType
	err = activityRepo.DB.Where("id = ?", activity.ActivityTypeId).First(&activityTypeTemp).Error
	if err != nil {
		return activityEntities.Activity{}, err
	}

	newActivityEnt := activityDb.FromActivityDbToActivityEntities()
	newActivityDetailEnt := activityDetailDb.FromActivityDbToActivityDetailEntities()
	newActivityType := activityTypeTemp.FromActivityDbToActivityTypeEntities()

	newActivityEnt.ActivityDetail = *newActivityDetailEnt
	newActivityEnt.ActivityType = *newActivityType

	return *newActivityEnt, nil
}

func (activityRepo *ActivityRepo) DeleteActivityById(activity activityEntities.Activity) error {
	var activityDb Activity
	activityDb.Id = activity.Id

	err := activityRepo.DB.Where("id = ?", activityDb.Id).First(&activityDb).Error
	if err != nil {
		return err
	}

	var activityDetailDb ActivityDetail 
	activityDetailDb.Id = activityDb.ActivityDetailId

	err = activityRepo.DB.Delete(&activityDetailDb).Error
	if err != nil {
		return err
	}

	err = activityRepo.DB.Delete(&activityDb).Error
	if err != nil {
		return err
	}

	return nil
}

func (ActivityRepo *ActivityRepo) CreateActivityType(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	activityTypeDb := ActivityType {
		Name:        activityType.Name,
		Description: activityType.Description,
	}

	err := ActivityRepo.DB.Create(&activityTypeDb).Error
	if err != nil {
		return activityEntities.ActivityType{}, err
	}

	activityEnt := activityEntities.ActivityType{
		Id:          activityTypeDb.Id,
		Name:        activityTypeDb.Name,
		Description: activityTypeDb.Description,
	}

	return activityEnt, nil
}

func (activityRepo *ActivityRepo) GetAllActivityType() ([]activityEntities.ActivityType, error) {
	var activityTypeDbs []ActivityType
	err := activityRepo.DB.Find(&activityTypeDbs).Error
	if err != nil {
		return []activityEntities.ActivityType{}, err
	}

	activityTypeEnts := make([]activityEntities.ActivityType, len(activityTypeDbs))

	for i := 0; i < len(activityTypeDbs); i++ {
		activityTypeEnts[i] = activityEntities.ActivityType{
			Id:          activityTypeDbs[i].Id,
			Name:        activityTypeDbs[i].Name,
			Description: activityTypeDbs[i].Description,
		}
	}

	return activityTypeEnts, nil
}

func (activityRepo *ActivityRepo) GetActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	var activityTypeDb ActivityType
	activityTypeDb.Id = activityType.Id

	err := activityRepo.DB.Where("id = ?", activityTypeDb.Id).First(&activityTypeDb).Error
	if err != nil {
		return activityEntities.ActivityType{}, err
	}

	activityTypeEnt := activityEntities.ActivityType{
		Id:          activityTypeDb.Id,
		Name:        activityTypeDb.Name,
		Description: activityTypeDb.Description,
	}

	return activityTypeEnt, nil
}

func (activityRepo *ActivityRepo) UpdateActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	var activityTypeDb ActivityType
	activityTypeDb.Id = activityType.Id
	activityTypeDb.Name = activityType.Name
	activityTypeDb.Description = activityType.Description

	err := activityRepo.DB.Save(&activityTypeDb).Error
	if err != nil {
		return activityEntities.ActivityType{}, err
	}

	return activityType, nil
}

func (activityRepo *ActivityRepo) DeleteActivityTypeById(activityType activityEntities.ActivityType) error {
	var activityTypeDb ActivityType
	activityTypeDb.Id = activityType.Id

	err := activityRepo.DB.Delete(&activityTypeDb).Error
	if err != nil {
		return err
	}

	return nil
}