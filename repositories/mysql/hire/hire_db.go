package hire

import (
	hireEntities "habit/entities/hire"
	"time"

	"gorm.io/gorm"
)

type Hire struct {
	gorm.Model
	Id            int `gorm:"primaryKey:autoIncrement"`
	HireStart     time.Time `gorm:"type:date"`
	HireEnd       time.Time `gorm:"type:date"`
	TotalFee      int
	PaymentStatus string `gorm:"type:ENUM('pending', 'paid')"`
	PaymentImage  string
	MeetTime      string `gorm:"type:time"`
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