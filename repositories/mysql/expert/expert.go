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

func (expertRepo *ExpertRepo) UpdateProfileExpertById(expert *expertEntities.Expert) (expertEntities.Expert, int64, error) {
	expertDb := FromExpertEntitiesToExpertDb(expert)
	bankAccountDb := FromBankAccountEntitiesToBankAccountDb(expert)

	var expertDbTemp Expert
	err := expertRepo.DB.Where("Id = ?", expertDb.Id).First(&expertDbTemp).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	var bankAccountDbTemp BankAccount
	err = expertRepo.DB.Where("Id = ?", expertDbTemp.BankAccountId).First(&bankAccountDbTemp).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	bankAccountDb.Id = expertDbTemp.BankAccountId
	bankAccountDb.BankAccountTypeId = bankAccountDbTemp.BankAccountTypeId
	expertDb.BankAccountId = expertDbTemp.BankAccountId
	expertDb.ExpertiseId = expertDbTemp.ExpertiseId

	err = expertRepo.DB.Save(&bankAccountDb).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	if expertDb.ProfilePicture == "" {
		expertDb.ProfilePicture = expertDbTemp.ProfilePicture
	}

	var counterUsername, counterEmail int64
	err = expertRepo.DB.Model(&expertDb).Where("Username = ?", expertDb.Username).Count(&counterUsername).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	if expertDb.Username != expertDbTemp.Username && counterUsername > 0 {
		return expertEntities.Expert{}, 2, err
	}

	if expertDb.Username == expertDbTemp.Username && counterUsername > 1 {
		return expertEntities.Expert{}, 2, err
	}

	err = expertRepo.DB.Model(&expertDb).Where("Email = ?", expertDb.Email).Count(&counterEmail).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	if expertDb.Email != expertDbTemp.Email && counterEmail > 0 {
		return expertEntities.Expert{}, 3, err
	}

	if expertDb.Email == expertDbTemp.Email && counterEmail > 1 {
		return expertEntities.Expert{}, 3, err
	}

	err = expertRepo.DB.Save(&expertDb).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	err = expertRepo.DB.Where("Id = ?", expertDb.Id).First(&expertDb).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	err = expertRepo.DB.Where("Id = ?", expertDb.BankAccountId).First(&bankAccountDb).Error
	if err != nil {
		return expertEntities.Expert{}, 1, err
	}

	expertFromDb := expertDb.FromExpertDbToExpertEntities()
	newBankAccount := bankAccountDb.FromBankAccountDbToBankAccountEntities()
	expertFromDb.BankAccount = *newBankAccount
	return *expertFromDb, 0, nil
}

func (expertRepo *ExpertRepo) GetAllExperts() ([]expertEntities.Expert, error) {
	var experts []Expert
	err := expertRepo.DB.Find(&experts).Error
	if err != nil {
		return []expertEntities.Expert{}, err
	}

	var expertise []Expertise
	err = expertRepo.DB.Find(&expertise).Error
	if err != nil {
		return []expertEntities.Expert{}, err
	}

	var bankAccount []BankAccount
	err = expertRepo.DB.Find(&bankAccount).Error
	if err != nil {
		return []expertEntities.Expert{}, err
	}

	expertEntities := make([]expertEntities.Expert, len(experts))

	for i := 0; i < len(experts); i++ {
		expertEntities[i] = *experts[i].FromExpertDbToExpertEntities()
		expertEntities[i].Expertise = *expertise[i].FromExpertiseDbToExpertiseEntities()
		expertEntities[i].BankAccount = *bankAccount[i].FromBankAccountDbToBankAccountEntities()
	}

	return expertEntities, nil
}

func (expertRepo *ExpertRepo) GetExpertById(expert *expertEntities.Expert) (expertEntities.Expert, error) {
	var expertDb Expert
	expertDb.Id = expert.Id
	err := expertRepo.DB.First(&expertDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	var expertiseDb Expertise
	expertiseDb.Id = expertDb.ExpertiseId
	err = expertRepo.DB.First(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	var bankAccountDb BankAccount
	bankAccountDb.Id = expertDb.BankAccountId
	err = expertRepo.DB.First(&bankAccountDb).Error
	if err != nil {
		return expertEntities.Expert{}, err
	}

	expertFromDb := expertDb.FromExpertDbToExpertEntities()
	expertFromDb.Expertise = *expertiseDb.FromExpertiseDbToExpertiseEntities()
	expertFromDb.BankAccount = *bankAccountDb.FromBankAccountDbToBankAccountEntities()
	return *expertFromDb, nil
}

func (expertRepo *ExpertRepo) CreateExpertise(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb Expertise
	expertiseDb.Name = expertise.Name
	expertiseDb.Description = expertise.Description
	err := expertRepo.DB.Create(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	expertise.Id = expertiseDb.Id
	return expertise, nil
}

func (expertRepo *ExpertRepo) GetAllExpertise() ([]expertEntities.Expertise, error) {
	var expertiseDb []Expertise
	err := expertRepo.DB.Find(&expertiseDb).Error
	if err != nil {
		return []expertEntities.Expertise{}, err
	}
	expertise := make([]expertEntities.Expertise, len(expertiseDb))
	for i := 0; i < len(expertiseDb); i++ {
		expertise[i].Id = expertiseDb[i].Id
		expertise[i].Name = expertiseDb[i].Name
		expertise[i].Description = expertiseDb[i].Description
	}
	return expertise, nil
}

func (expertRepo *ExpertRepo) GetExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb Expertise
	expertiseDb.Id = expertise.Id
	err := expertRepo.DB.First(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	expertise.Name = expertiseDb.Name
	expertise.Description = expertiseDb.Description
	return expertise, nil
}

func (expertRepo *ExpertRepo) UpdateExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb Expertise
	expertiseDb.Id = expertise.Id
	expertiseDb.Name = expertise.Name
	expertiseDb.Description = expertise.Description
	err := expertRepo.DB.Save(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	return expertise, nil
}

func (expertRepo *ExpertRepo) DeleteExpertiseById(expertise expertEntities.Expertise) error {
	var expertiseDb Expertise
	expertiseDb.Id = expertise.Id
	err := expertRepo.DB.Delete(&expertiseDb).Error
	if err != nil {
		return err
	}
	return nil
}

func (expertRepo *ExpertRepo) CreateBankAccountType(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	var bankTypeDb BankAccountType
	bankTypeDb.Name = bankType.Name
	err := expertRepo.DB.Create(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, err
	}
	bankType.Id = bankTypeDb.Id
	return bankType, nil
}

func (expertRepo *ExpertRepo) GetBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	var bankTypeDb BankAccountType
	bankTypeDb.Id = bankType.Id
	err := expertRepo.DB.First(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, err
	}
	bankType.Name = bankTypeDb.Name
	return bankType, nil
}

func (expertRepo *ExpertRepo) GetAllBankAccountType() ([]expertEntities.BankAccountType, error) {
	var bankTypesDb []BankAccountType
	err := expertRepo.DB.Find(&bankTypesDb).Error
	if err != nil {
		return []expertEntities.BankAccountType{}, err
	}

	bankTypes := make([]expertEntities.BankAccountType, len(bankTypesDb))

	for i := 0; i < len(bankTypesDb); i++ {
		bankTypes[i].Id = bankTypesDb[i].Id
		bankTypes[i].Name = bankTypesDb[i].Name
	}

	return bankTypes, nil
}

func (expertRepo *ExpertRepo) UpdateBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, int64, error) {
	var bankTypeDb BankAccountType
	bankTypeDb.Id = bankType.Id
	bankTypeDb.Name = bankType.Name

	var counter int64
	err := expertRepo.DB.Model(&bankTypeDb).Where("Id = ?", bankType.Id).Count(&counter).Error
	if err != nil {
		return expertEntities.BankAccountType{}, 0, err
	}

	if counter == 0 {
		return expertEntities.BankAccountType{}, 1, nil
	}

	err = expertRepo.DB.Save(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, 0, err
	}
	return bankType, 0, nil
}

func (expertRepo *ExpertRepo) DeleteBankAccountTypeById(bankType expertEntities.BankAccountType) error {
	var bankTypeDb BankAccountType
	bankTypeDb.Id = bankType.Id
	err := expertRepo.DB.Delete(&bankTypeDb).Error
	if err != nil {
		return err
	}
	return nil
}