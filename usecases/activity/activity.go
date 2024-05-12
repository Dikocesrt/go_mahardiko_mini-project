package activity

import (
	"context"
	"habit/configs"
	"habit/constants"
	activityEntities "habit/entities/activity"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type ActivityUseCase struct {
	repository activityEntities.RepositoryInterface
}

func NewActivityUseCase(repository activityEntities.RepositoryInterface) *ActivityUseCase {
	return &ActivityUseCase{
		repository: repository,
	}
}

func uploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	cloudinaryURL := configs.InitConfigCloudinary()
	if cloudinaryURL == "" {
		return "", constants.ErrCloudinary
	}

	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func (activityUseCase *ActivityUseCase) CreateActivity(activity activityEntities.Activity, file *multipart.FileHeader) (activityEntities.Activity, error) {
	if activity.Title == "" || activity.ActivityStart == "" || activity.ActivityFinish == "" {
		return activityEntities.Activity{}, constants.ErrEmptyInputCreateActivity
	}

	if file != nil {
		SecureURL, err := uploadImage(file)
		if err != nil {
			return activityEntities.Activity{}, constants.ErrUploadImage
		}

		activity.ActivityDetail.ImageUrl = SecureURL
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
		return []activityEntities.Activity{}, constants.ErrGetDatabase
	}
	return activities, nil
}

func (activityUseCase *ActivityUseCase) GetActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	activity, err := activityUseCase.repository.GetActivityById(activity)
	if err != nil {
		return activityEntities.Activity{}, constants.ErrGetDatabase
	}
	return activity, nil
}

func (activityUseCase *ActivityUseCase) UpdateActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	if activity.Title == "" || activity.ActivityStart == "" || activity.ActivityFinish == "" {
		return activityEntities.Activity{}, constants.ErrEmptyInputCreateActivity
	}

	activity, err := activityUseCase.repository.UpdateActivityById(activity)
	if err != nil {
		return activityEntities.Activity{}, constants.ErrUpdateDatabase
	}
	return activity, nil
}

func (activityUseCase *ActivityUseCase) DeleteActivityById(activity activityEntities.Activity) error {
	err := activityUseCase.repository.DeleteActivityById(activity)
	if err != nil {
		return constants.ErrDeleteDatabase
	}
	return nil
}

func (activityUseCase *ActivityUseCase) CreateActivityType(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	activityType, err := activityUseCase.repository.CreateActivityType(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrGetDatabase
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) GetAllActivityType() ([]activityEntities.ActivityType, error) {
	activityTypes, err := activityUseCase.repository.GetAllActivityType()
	if err != nil {
		return []activityEntities.ActivityType{}, constants.ErrGetDatabase
	}
	return activityTypes, nil
}

func (activityUseCase *ActivityUseCase) GetActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	activityType, err := activityUseCase.repository.GetActivityTypeById(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrGetDatabase
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) UpdateActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	activityType, err := activityUseCase.repository.UpdateActivityTypeById(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrGetDatabase
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) DeleteActivityTypeById(activityType activityEntities.ActivityType) error {
	err := activityUseCase.repository.DeleteActivityTypeById(activityType)
	if err != nil {
		return constants.ErrDeleteDatabase
	}
	return nil
}