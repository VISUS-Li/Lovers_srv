package Cache

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"reflect"
	"time"
)

var (
	RedisPool *redis.Pool
	redisExpire int32
	DB *gorm.DB
)

func NewRedisPool(serverDB *gorm.DB) *redis.Pool {
	//初始化服务数据库
	DB = serverDB

	//计算Redis存储元素的超时时间
	d := time.Duration(config.GlobalConfig.Redis_ExpireTime)
	redisExpire = int32(d / time.Second)

	//构建新Redis连接池
	RedisPool = &redis.Pool{
		MaxIdle:     config.GlobalConfig.Redis_MaxIdle,
		MaxActive:   config.GlobalConfig.Redis_MaxActive,
		IdleTimeout: config.GlobalConfig.Redis_IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(config.GlobalConfig.Redis_NetWork, config.GlobalConfig.Redis_Addr,
				redis.DialConnectTimeout(config.GlobalConfig.Redis_DialTimeout),
				redis.DialReadTimeout(config.GlobalConfig.Redis_ReadTimeout),
				redis.DialWriteTimeout(config.GlobalConfig.Redis_WriteTimeout),
				redis.DialPassword(config.GlobalConfig.Redis_Pwd),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
	return RedisPool
}

/******
关闭Redis Pool
******/
func CloseRedis() error{
	return RedisPool.Close()
}

/******
ping Redis
*******/
func PingRedis() error{
	conn := RedisPool.Get()
	_, err := conn.Do("SET","PING","PONG")
	conn.Close()
	return err
}


/******
将整个结构结构化成json，再存成redis-hash
key:存redis时的键，用以确定模块，如登录可以用login作为key
fieldKey:用以确定用户，采用用户的userID作为存储
data: 存redis时的值
******/
func SetHashByJson(key string, fieldKey string, data interface{}) (error){

	//从Redis的连接库中获取一个连接，进行set操作
	conn := RedisPool.Get()
	defer conn.Close()

	//对传入的参数进行参数合法性判断
	if data == nil{
		return errors.New("the second param \"data\" is nil")
	}
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct{
		return errors.New("the second param \"data\" is not a struct")
	}

	//json化struct，前提是该struct中需要json化的field要有json的tag
	b, err := json.Marshal(data)
	if err != nil{
		return errors.New("json Marshal error:"+err.Error())
	}

	//进行HSET操作，将HSET写入缓存中，还没有正式执行该操作的
	if err = conn.Send("HSET", key, fieldKey, b ); err != nil{
		return Utils.ErrorOutputf("conn.Send(SET %s,%s) error(%v)", key, fieldKey, err)
	}

	//设置该Key的过期时间
	if err = conn.Send("EXPIRE", key, redisExpire); err != nil {
		return Utils.ErrorOutputf("conn.Send(EXPIRE %s) error(%v)", key, err)
	}

	//正式执行HSET和EXPIRE操作
	if err = conn.Flush(); err != nil{
		return Utils.ErrorOutputf("conn.Flush() error(%v)", err)
	}

	//接受执行HSET和EXPIRE的执行结果
	for i := 0; i < 2; i++ {
		if _, err = conn.Receive(); err != nil {
			return Utils.ErrorOutputf("conn.Receive() error(%v)", err)
		}
	}
	return nil
}

func GetHashByJson(key string, fieldKey string) (data []byte, err error){
	conn := RedisPool.Get()
	defer conn.Close()
	b, err := redis.Bytes(conn.Do("HGET", key,fieldKey))
	if err != nil{
		if err != redis.ErrNil {
			return nil, Utils.ErrorOutputf("conn.Do(HGET %s %s) error(%v)", key, fieldKey, err.Error())
		}
		return nil, nil
	}
	return b, nil
}

func DelHash(key string, fieldKey string) error{
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("HDEL", key, fieldKey)
	if err != nil{
		return Utils.ErrorOutputf("conn.Do(HDEL %s %s) errovr(%v)", key,fieldKey,err.Error())
	}
	return nil
}

func SetString(key string, val string)(error){
	conn := RedisPool.Get()
	defer conn.Close()

	if key == "" || len(key) <= 0 || val == "" || len(val) <= 0{
		return Utils.ErrorOutputf("param is nil")
	}
	//进行SET操作，将SET写入缓存中，还没有正式执行该操作的
	if err := conn.Send("SET", key, val); err != nil{
		return Utils.ErrorOutputf("conn.Send(SET %s,%s) error(%v)", key, val, err)
	}

	//设置该Key的过期时间
	if err := conn.Send("EXPIRE", key, redisExpire); err != nil {
		return Utils.ErrorOutputf("conn.Send(EXPIRE %s) error(%v)", key, err)
	}

	//正式执行HSET和EXPIRE操作
	if err := conn.Flush(); err != nil{
		return Utils.ErrorOutputf("conn.Flush() error(%v)", err)
	}

	//接受执行HSET和EXPIRE的执行结果
	for i := 0; i < 2; i++ {
		if _, err := conn.Receive(); err != nil {
			return Utils.ErrorOutputf("conn.Receive() error(%v)", err)
		}
	}
	return nil
}

func GetString(key string) (string, error){
	conn := RedisPool.Get()
	defer conn.Close()
	b, err := redis.String(conn.Do("GET", key))
	if err != nil{
		if err != redis.ErrNil {
			return "", Utils.ErrorOutputf("conn.Do(GET %s) error(%v)", key, err.Error())
		}
		return "", nil
	}
	return b, nil
}

func DelString(key string) error{
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil{
		return Utils.ErrorOutputf("conn.Do(DEL %s) errovr(%v)", key,err.Error())
	}
	return nil
}

func IsExistString(key string) (bool, error){
	conn := RedisPool.Get()
	defer conn.Close()
	b, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil{
		if err != redis.ErrNil {
			return false, Utils.ErrorOutputf("conn.Do(EXISTS %s) error(%v)", key, err.Error())
		}
		return false, nil
	}
	return b, nil
}