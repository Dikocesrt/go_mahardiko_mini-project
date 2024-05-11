package admin

type Admin struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Username string
	Email    string
	Password string
}