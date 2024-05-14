package activity

import (
	"errors"
	activityEntities "habit/entities/activity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockActivityRepository struct {
	Err error
	Kode int64
}

func (m *MockActivityRepository) CreateActivity(activity activityEntities.Activity) (activityEntities.Activity, error) {
	return activity, m.Err
}

func (m *MockActivityRepository) GetActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	return activity, m.Err
}

func (m *MockActivityRepository) GetActivityByUserId(userId int) ([]activityEntities.Activity, int64, error) {
	return []activityEntities.Activity{}, m.Kode, m.Err
}

func (m *MockActivityRepository) UpdateActivityById(activity activityEntities.Activity) (activityEntities.Activity, error) {
	return activity, m.Err
}

func (m *MockActivityRepository) DeleteActivityById(activity activityEntities.Activity) error {
	return m.Err
}

func (m *MockActivityRepository) CreateActivityType(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	return activityType, m.Err
}

func (m *MockActivityRepository) GetActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, error) {
	return activityType, m.Err
}

func (m *MockActivityRepository) GetAllActivityType() ([]activityEntities.ActivityType, error) {
	return []activityEntities.ActivityType{}, m.Err
}

func (m *MockActivityRepository) UpdateActivityTypeById(activityType activityEntities.ActivityType) (activityEntities.ActivityType, int64, error) {
	return activityType, m.Kode, m.Err
}

func (m *MockActivityRepository) DeleteActivityTypeById(activityType activityEntities.ActivityType) (int64, error) {
	return m.Kode, m.Err
}

func TestCreateActivity(t *testing.T) {
	t.Run("Create Activity Success", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Title:        "Hiking",
			ActivityStart: "2022-01-01 00:00:00",
			ActivityFinish: "2022-01-02 00:00:00",
		}

		createdActivity, err := activityUseCase.CreateActivity(activity, nil)

		assert.NoError(t, err)
		assert.Equal(t, activity, createdActivity)
	})

	t.Run("Create Activity Empty Input", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}
		
		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{}

		createdActivity, err := activityUseCase.CreateActivity(activity, nil)

		assert.Error(t, err)
		assert.Equal(t, activity, createdActivity)
	})

	t.Run("Create Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to create activity")}
		
		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Title:        "Hiking",
			ActivityStart: "2022-01-01 00:00:00",
			ActivityFinish: "2022-01-02 00:00:00",
		}

		createdActivity, err := activityUseCase.CreateActivity(activity, nil)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.Activity{}, createdActivity)
	})
}


func TestGetActivityById(t *testing.T) {
	t.Run("Get Activity Success", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}

		activityUseCase := NewActivityUseCase(mockActivityRepository)
		
		activity := activityEntities.Activity{
			Id: 1,
		}

		getActivity, err := activityUseCase.GetActivityById(activity)
		
		assert.NoError(t, err)
		assert.Equal(t, activity, getActivity)
	})

	t.Run("Get Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to get activity")}
		
		activityUseCase := NewActivityUseCase(mockActivityRepository)
		
		activity := activityEntities.Activity{
			Id: 10,
		}
		
		getActivity, err := activityUseCase.GetActivityById(activity)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.Activity{}, getActivity)
	})
}


func TestGetActivityByUserId(t *testing.T) {
	t.Run("Get Activity Success", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		getActivity, err := activityUseCase.GetActivityByUserId(1)
		
		assert.NoError(t, err)
		assert.Equal(t, []activityEntities.Activity{}, getActivity)
	})

	t.Run("Get Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to get activity")}

		activityUseCase := NewActivityUseCase(mockActivityRepository)
		
		activity := activityEntities.Activity{
			UserId: 10,
		}

		getActivities, err := activityUseCase.GetActivityByUserId(activity.UserId)

		assert.Error(t, err)
		assert.Equal(t, []activityEntities.Activity{}, getActivities)
	})

	t.Run("Get Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to get activity"), Kode: 1}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			UserId: 10,
		}

		getActivities, err := activityUseCase.GetActivityByUserId(activity.UserId)

		assert.Error(t, err)
		assert.Equal(t, []activityEntities.Activity{}, getActivities)
	})
}


func TestUpdateActivityById(t *testing.T) {
	t.Run("Update Activity Success", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}
		
		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Id: 10,
			Title: "Hiking",
			ActivityStart: "2022-01-01 00:00:00",
			ActivityFinish: "2022-01-02 00:00:00",
		}

		updatedActivity, err := activityUseCase.UpdateActivityById(activity)

		assert.NoError(t, err)
		assert.Equal(t, activity, updatedActivity)
	})

	t.Run("Update Activity Empty Input", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{}

		updatedActivity, err := activityUseCase.UpdateActivityById(activity)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.Activity{}, updatedActivity)
	})

	t.Run("Update Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to update activity")}
		
		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Id: 10,
			Title: "Hiking",
			ActivityStart: "2022-01-01 00:00:00",
			ActivityFinish: "2022-01-02 00:00:00",
		}

		updatedActivity, err := activityUseCase.UpdateActivityById(activity)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.Activity{}, updatedActivity)
	})
}


func TestDeleteActivityById(t *testing.T) {
	t.Run("Delete Activity Success", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: nil}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Id: 1,
		}

		err := activityUseCase.DeleteActivityById(activity)

		assert.NoError(t, err)
	})

	t.Run("Delete Activity Failed", func(t *testing.T) {
		mockActivityRepository := &MockActivityRepository{Err: errors.New("failed to delete activity")}

		activityUseCase := NewActivityUseCase(mockActivityRepository)

		activity := activityEntities.Activity{
			Id: 10,
		}

		err := activityUseCase.DeleteActivityById(activity)

		assert.Error(t, err)
	})
}

func TestCreateActivityType(t *testing.T) {
	t.Run("Create Activity Type Success", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Name: "Hiking",
			Description: "Hiking",
		}

		createdActivityType, err := activityTypeUseCase.CreateActivityType(activityType)

		assert.NoError(t, err)
		assert.Equal(t, activityType, createdActivityType)
	})

	t.Run("Create Activity Type Empty Input", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{}

		createdActivityType, err := activityTypeUseCase.CreateActivityType(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, createdActivityType)
	})

	t.Run("Create Activity Type Failed", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: errors.New("failed to create activity type")}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Name: "Hiking",
			Description: "Hiking",
		}

		createdActivityType, err := activityTypeUseCase.CreateActivityType(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, createdActivityType)
	})
}

func TestGetAllActivityType(t *testing.T) {
	t.Run("Get All Activity Type Success", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityTypes, err := activityTypeUseCase.GetAllActivityType()

		assert.NoError(t, err)
		assert.Equal(t, []activityEntities.ActivityType{}, activityTypes)
	})

	t.Run("Get All Activity Type Failed", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: errors.New("failed to get all activity type")}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityTypes, err := activityTypeUseCase.GetAllActivityType()

		assert.Error(t, err)
		assert.Equal(t, []activityEntities.ActivityType{}, activityTypes)
	})
}

func TestGetActivityTypeById(t *testing.T) {
	t.Run("Get Activity Type Success", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
		}

		getActivityType, err := activityTypeUseCase.GetActivityTypeById(activityType)

		assert.NoError(t, err)
		assert.Equal(t, activityType, getActivityType)
	})

	t.Run("Get Activity Type Failed", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: errors.New("failed to get activity type")}
		
		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
		}

		getActivityType, err := activityTypeUseCase.GetActivityTypeById(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, getActivityType)
	})
}

func TestUpdateActivityTypeById(t *testing.T){
	t.Run("Update Activity Type Success", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
			Name: "Hiking",
			Description: "Hiking",
		}

		updatedActivityType, err := activityTypeUseCase.UpdateActivityTypeById(activityType)

		assert.NoError(t, err)
		assert.Equal(t, activityType, updatedActivityType)
	})

	t.Run("Update Activity Type Empty Input", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{}

		updatedActivityType, err := activityTypeUseCase.UpdateActivityTypeById(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, updatedActivityType)
	})

	t.Run("Update Activity Type Failed", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: errors.New("failed to update activity type")}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
			Name: "Hiking",
			Description: "Hiking",
		}

		updatedActivityType, err := activityTypeUseCase.UpdateActivityTypeById(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, updatedActivityType)
	})

	t.Run("Update Activity Type Not Found", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil,
		Kode: 1,}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
			Name: "Hiking",
			Description: "Hiking",
		}

		updatedActivityType, err := activityTypeUseCase.UpdateActivityTypeById(activityType)

		assert.Error(t, err)
		assert.Equal(t, activityEntities.ActivityType{}, updatedActivityType)
	})
}

func TestDeleteActivityTypeById(t *testing.T){
	t.Run("Delete Activity Type Success", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
		}

		err := activityTypeUseCase.DeleteActivityTypeById(activityType)

		assert.NoError(t, err)
	})

	t.Run("Delete Activity Type Failed", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: errors.New("failed to delete activity type")}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 1,
		}

		err := activityTypeUseCase.DeleteActivityTypeById(activityType)

		assert.Error(t, err)
	})

	t.Run("Delete Activity Type Not Found", func(t *testing.T) {
		mockActivityTypeRepository := &MockActivityRepository{Err: nil, Kode: 1}

		activityTypeUseCase := NewActivityUseCase(mockActivityTypeRepository)

		activityType := activityEntities.ActivityType{
			Id: 10,
		}

		err := activityTypeUseCase.DeleteActivityTypeById(activityType)

		assert.Error(t, err)
	})
}