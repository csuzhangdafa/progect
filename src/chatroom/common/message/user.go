package message

//定义一个用户的结构体

type User struct {

	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证用户信息的json字符串key和结构体对应的tag名字一致
	Userid     int    `json:"userid"`
	Userpwd    string `json:"userpwd"`
	Username   string `json:"username"`
	Userstatus int    `json:"userstatus"`
}
