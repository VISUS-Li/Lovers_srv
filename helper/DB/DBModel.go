package DB

import "github.com/jinzhu/gorm"

type UserBaseInfo struct{
	gorm.Model
	UserId 		string `gorm:"unique;not null"`
	RealName 	string
	Phone 		string
	Sex 		int /*1.男，2.女*/
	Birth 		string
	Sculpture	string /*头像地址*/
	HomeTown	string /*所在地*/
	LoverId			string
	LoverPhone		string
	LoverNickName 	string /*对另一半的昵称*/
	LoveDuration	int64 `gorm:type:BIGINT`
}

type LoginInfo struct{
	gorm.Model
	UserId string  `gorm:"unique;not null"`
	UserName string
	PassWord string  `gorm:"not null"`
	Phone string	 `gorm:"unique;not null"`
}