package process

import (
	"fmt"
)

//因为UserMgr实例在服务器端有且只有一个，在很多地方都会使用到
//因此我们将它定义为全局变量

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*Userprocess
}

// 初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*Userprocess, 1024),
	}
}

// 完成添加
func (this *UserMgr) AddOnlineuser(up *Userprocess) {
	this.onlineUsers[up.Userid] = up
}

// 完成删除
func (this *UserMgr) DelOnlineUser(userid int) {
	delete(this.onlineUsers, userid)
}

// 查询返回当前所有的在线用户
func (this *UserMgr) Getallonlineuser() map[int]*Userprocess {
	return this.onlineUsers
}

// 根据id返回对应的值
func (this *UserMgr) GetOnlineuserByte(userid int) (up *Userprocess, err error) {

	//如何从map中取出一个值，带检测方式
	up, ok := this.onlineUsers[userid]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userid)
		return
	}
	return
}
