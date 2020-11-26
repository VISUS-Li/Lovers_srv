package DB

import "github.com/jinzhu/gorm"

type UserBaseInfo struct{
	gorm.Model
	UserId 		string `gorm:"unique;not null" json:"user_id"`
	RealName 	string `json:"real_name"`
	Phone 		string `json:"phone"`
	Sex 		int  `json:"set""`/*1.男，2.女*/
	Birth 		string `json:"birth""`
	Sculpture	string  `json:"sculpture"`/*头像地址*/
	HomeTown	string  `json:"home_town"`/*所在地*/
	LoverId			string `json:"lover_id"`
	LoverPhone		string `json:"lover_phone"`
	LoverNickName 	string `json:"lover_nick_name"`/*对另一半的昵称*/
	LoveDuration	int64 `gorm:type:BIGINT json:"lover_duration"`
	CoupleId		string `gorm:"unique" json:"couple_id"` /*情侣ID，绑定情侣后，双方有一个共同的ID*/
}

type LoginInfo struct{
	gorm.Model
	UserId string  `gorm:"unique;not null" json:"user_id"`
	PassWord string  `gorm:"not null" json:"pass_word"`
	Phone string	 `gorm:"unique;not null" json:"phone"`
}