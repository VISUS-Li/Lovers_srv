package config

/******
枚举定义
******/

////用户相关

//登录相关
const(
	ENUM_LOGIN_USERNAME 	 = 1
	ENUM_LOGIN_WEIXIN 		 = 2
	ENUM_LOGIN_QQ			 = 3
	ENUM_LOGIN_VERCODE	     = 4 //验证码登录
)


/******
	错误码定义
******/

//通用错误码定义
const(
	ENUM_ERR_OK			   = 0
	ENUM_ERR_INVALID_PARAM = 1

)

//redis缓存错误码定义
const(
	ENUM_ERR_SETHASH_FAILED = 10
	ENUM_ERR_GETHASH_FAILED = 11
	ENUM_ERR_DELHASH_FAILED = 12
)

//数据库读写错误码定义
const(
	ENUM_ERR_DB_QUERY_FAILED  		= 20
	ENUM_ERR_DB_QUERY_NOT_FOUND		= 21
	ENUM_ERR_DB_INSERT_FAILED 		= 22
	ENUM_ERR_DB_INSERT_DUPLICATE 	= 23 //插入数据库时，主键重复
	ENUM_ERR_DB_UPDATE_FAILED 		= 24
	ENUM_ERR_DB_DELETE_FAILED		= 25
)