package hire

import (
	hireEntities "habit/entities/hire"

	"gorm.io/gorm"
)

type Hire struct {
	gorm.Model
	Id            int `gorm:"primaryKey:autoIncrement"`
	HireStart     string
	HireEnd       string
	TotalFee      int
	PaymentStatus string
	PaymentImage  string
	MeetTime      string
	MeetDay       string
	MeetUrl       string
	UserId        int `gorm:"index"`
	ExpertId      int `gorm:"index"`
}

func FromHireEntitiesToHireDb(hire *hireEntities.Hire) *Hire {
	return &Hire{
		Id:            hire.Id,
		HireStart:     hire.HireStart,
		HireEnd:       hire.HireEnd,
		TotalFee:      hire.TotalFee,
		PaymentStatus: hire.PaymentStatus,
		PaymentImage:  hire.PaymentImage,
		MeetTime:      hire.MeetTime,
		MeetDay:       hire.MeetDay,
		MeetUrl:       hire.MeetUrl,
		UserId:        hire.UserId,
		ExpertId:      hire.ExpertId,
	}
}

func (hire *Hire) FromHireDbToHireEntities() *hireEntities.Hire {
	return &hireEntities.Hire{
		Id:            hire.Id,
		HireStart:     hire.HireStart,
		HireEnd:       hire.HireEnd,
		TotalFee:      hire.TotalFee,
		PaymentStatus: hire.PaymentStatus,
		PaymentImage:  hire.PaymentImage,
		MeetTime:      hire.MeetTime,
		MeetDay:       hire.MeetDay,
		MeetUrl:       hire.MeetUrl,
		UserId:        hire.UserId,
		ExpertId:      hire.ExpertId,
	}
}