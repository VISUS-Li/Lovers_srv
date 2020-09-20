package config

//通用状态定义
const(
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
)

//用户相关
const(
	DB_LOGIN_OK              = "登录成功"
	DB_LOGIN_IN_EMPTY        = "用户名或密码为空"
	DB_LOGIN_NO_UNIQUE_IN_DB = "内部错误，用户在数据库中不唯一"
	DB_LOGIN_NO_USER         = "用户未注册"
	DB_LOGIN_PWD_ERROR       = "密码错误"
)

const(
	DB_REG_OK				 = "注册成功"
	DB_REG_IN_EMPTY			 = "传入用户名密码为空"
	DB_REG_EXIST			 = "该账号已存在"
	DB_REG_REG_ERR			 = "注册失败，可能插入数据库失败"
	DB_REG_PHONE_ERR		 = "手机格式不正确"
	DB_REG_PARAM_nil		 = "传入参数为空"
)

const(
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 1001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 1002
	ERROR_AUTH_TOKEN               = 1003
	ERROR_AUTH                     = 1004
)

const(
	NOTELISTUP_INVALID_PARAM = "事件清单参数错误"

)