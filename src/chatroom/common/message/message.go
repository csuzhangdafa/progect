package message

const (
	LoginMestype    = "LoginMes"
	LoginResMestype = "LoginResMes"
	RegisterMestype = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 定义两个消息，两个返回类型
type LoginMes struct {
	Userid   int    `json:"userid"`  //用户ID
	Userpwd  string `json:"userpwd"` //用户密码
	Username string `json:"username"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码
	Error string `json:"error"` //返回错误信息
}

type RegisterMes struct {
}
