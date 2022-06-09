package model

import (
	"encoding/json"
	"fmt"
	"go_code/project01/main/chapter19/common/message"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const (
	_usersRedisKey                = "users"         //用户库
	_offlineUserMesRedisKeyPrefix = "off_mes_user_" //离线用户的消息列表
)

//我们在服务器启动后，就初始化一个userDao实例
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//UserDao 定义一个UserDao 结构体
//完成对User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//NewUserDao 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//1. 根据用户ID 返回一个User实例+err
func (dao *UserDao) getUserByID(conn redis.Conn, id int) (user *message.User, err error) {
	//通过给定ID 去 redis 中查询用户
	res, err := redis.String(conn.Do("HGet", _usersRedisKey, id))
	if err != nil {
		//错误!
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应的ID
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	//这里我们需要把res 反序列化成User结构体
	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

//Login 2. 完成登录校验
//如果用户的ID和密码都正确，则返回一个user实例
//如果用户的ID或者密码有错误，则返回对应的的错误信息
func (dao *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	//先从userDao 的 连接池中取出一个连接
	conn := dao.pool.Get()
	defer conn.Close()

	user, err = dao.getUserByID(conn, userID)
	if err != nil {
		return
	}

	//这时至少证明用户是存在的，需要继续判断密码等信息
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//Register 3.完成注册逻辑
func (dao *UserDao) Register(user *message.User) (err error) {
	//先从userDao 的 连接池中取出一个连接
	conn := dao.pool.Get()
	defer conn.Close()

	_, err = dao.getUserByID(conn, user.UserID)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//这时，说明id在redis还没有，则可以完成注册
	//这里我们需要把user序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//入库
	_, err = conn.Do("HSet", _usersRedisKey, user.UserID, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}

	return
}

//SaveUserOfflineMessage 4. 保存离线消息
func (dao *UserDao) SaveUserOfflineMessage(userID int, mes *message.OfflineSmsMes) (err error) {
	conn := dao.pool.Get()
	defer conn.Close()

	//mes序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//存入list中
	_, err = conn.Do("RPush", _offlineUserMesRedisKeyPrefix+strconv.Itoa(userID), string(data))
	if err != nil {
		fmt.Println("保存离线消息错误 err=", err)
		return
	}
	return
}

//deleteUserAllOfflineMessage 删除用户所有的离线消息
func (dao *UserDao) deleteUserAllOfflineMessage(conn redis.Conn, userID int) (err error) {
	//直接删除掉list
	_, err = conn.Do("Del", _offlineUserMesRedisKeyPrefix+strconv.Itoa(userID))
	if err != nil {
		fmt.Println("删除用户离线消息列表错误 err=", err)
		return
	}
	return
}

//GetUserAllOfflineMessage 获取用户所有的离线消息
func (dao *UserDao) GetUserAllOfflineMessage(userID int) ([]message.OfflineSmsMes, error) {
	conn := dao.pool.Get()
	defer conn.Close()

	messages, err := redis.Strings(conn.Do("LRange", _offlineUserMesRedisKeyPrefix+strconv.Itoa(userID), 0, -1))
	if err != nil {
		fmt.Println("获取离线消息列表错误，err=", err)
		return nil, err
	}

	var list []message.OfflineSmsMes
	//此处循环解析成OfflineSmsMes对象
	for _, s := range messages {
		var sms message.OfflineSmsMes
		err = json.Unmarshal([]byte(s), &sms)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			continue
		}
		list = append(list, sms)
	}

	//获取全部消息后需要删除一下
	dao.deleteUserAllOfflineMessage(conn, userID)

	return list, nil
}

//ExistsUserByID 7. 查看用户是否存在
func (dao *UserDao) ExistsUserByID(userID int) (b bool, err error) {
	conn := dao.pool.Get()
	defer conn.Close()

	//通过给定ID 去 redis 中查询用户
	b, err = redis.Bool(conn.Do("HExists", _usersRedisKey, userID))
	if err != nil {
		return
	}

	return
}

//GetAllUsers 8. 获取所有的在线用户
func (dao *UserDao) GetAllUsers() (users []*message.User, err error) {
	conn := dao.pool.Get()
	defer conn.Close()

	us, err := redis.Strings(conn.Do("HGetAll", _usersRedisKey))
	if err != nil {
		fmt.Println("HGetAll err=", err)
		return nil, err
	}

	for i := 1; i < len(us); i = i + 2 {
		//这里我们需要把res 反序列化成User结构体
		user := &message.User{}
		err = json.Unmarshal([]byte(us[i]), user)
		if err != nil {
			fmt.Println("json.Unmarshal err=", err)
			continue
		}

		users = append(users, user)
	}

	return
}
