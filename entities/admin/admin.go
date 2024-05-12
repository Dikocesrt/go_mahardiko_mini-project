package admin

type Admin struct {
	Id       int
	Username string
	Email    string
	Password string
	Token    string
}

type RepositoryInterface interface {
	Register(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
}

type UseCaseInterface interface {
	Register(admin *Admin) (Admin, error)
	Login(admin *Admin) (Admin, error)
}