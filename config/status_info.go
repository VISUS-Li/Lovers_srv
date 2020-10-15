package config

/******
状态码定义
******/
//通用状态定义
const(
	CODE_ERR_SUCCESS        = 1000
	CODE_ERR_UNKNOW         = 50000
	INVALID_PARAMS 			= 400
)

////用户相关状态定义

//用户验证相关
const(
	CODE_ERR_AUTH_CHECK_TOKEN_FAIL    = 1001
	CODE_ERR_AUTH_CHECK_TOKEN_TIMEOUT = 1002
	CODE_ERR_AUTH_TOKEN               = 1003
	CODE_ERR_AUTH                     = 1004
	CODE_ERR_AUTH_TOKEN_EMPTY		  = 1017
)

const(
	CODE_ERR_PARAM_EMPTY			 = 1005
	CODE_ERR_SERVER_INTERNAL		 = 1011 //服务器内部错误
	CODE_ERR_INSERT_DB_FAIL		  	 = 1013	//插入数据库失败
	CODE_ERR_SELECT_DB_FAIL		  	 = 1014	//查询数据库失败

)
//用户登录相关
const(
	CODE_ERR_LOGIN_QUERY       = 1006
	CODE_ERR_LOGIN_NO_USER     = 1007
	CODE_ERR_LOGIN_PWD_ERROR   = 1008
	CODE_ERR_LOGIN_TOKEN_ERROR = 1009
	CODE_ERR_LOGIN_EXPIRE	   = 1016 	//用戶token过期，需重新登录
)

//用户注册相关
const(
	CODE_ERR_REG_PHONE_ERR		= 1010
	CODE_ERR_REG_PHONE_EXIST	= 1018
)

//用户操作相关
const(
	CODE_ERR_USER_NOT_EXIST = 1015
)

//主页相关
const(
	CODE_ERR_HOME_NOT_ENOUGH_CARD = 1012
)

/******
状态信息定义（信息文字内容）
******/

//通用
const(

	MSG_REQUEST_SUCCESS	 	= "请求成功"
	MSG_SERVER_INTERNAL 	= "服务器内部错误"
	MSG_ERR_INSERT_DB_FAIL 	= "插入数据库失败"
	MSG_ERR_SELECT_DB_FAIL 	= "查询数据库失败"
	MSG_ERR_PARAM_WRONG		= "传入参数错误"
)

//用户相关
const(
	MSG_DB_LOGIN_OK              = "登录成功"
	MSG_DB_LOGIN_IN_EMPTY        = "用户名或密码为空"
	MSG_DB_LOGIN_QUERY_ERR       = "查询数据库失败"
	MSG_DB_LOGIN_NO_USER         = "用户未注册"
	MSG_DB_LOGIN_PWD_ERROR       = "密码错误"
	MSG_DB_LOGIN_TOKEN_ERROR     = "获取token失败"
	MSG_AUTH_TOKEN_EXPIRE		 = "Token过期"
	MSG_AUTH_TOKEN_EMPTY		 = "Token为空"
	MSG_AUTH_TOKEN_ERROR		 = "Token解析错误"
)

//注册相关
const(
	MSG_DB_REG_OK        = "注册成功"
	MSG_DB_REG_IN_EMPTY  = "传入用户名密码为空"
	MSG_DB_REG_EXIST     = "该账号已存在"
	MSG_DB_REG_REG_ERR   = "注册失败，可能插入数据库失败"
	MSG_DB_REG_PHONE_ERR = "手机格式不正确"
	MSG_DB_REG_PARAM_nil = "传入参数为空"
)

//用户操作相关
const(
	MSG_USER_NOT_EXIST = "用户不存在"
)

//主页相关
const(
	MSG_HOME_NOT_ENOUGH_CARD = "没有足够的主页卡片"
)

const(
	NOTELISTUP_INVALID_PARAM = "事件清单参数错误"

)