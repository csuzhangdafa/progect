package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

// 处理群发消息
func outputgroupmes(mes *message.Message) {
	//反序列化mes.data
	var smsmes message.Smsmes
	err := json.Unmarshal([]byte(mes.Data), &smsmes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示
	info := fmt.Sprintf("用户id:\t%d对大家说\t%s", smsmes.Userid, smsmes.Content)
	fmt.Println(info)

}
