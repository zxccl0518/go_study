package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 我们在服务器启动后，就初始化一个userDao的实例，
// 把它做成全局的，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

// 定义一个结构体，完成对User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao的实例。
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 思考一下，UserDao 应该提供哪些方法
// 1.根据用户id 返回一个User实例+err
func (this *UserDao) getUserByID(conn redis.Conn, userID int) (user *User, err error) {
	// 通过给定的id 去redis查询用户，
	res, err := redis.String(conn.Do("HGET", "users", userID))
	if err != nil {
		// 表示在user哈希中，没有找到对应的id
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}

	fmt.Println("UserDao.go === > res = ", res)
	user = &User{}
	// 我们需要把 res反序列化成user实例。
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("user 反序列化失败。")
	}

	return
}

// 完成登录的一个校验
// 1.Login() 完成对用户的验证，
// 2.如果用户的id 密码都正确，返回一个user实例。
// 3.如果用户的id 或者密码不正确。返回一个错误的信息。
func (this *UserDao) Login(userID int, userPwd string) (user *User, err error) {
	// 先从 userDao连接池中，取出一个连接。
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserByID(conn, userID)
	if err != nil {
		return
	}

	// 到了这步，证明用户是获取到了，但是不知道密码是否正确。
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}
