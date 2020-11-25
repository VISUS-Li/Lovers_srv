package DB

import (
	"github.com/jinzhu/gorm"
)

/******
	跑步数据统计
******/
type RunStatistics struct{
	gorm.Model
	UserId        string `gorm:"unique;not null" json:"user_id"` //用户ID
	TotalDis      int32  `json:"total_dis"`                      //总共跑步长度（米）
	TotalTime     int32  `json:"total_time"`                     //总耗时长(秒)
	TotalRunTime  int32  `json:"total_run_time"`				 //总跑步耗时，不包含暂停等未跑步耗时
	TotalCalorie  int32  `json:"total_calorie"`                  //总消耗卡路里
	TotalDays     int32  `json:"total_days"`                     //总跑步天数
	LatestTime    int64  `json:"latest_time"`                    //最近一次运动时间
	Farthest      int32  `json:"farthest"`                       //最远距离
	LongestTime   int32  `json:"longest_time"`                   //最长距离
	FastestPace   int32  `json:"fastest_pace"`                   //最快配速
	LastRunIndex  int32  `json:"last_run_index"`				 //最后一次跑步的序号，也是总跑步次数
	LastRunId	  string `json:"last_run_id"`					 //最后一次跑步的ID
}

type RunItemData struct{
	gorm.Model
	UserId			string `gorm:"unique;not null" json:"user_id"` 		//用户ID
	RunId			string `gorm:"unique;not null" json:"run_id"` 		//当前跑步项的ID
	RunIndex		int32  `gorm:"unique;not null" json:"run_indxe"` 	//当前跑步项的序号
	FastestPace		int32  `json:"fastest_pace"`						//当前跑步最快配速
	SlowestPace		int32  `json:"slowest_pace"`						//当前跑步最慢配速
	RunDistance		int32  `json:"run_distance"`						//当前跑步距离
	RunCalorie		int32  `json:"run_calorie"`							//当前跑步消耗卡路里
	RunTotalTime	int64  `json:"run_total_time"`						//本次跑步总耗时
	StartTime 		int64  `json:"start_time"`							//开始跑步时间
	EndTime			int64  `json:"end_time"`							//结束跑步时间
	PauseTime		int32  `json:"pause_time"`							//暂停耗时
}