/**
 * Created by Elvizlai on 2015/10/31 22:26
 * Copyright © PubCloud
 */

package enum

var RespCode *respCode

type Code struct {
	key   int
	value string
}

type respCode struct {
	OK                *Code
	UserNotExist      *Code
	PasswordIncorrect *Code
	EmailExist        *Code
	NickNameExist     *Code

	BadRequest        *Code
	UnAuthorized      *Code
	NotFound          *Code
}

func init() {
	RespCode = new(respCode)
	RespCode.OK = &Code{0, "OK"}
	RespCode.UserNotExist = &Code{1, "用户不存在"}
	RespCode.PasswordIncorrect = &Code{2, "密码错误"}
	RespCode.EmailExist = &Code{3, "该邮箱已注册"}
	RespCode.NickNameExist = &Code{4, "该昵称已被使用"}

	//system error
	RespCode.BadRequest = &Code{400, "BadRequest"}
	RespCode.UnAuthorized = &Code{401, "UnAuthorized"}
	RespCode.NotFound = &Code{404, "NotFound"}
}

func (c Code)Str() string {
	return c.value
}

func (c Code)Code() int {
	return c.key
}


//type Code int
//
//const (
//	OK Code = iota
//)
//
////100~ 用户相关
//const (
//	EmailAlreadyExist Code = iota + 100
//	NickNameAlreadyExist
//	PasswordIncorrect
//	UserNotExist
//)
//
//const (
//	UNKNOWN Code = 999
//)
//
//func (c Code)Str() string {
//	switch c{
//	case OK:
//		return "OK"
//	case NickNameAlreadyExist:
//		return "用户名已存在"
//	case EmailAlreadyExist:
//		return "邮箱已存在"
//	case PasswordIncorrect:
//		return "密码错误"
//	case UserNotExist:
//		return "用户不存在"
//	case UnAuthorized:
//		return "未授权"
//	case BadRequest:
//		return "Bad Request"
//	case NotFound:
//		return "Not Found"
//	default:
//		return "UNKNOWN:" + fmt.Sprint(c)
//	}
//}
//
//func (c Code)Code() int {
//	return int(c)
//}