package hire

import (
	hireEntities "habit/entities/hire"
	"habit/repositories/mysql/expert"
	"habit/repositories/mysql/user"

	"gorm.io/gorm"
)

type HireRepo struct {
	DB *gorm.DB
}

func NewHireRepo(db *gorm.DB) *HireRepo {
	return &HireRepo{
		DB: db,
	}
}

func (hireRepo *HireRepo) CreateHire(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	hireDb := FromHireEntitiesToHireDb(hire)

	err := hireRepo.DB.Create(&hireDb).Error

	if err != nil {
		return hireEntities.Hire{}, err
	}

	var userDb user.User
	err = hireRepo.DB.Where("id = ?", hireDb.UserId).First(&userDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	userEnt := userDb.FromUserDbToUserEntities()

	var expertDb expert.Expert
	err = hireRepo.DB.Where("id = ?", hireDb.ExpertId).First(&expertDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	expertEnt := expertDb.FromExpertDbToExpertEntities()

	hireEnt := hireDb.FromHireDbToHireEntities()
	hireEnt.User = *userEnt
	hireEnt.Expert = *expertEnt

	return *hireEnt, nil
}

func (hireRepo *HireRepo) GetHiresByExpertId(id int) ([]hireEntities.Hire, error) {
	var hiresDb []Hire
	err := hireRepo.DB.Where("expert_id = ?", id).Find(&hiresDb).Error
	if err != nil {
		return []hireEntities.Hire{}, err
	}

	usersDb := make([]user.User, len(hiresDb))

	for i := 0; i < len(hiresDb); i++ {
		err = hireRepo.DB.Where("id = ?", hiresDb[i].UserId).First(&usersDb[i]).Error
		if err != nil {
			return []hireEntities.Hire{}, err
		}
	}

	var expertDb expert.Expert
	err = hireRepo.DB.Where("id = ?", id).First(&expertDb).Error
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

func (hireRepo *HireRepo) GetHiresByUserId(id int) ([]hireEntities.Hire, error) {
	var hiresDb []Hire
	err := hireRepo.DB.Where("user_id = ?", id).Find(&hiresDb).Error
	if err != nil {
		return []hireEntities.Hire{}, err
	}

	expertsDb := make([]expert.Expert, len(hiresDb))

	for i := 0; i < len(hiresDb); i++ {
		err = hireRepo.DB.Where("id = ?", hiresDb[i].ExpertId).First(&expertsDb[i]).Error
		if err != nil {
			return []hireEntities.Hire{}, err
		}
	}

	var userDb user.User
	err = hireRepo.DB.Where("id = ?", id).First(&userDb).Error
	if err != nil {
		return []hireEntities.Hire{}, err
	}

	hiresEnt := make([]hireEntities.Hire, len(hiresDb))

	for i := 0; i < len(hiresDb); i++ {
		hiresEnt[i] = *hiresDb[i].FromHireDbToHireEntities()
		hiresEnt[i].User = *userDb.FromUserDbToUserEntities()
		hiresEnt[i].Expert = *expertsDb[i].FromExpertDbToExpertEntities()
	}

	return hiresEnt, nil
}

func (hireRepo *HireRepo) GetHireById(id int) (hireEntities.Hire, error) {
	var hireDb Hire

	err := hireRepo.DB.Where("id = ?", id).First(&hireDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	var userDb user.User
	err = hireRepo.DB.Where("id = ?", hireDb.UserId).First(&userDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	var expertDb expert.Expert
	err = hireRepo.DB.Where("id = ?", hireDb.ExpertId).First(&expertDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	hireEnt := hireDb.FromHireDbToHireEntities()
	hireEnt.User = *userDb.FromUserDbToUserEntities()
	hireEnt.Expert = *expertDb.FromExpertDbToExpertEntities()

	return *hireEnt, nil
}

func (hireRepo *HireRepo) VerifyPayment(hire *hireEntities.Hire) (hireEntities.Hire, error) {
	var hireDbTemp Hire
	hireDbTemp.Id = hire.Id
	err := hireRepo.DB.First(&hireDbTemp).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	hireDbTemp.PaymentStatus = hire.PaymentStatus
	hireDbTemp.MeetUrl = hire.MeetUrl
	hireDbTemp.HireStart = hire.HireStart
	hireDbTemp.HireEnd = hire.HireEnd

	err = hireRepo.DB.Save(&hireDbTemp).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	var userDb user.User
	err = hireRepo.DB.Where("id = ?", hireDbTemp.UserId).First(&userDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	var expertDb expert.Expert
	err = hireRepo.DB.Where("id = ?", hireDbTemp.ExpertId).First(&expertDb).Error
	if err != nil {
		return hireEntities.Hire{}, err
	}

	hireEnt := hireDbTemp.FromHireDbToHireEntities()
	hireEnt.User = *userDb.FromUserDbToUserEntities()
	hireEnt.Expert = *expertDb.FromExpertDbToExpertEntities()

	return *hireEnt, nil
}