package types

import (
	"encoding/json"
	"fmt"
)

type AuthType string

const (
	AuthType_Null       AuthType = ""
	AuthType_Basic      AuthType = "Basic"
	AuthType_Bearer     AuthType = "Bearer"
	AuthType_ProjectKey AuthType = "ProjectKey"
)

func (t AuthType) String() string {
	return string(t)
}
func (t AuthType) Valid() bool {
	switch t {
	case AuthType_Basic:
		return true
	case AuthType_Bearer:
		return true
	case AuthType_ProjectKey:
		return true
	}
	return false
}

const (
	HEADER_AUTHORIZATION = "Authorization"
	HEADER_AUTH_TOKEN    = "Auth-Token"
)

const (
	MAX_DURATION int32  = 1440
	MIN_DURATION int32  = 1
	JSON_RPC_VER string = "2.0"
)

type BizCode int

const (
	CODE_ERROR                           BizCode = -1   //unknown error
	CODE_OK                              BizCode = 0    //success
	CODE_TOO_MANAY_REQUESTS              BizCode = 429  //too many requests
	CODE_INTERNAL_SERVER_ERROR           BizCode = 500  //internal service error
	CODE_DATABASE_ERROR                  BizCode = 501  //database server error
	CODE_ACCESS_DENY                     BizCode = 1000 //access deny
	CODE_UNAUTHORIZED                    BizCode = 1001 //user unauthorized
	CODE_INVALID_USER_OR_PASSWORD        BizCode = 1002 //user or password incorrect
	CODE_INVALID_PARAMS                  BizCode = 1003 //parameters invalid
	CODE_INVALID_JSON_OR_REQUIRED_PARAMS BizCode = 1004 //json format is invalid
	CODE_ALREADY_EXIST                   BizCode = 1005 //account name already exist
	CODE_NOT_FOUND                       BizCode = 1006 //record not found
	CODE_INVALID_PASSWORD                BizCode = 1007 //wrong password
	CODE_INVALID_AUTH_CODE               BizCode = 1008 //invalid auth code
	CODE_ACCESS_VIOLATE                  BizCode = 1009 //access violate
	CODE_TYPE_UNDEFINED                  BizCode = 1010 //type undefined
	CODE_BAD_DID_OR_SIGNATURE            BizCode = 1011 //bad did or signature
	CODE_ACCOUNT_BANNED                  BizCode = 1012 //account was banned
)

var codeMessages = map[BizCode]string{
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

func (c BizCode) Ok() bool {
	return c == CODE_OK
}

func (c BizCode) String() string {
	if m, ok := codeMessages[c]; ok {
		return m
	}
	return fmt.Sprintf("CODE_UNKNOWN<%d>", c)
}

func (c BizCode) GoString() string {
	return c.String()
}

type HttpHeader struct {
	Code    BizCode `json:"code"`    //response code of business (0=OK, other fail)
	Message string  `json:"message"` //error message
	Total   int64   `json:"total"`   //result total
	Count   int     `json:"count"`   //result count (single page)
}

type HttpResponse struct {
	Header HttpHeader  `json:"header"` //response header
	Data   interface{} `json:"data"`   //response data body
}

type RpcRequest struct {
	Id      interface{} `json:"id"`       //0
	JsonRpc string      `json:"json_rpc"` //2.0
	Method  string      `json:"method"`   //JSON-RPC method
	//Params  []interface{} `json:"params"`   //JSON-RPC parameters [any...]
}

type RpcError struct {
	Code    BizCode     `json:"code"`    //response code of business (0=OK, other fail)
	Message string      `json:"message"` //error message
	Data    interface{} `json:"data"`    //error attach data
}

type RpcResponse struct {
	Id      interface{} `json:"id"`       //0
	JsonRpc string      `json:"json_rpc"` //2.0
	Error   RpcError    `json:"error"`    //error message
	Result  interface{} `json:"result"`   //JSON-RPC result
}

func (r *RpcResponse) String() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func NewRpcResponse(id interface{}, result interface{}) *RpcResponse {
	return &RpcResponse{
		Id:      id,
		JsonRpc: JSON_RPC_VER,
		Result:  result,
	}
}

func NewRpcError(id interface{}, code BizCode, strError string) *RpcResponse {
	return &RpcResponse{
		Id:      id,
		JsonRpc: JSON_RPC_VER,
		Error: RpcError{
			Code:    code,
			Message: strError,
			Data:    nil,
		},
	}
}
