package admin

import (
	adminEntities "habit/entities/admin"
	expertEntities "habit/entities/expert"
	expertDb "habit/repositories/mysql/expert"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func NewAdminRepo(db *gorm.DB) *AdminRepo {
	return &AdminRepo{
		DB: db,
	}
}

func (adminRepo *AdminRepo) Register(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	var adminDb Admin
	adminDb.Username = admin.Username
	adminDb.Email = admin.Email
	adminDb.Password = admin.Password
	err := adminRepo.DB.Create(&adminDb).Error
	if err != nil {
		return adminEntities.Admin{}, err
	}
	admin.Id = adminDb.Id
	return *admin, nil
}

func (adminRepo *AdminRepo) Login(admin *adminEntities.Admin) (adminEntities.Admin, error) {
	var adminDb Admin
	adminDb.Username = admin.Username
	adminDb.Password = admin.Password

	err := adminRepo.DB.Where("Username = ?", adminDb.Username).First(&adminDb).Error
	if err != nil {
		err := adminRepo.DB.Where("Email = ?", adminDb.Username).First(&adminDb).Error
		if err != nil {
			return adminEntities.Admin{}, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminDb.Password), []byte(admin.Password))
	if err != nil {
		return adminEntities.Admin{}, err
	}
	admin.Id = adminDb.Id
	admin.Username = adminDb.Username
	admin.Email = adminDb.Email
	return *admin, nil
}

func (adminRepo *AdminRepo) CreateBankAccountType(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	var bankTypeDb expertDb.BankAccountType
	bankTypeDb.Name = bankType.Name
	err := adminRepo.DB.Create(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, err
	}
	bankType.Id = bankTypeDb.Id
	return bankType, nil
}

func (adminRepo *AdminRepo) GetBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	var bankTypeDb expertDb.BankAccountType
	bankTypeDb.Id = bankType.Id
	err := adminRepo.DB.First(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, err
	}
	bankType.Name = bankTypeDb.Name
	return bankType, nil
}

func (adminRepo *AdminRepo) GetAllBankAccountType() ([]expertEntities.BankAccountType, error) {
	var bankTypesDb []expertDb.BankAccountType
	err := adminRepo.DB.Find(&bankTypesDb).Error
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

func (adminRepo *AdminRepo) UpdateBankAccountTypeById(bankType expertEntities.BankAccountType) (expertEntities.BankAccountType, error) {
	var bankTypeDb expertDb.BankAccountType
	bankTypeDb.Id = bankType.Id
	bankTypeDb.Name = bankType.Name
	err := adminRepo.DB.Save(&bankTypeDb).Error
	if err != nil {
		return expertEntities.BankAccountType{}, err
	}
	return bankType, nil
}

func (adminRepo *AdminRepo) DeleteBankAccountTypeById(bankType expertEntities.BankAccountType) error {
	var bankTypeDb expertDb.BankAccountType
	bankTypeDb.Id = bankType.Id
	err := adminRepo.DB.Delete(&bankTypeDb).Error
	if err != nil {
		return err
	}
	return nil
}

func (adminRepo *AdminRepo) CreateExpertise(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb expertDb.Expertise
	expertiseDb.Name = expertise.Name
	expertiseDb.Description = expertise.Description
	err := adminRepo.DB.Create(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	expertise.Id = expertiseDb.Id
	return expertise, nil
}

func (adminRepo *AdminRepo) GetAllExpertise() ([]expertEntities.Expertise, error) {
	var expertiseDb []expertDb.Expertise
	err := adminRepo.DB.Find(&expertiseDb).Error
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

func (adminRepo *AdminRepo) GetExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb expertDb.Expertise
	expertiseDb.Id = expertise.Id
	err := adminRepo.DB.First(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	expertise.Name = expertiseDb.Name
	expertise.Description = expertiseDb.Description
	return expertise, nil
}

func (adminRepo *AdminRepo) UpdateExpertiseById(expertise expertEntities.Expertise) (expertEntities.Expertise, error) {
	var expertiseDb expertDb.Expertise
	expertiseDb.Id = expertise.Id
	expertiseDb.Name = expertise.Name
	expertiseDb.Description = expertise.Description
	err := adminRepo.DB.Save(&expertiseDb).Error
	if err != nil {
		return expertEntities.Expertise{}, err
	}
	return expertise, nil
}

func (adminRepo *AdminRepo) DeleteExpertiseById(expertise expertEntities.Expertise) error {
	var expertiseDb expertDb.Expertise
	expertiseDb.Id = expertise.Id
	err := adminRepo.DB.Delete(&expertiseDb).Error
	if err != nil {
		return err
	}
	return nil
}