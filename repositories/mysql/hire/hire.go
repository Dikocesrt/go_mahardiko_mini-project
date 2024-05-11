package hire

import (
	hireEntities "habit/entities/hire"
	"habit/repositories/mysql/expert"
	"habit/repositories/mysql/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (userRepo *UserRepo) CreateHire(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	hireDb := FromHireEntitiesToHireDb(hire)

	err := userRepo.DB.Create(&hireDb).Error

	if err != nil {
		return hireEntities.Hire{}, err
	}

	var userDb user.User
	err = userRepo.DB.Where("id = ?", hireDb.UserId).First(&userDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	userEnt := userDb.FromUserDbToUserEntities()

	var expertDb expert.Expert
	err = userRepo.DB.Where("id = ?", hireDb.ExpertId).First(&expertDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	expertEnt := expertDb.FromExpertDbToExpertEntities()

	hireEnt := hireDb.FromHireDbToHireEntities()
	hireEnt.User = *userEnt
	hireEnt.Expert = *expertEnt

	return *hireEnt, nil
}

func (userRepo *UserRepo) GetHiresByExpertId(id int) ([]hireEntities.Hire, error) {
	var hiresDb []Hire
	err := userRepo.DB.Where("expert_id = ?", id).Find(&hiresDb).Error
	if err != nil {
		return []hireEntities.Hire{}, err
	}

	usersDb := make([]user.User, len(hiresDb))

	for i := 0; i < len(hiresDb); i++ {
		err = userRepo.DB.Where("id = ?", hiresDb[i].UserId).First(&usersDb[i]).Error
		if err != nil {
			return []hireEntities.Hire{}, err
		}
	}

	var expertDb expert.Expert
	err = userRepo.DB.Where("id = ?", id).First(&expertDb).Error
	if err != nil {
		return []hireEntities.Hire{}, err
	}

	hiresEnt := make([]hireEntities.Hire, len(hiresDb))

	for i := 0; i < len(hiresDb); i++ {
		hiresEnt[i] = *hiresDb[i].FromHireDbToHireEntities()
		hiresEnt[i].User = *usersDb[i].FromUserDbToUserEntities()
		hiresEnt[i].Expert = *expertDb.FromExpertDbToExpertEntities()
	}

	return hiresEnt, nil
}