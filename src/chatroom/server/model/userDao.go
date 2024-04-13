package model

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// 我们在服务器启动后，就初始化一个UserDao实例
// 把他做成全局变量，在需要和redis操作时，就直接使用
var (
	MyUserDao *UserDao
)

//定义一个Userdao结构体，完成对user结构体的各种操作

type UserDao struct {
	pool *redis.Client
}

// 使用工厂模式创建一个UserDao实例
/*
函数实现：在函数内部，首先创建了一个 UserDao 结构体的实例，
并将传入的 pool 参数赋值给该实例的 pool 字段。然后将该实例的指针返回给调用者。
工厂模式：工厂模式是一种创建型设计模式，它提供了一种将对象的实例化过程抽象出来的方法，
便在程序中通过调用工厂方法来创建对象，而不是直接使用 new 关键字或者手动创建对象。
这种方法可以隐藏对象的创建细节，使得代码更易于维护和扩展。
在这段代码中，通过使用工厂模式，您可以将 UserDao 的创建逻辑封装起来，
使得在其他地方需要创建 UserDao 实例时，只需要调用 NewUserDao 函数即可，
而不需要了解 UserDao 结构体的具体实现细节。这种做法有利于降低耦合性，提高代码的可维护性。
*/
func NewUserDao(pool *redis.Client) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 根据用户id返回一个User实例+err
func (this *UserDao) getUserByid(ctx context.Context, id int) (user *User, err error) {

	//通过给定的id去redis查找这个用户
	res, err := this.pool.HGet(ctx, "users", fmt.Sprintf("%d", id)).Result()
	if err != nil {
		//错误
		if err == redis.Nil {
			err = ERROR_USER_NOTEEXISTS
		}
		return
	}

	user = &User{}
	//把res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal ERR=", err)
		return
	}

	return
}

// 完成登录校验 login
// 如果用户id和pwd都正确返回一个user实例
// 如果id或者pwd有错误。返回错误信息
func (this *UserDao) Login(userid int, userpwd string) (user *User, err error) {

	ctx := context.Background()
	user, err = this.getUserByid(ctx, userid)
	if err != nil {
		return
	}

	//用户获取到了,验证密码
	if user.Userpwd != userpwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *User) (err error) {

	ctx := context.Background()
	_, err = this.getUserByid(ctx, user.Userid)
	if err == nil {
		//如果err = nil说明在数据库中已经有了这个用户，返回错误
		err = ERROR_USER_EXISTS
		return
	}

	//这时，说明id没在redis，可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//入库
	err = this.pool.HSet(ctx, "users", fmt.Sprintf("%d", user.Userid), data).Err()
	if err != nil {
		return
	}
	return
}
