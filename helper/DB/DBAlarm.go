package DB

import "github.com/jinzhu/gorm"

type Alarm struct{
	AlarmTime 			int64 	//响铃时间
	Location  			string	//闹铃地点
	Remark	  			string	//闹铃备注
	PreSetAlarmAudio 	string  //系统预制闹铃
	UserSetAlarmAudio	string  //用户上传的闹铃
	RepeatType 			int 	//重复类型
}

type AlarmModel struct{
	gorm.Model
	UserId 			string //闹钟属于的用户
	CreatedUserId 	string //创建这条闹钟的用户
	AlarmInfo		Alarm  //闹钟的信息
}