package model

import (
	"errors"
)

//根据业务逻辑，自定义一些错误

var (
	ERROR_USER_NOTEEXISTS = errors.New("用户名不存在")
	ERROR_USER_EXISTS     = errors.New("用户名已经存在")
	ERROR_USER_PWD        = errors.New("密码不正确")
)
