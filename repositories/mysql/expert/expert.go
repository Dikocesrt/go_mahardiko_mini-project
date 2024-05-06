package expert

import (
	expertEntities "habit/entities/expert"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ExpertRepo struct {
	DB *gorm.DB
}

func NewExpertRepo(db *gorm.DB) *ExpertRepo {
	return &ExpertRepo{
		DB: db,
	}
}

func (expertRepo *ExpertRepo) Register(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	expertDb := FromExpertEntitiesToExpertDb(expert)
	bankAccountDb := FromBankAccountEntitiesToBankAccountDb(expert)

	err := expertRepo.DB.Create(&bankAccountDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	expertDb.BankAccountId = bankAccountDb.Id
	err = expertRepo.DB.Create(&expertDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	var expertiseDb Expertise
	err = expertRepo.DB.Where("id = ?", expertDb.ExpertiseId).First(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	newExpert := expertDb.FromExpertDbToExpertEntities()
	newBankAccount := bankAccountDb.FromBankAccountDbToBankAccountEntities()
	newExpertise := expertiseDb.FromExpertiseDbToExpertiseEntities()

	newExpert.BankAccount = *newBankAccount
	newExpert.Expertise = *newExpertise
	return *newExpert, nil
}

func (expertRepo *ExpertRepo) Login(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	expertDb := FromExpertEntitiesToExpertDb(expert)

	password := expertDb.Password
	err := expertRepo.DB.Where("Username = ?", expertDb.Username).First(&expertDb).Error
	if err != nil {
		err := expertRepo.DB.Where("Email = ?", expertDb.Username).First(&expertDb).Error
		if err != nil {
			return expertEntities.Expert{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(expertDb.Password), []byte(password))
	if err != nil {
		return expertEntities.Expert{}, err
	}

	expertFromDb := expertDb.FromExpertDbToExpertEntities()
	return *expertFromDb, nil
}