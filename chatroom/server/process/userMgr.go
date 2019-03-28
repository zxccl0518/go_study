package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr的初始化。
func init() {
	userMgr = &UserMgr{
		make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加。
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserID] = up
}

// 完成对onlineUsers的删除
func (this *UserMgr) deleteOnlineUser(up *UserProcess) {
	delete(this.onlineUsers, up.UserID)
}

// 返回当前所有在线用户。
func (this *UserMgr) getAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// 根绝id返回对应的值
func (this *UserMgr) getOnlineUserByID(userID int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userID]
	if !ok { //说明你要查询的那个用户 当前不在线
		err = fmt.Errorf("用户 %d 不存在", userID)
		return
	}

	return
}
