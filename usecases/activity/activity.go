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
		return activityEntities.Activity{}, constants.ErrEmptyInputActivity
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
		return activityEntities.Activity{}, constants.ErrInsertDatabase
	}

	return activity, nil
}

func (activityUseCase *ActivityUseCase) GetActivityByUserId(userId int) ([]activityEntities.Activity, error) {
	activities, kode, err := activityUseCase.repository.GetActivityByUserId(userId)
	if err != nil {
		return []activityEntities.Activity{}, constants.ErrGetAllData
	}

	if kode == 1 {
		return []activityEntities.Activity{}, constants.ErrGetActivitiesByUserId
	}
	return activities, nil
}

func (activityUseCase *ActivityUseCase) GetActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	activity, err := activityUseCase.repository.GetActivityById(activity)
	if err != nil {
		return activityEntities.Activity{}, constants.ErrActivityNotFound
	}
	return activity, nil
}

func (activityUseCase *ActivityUseCase) UpdateActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	if activity.Title == "" || activity.ActivityStart == "" || activity.ActivityFinish == "" {
		return activityEntities.Activity{}, constants.ErrEmptyInputActivity
	}

	activity, err := activityUseCase.repository.UpdateActivityById(activity)
	if err != nil {
		return activityEntities.Activity{}, constants.ErrActivityNotFound
	}
	return activity, nil
}

func (activityUseCase *ActivityUseCase) DeleteActivityById(activity activityEntities.Activity) error {
	err := activityUseCase.repository.DeleteActivityById(activity)
	if err != nil {
		return constants.ErrActivityNotFound
	}
	return nil
}

func (activityUseCase *ActivityUseCase) CreateActivityType(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	if activityType.Name == "" || activityType.Description == "" {
		return activityEntities.ActivityType{}, constants.ErrEmptyInputActivityType
	}
	activityType, err := activityUseCase.repository.CreateActivityType(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrInsertDatabase
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) GetAllActivityType() ([]activityEntities.ActivityType, error) {
	activityTypes, err := activityUseCase.repository.GetAllActivityType()
	if err != nil {
		return []activityEntities.ActivityType{}, constants.ErrGetAllData
	}
	return activityTypes, nil
}

func (activityUseCase *ActivityUseCase) GetActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	activityType, err := activityUseCase.repository.GetActivityTypeById(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrActivityTypeNotFound
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) UpdateActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	if activityType.Name == "" || activityType.Description == "" {
		return activityEntities.ActivityType{}, constants.ErrEmptyInputActivityType
	}
	activityType, kode, err := activityUseCase.repository.UpdateActivityTypeById(activityType)
	if err != nil {
		return activityEntities.ActivityType{}, constants.ErrUpdateData
	}

	if kode == 1 {
		return activityEntities.ActivityType{}, constants.ErrActivityTypeNotFound
	}
	return activityType, nil
}

func (activityUseCase *ActivityUseCase) DeleteActivityTypeById(activityType activityEntities.ActivityType) error {
	kode, err := activityUseCase.repository.DeleteActivityTypeById(activityType)
	if err != nil {
		return constants.ErrDeleteData
	}

	if kode == 1 {
		return constants.ErrActivityTypeNotFound
	}
	return nil
}