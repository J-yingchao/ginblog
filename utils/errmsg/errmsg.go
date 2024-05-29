package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... user model err
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	//code = 2000... article model err

	//code = 3000... category model err
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "Username already exists!",
	ERROR_PASSWORD_WRONG:   "Incorrect password",
	ERROR_USER_NOT_EXIST:   "User does not exist",
	ERROR_TOKEN_EXIST:      "TOKEN does not exist, please log in again",
	ERROR_TOKEN_RUNTIME:    "TOKEN has expired, please log in again",
	ERROR_TOKEN_WRONG:      "Incorrect TOKEN, please log in again",
	ERROR_TOKEN_TYPE_WRONG: "Incorrect TOKEN format, please log in again",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
