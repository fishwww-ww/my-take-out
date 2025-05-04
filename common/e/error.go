package e

import "errors"

var (
	Error_PASSWORD_ERROR    = errors.New("密码错误")
	Error_ACCOUNT_NOT_FOUND = errors.New("账号不存在")
	Error_ACCOUNT_LOCKED    = errors.New("账号被锁定")
)
