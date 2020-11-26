////运动常量定义文件
package SportConst

//运动类型定义
const (
	SPORT_TYPE_UNKNOWN = 0
	SPORT_TYPE_RUN = 1
	SPORT_TYPE_CYCLE = 2
	SPORT_TYPE_OTHER = 3
)

//运动状态定义
const (
	SPORT_STATUS_UNKNOWN      = 0
	SPROT_STATUS_ING          = 1 //正在运动
	SPORT_STATUS_PAUSE        = 2
	SPORT_STATUS_STOP         = 3
	SPORT_STATUS_ABNORMALEXIT = 4 //异常退出
)
const(
	SPORT_STATUS_SIZE			= 4 //运动状态数据大小
)

const(
	//跑步收到的最大数量的来自客户端的推送消息后，进行自动保存
	MAX_RECIEVE_RUN_SAVE_COUNT	= 100
)
