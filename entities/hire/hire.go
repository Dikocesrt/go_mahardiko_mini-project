package hire

import (
	"habit/entities/expert"
	"habit/entities/user"
	"mime/multipart"
)

type Hire struct {
	Id            int
	HireStart     string
	HireEnd       string
	TotalFee      int
	PaymentStatus string
	PaymentImage  string
	MeetTime      string
	MeetDay       string
	MeetUrl       string
	UserId        int
	ExpertId      int
	User          user.User
	Expert        expert.Expert
}

type RepositoryInterface interface {
	CreateHire(hire *Hire) (Hire, error)
	GetHiresByExpertId(id int) ([]Hire, error)
}

type UseCaseInterface interface {
	CreateHire(hire *Hire, file *multipart.FileHeader) (Hire, error)
	GetHiresByExpertId(id int) ([]Hire, error)
}