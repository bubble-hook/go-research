package model

//UserReponse UserReponse
type UserReponse struct {
	ID uint   `gorm:"primarykey" json:"id"`
	X1 string `json:"x1" binding:"required"`
	X2 string `json:"x2" binding:"required"`
	X3 string `json:"x3" binding:"required"`
}

//User User
type User struct {
	ID          uint   `gorm:"primarykey" json:"id"`
	UserName    string `json:"userName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
}

//ParseToReponse ParseToReponse
func (u User) ParseToReponse() *UserReponse {
	return &UserReponse{
		X1: u.UserName,
		X2: u.Password,
		X3: u.DisplayName,
	}

}
