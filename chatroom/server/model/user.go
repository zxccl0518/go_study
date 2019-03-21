package model

// 定义一个用户的结构体。
type User struct {
	// 为了序列化和反序列化 成功我们必选保证
	// 用户信息的json字符串的key 和 结构体的字段对应的tag 名字一致
	UserID   int    `json:"userID"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
