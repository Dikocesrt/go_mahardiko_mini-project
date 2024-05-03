package activity

import (
	"habit/constants"
	activityEntities "habit/entities/activity"
)

type ActivityUseCase struct {
	repository activityEntities.RepositoryInterface
}

func NewActivityUseCase(repository activityEntities.RepositoryInterface) *ActivityUseCase {
	return &ActivityUseCase{
		repository: repository,
	}
}

func (activityUseCase *ActivityUseCase) CreateActivity(activity activityEntities.Activity) (activityEntities.Activity, error) {
	if activity.Title == "" || activity.ActivityStart == "" || activity.ActivityFinish == "" {
		return activityEntities.Activity{}, constants.ErrEmptyInputCreateActivity
	}

	activity, err := activityUseCase.repository.CreateActivity(activity)
	if err != nil {
		return activityEntities.Activity{}, err
	}

	return activity, nil
}

func (activityUseCase *ActivityUseCase) GetActivityByUserId(userId int) ([]activityEntities.Activity, error) {
	activities, err := activityUseCase.repository.GetActivityByUserId(userId)
	if err != nil {
		return []activityEntities.Activity{}, err
	}
	return activities, nil
}

func (activityUseCase *ActivityUseCase) UpdateActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	if activity.Title == "" || activity.ActivityStart == "" || activity.ActivityFinish == "" {
		return activityEntities.Activity{}, constants.ErrEmptyInputCreateActivity
	}

	activity, err := activityUseCase.repository.UpdateActivityById(activity)
	if err != nil {
		return activityEntities.Activity{}, err
	}
	return activity, nil
}