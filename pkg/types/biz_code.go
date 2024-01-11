package types

import "fmt"

type BizCode struct {
	Code    int
	Message string
}

const (
	CODE_ERROR                           = -1   //unknown error
	CODE_OK                              = 0    //success
	CODE_TOO_MANAY_REQUESTS              = 429  //too many requests
	CODE_INTERNAL_SERVER_ERROR           = 500  //internal service error
	CODE_DATABASE_ERROR                  = 501  //database server error
	CODE_ACCESS_DENY                     = 1000 //access deny
	CODE_UNAUTHORIZED                    = 1001 //user unauthorized
	CODE_INVALID_USER_OR_PASSWORD        = 1002 //user or password incorrect
	CODE_INVALID_PARAMS                  = 1003 //parameters invalid
	CODE_INVALID_JSON_OR_REQUIRED_PARAMS = 1004 //json format is invalid
	CODE_ALREADY_EXIST                   = 1005 //account name already exist
	CODE_NOT_FOUND                       = 1006 //record not found
	CODE_INVALID_PASSWORD                = 1007 //wrong password
	CODE_INVALID_AUTH_CODE               = 1008 //invalid auth code
	CODE_ACCESS_VIOLATE                  = 1009 //access violate
	CODE_TYPE_UNDEFINED                  = 1010 //type undefined
	CODE_BAD_DID_OR_SIGNATURE            = 1011 //bad did or signature
	CODE_ACCOUNT_BANNED                  = 1012 //account was banned
)

var codeMessages = map[int]string{
	CODE_ERROR:                           "unknown error",
	CODE_OK:                              "OK",
	CODE_TOO_MANAY_REQUESTS:              "too many requests",
	CODE_INTERNAL_SERVER_ERROR:           "internal server error",
	CODE_DATABASE_ERROR:                  "database error",
	CODE_UNAUTHORIZED:                    "unauthorized",
	CODE_ACCESS_DENY:                     "access deny",
	CODE_INVALID_USER_OR_PASSWORD:        "invalid user or password",
	CODE_INVALID_PARAMS:                  "invalid params",
	CODE_INVALID_JSON_OR_REQUIRED_PARAMS: "invalid json request",
	CODE_ALREADY_EXIST:                   "data already exist",
	CODE_NOT_FOUND:                       "data not found",
	CODE_INVALID_PASSWORD:                "invalid password",
	CODE_INVALID_AUTH_CODE:               "invalid auth code",
	CODE_ACCESS_VIOLATE:                  "access violate",
	CODE_TYPE_UNDEFINED:                  "type undefined",
	CODE_BAD_DID_OR_SIGNATURE:            "bad id or signature",
	CODE_ACCOUNT_BANNED:                  "account banned",
}

var BizOK = BizCode{
	Code: CODE_OK,
}

func (c BizCode) Ok() bool {
	return c.Code == CODE_OK
}

func (c BizCode) String() string {
	if c.Message != "" {
		return c.Message
	}
	if m, ok := codeMessages[c.Code]; ok {
		return m
	}
	return fmt.Sprintf("unknown_code<%d>", c.Code)
}

func (c BizCode) GoString() string {
	return c.String()
}

func NewBizCode(code int, messages ...string) BizCode {
	if code == CODE_OK {
		return BizCode{}
	}
	var msg string
	if len(messages) > 0 {
		msg = messages[0]
	} else {
		msg = codeMessages[code]
	}

	return BizCode{
		Code:    code,
		Message: msg,
	}
}

func NewBizCodeDatabaseError(msg string) BizCode {
	return NewBizCode(CODE_DATABASE_ERROR, msg)
}
