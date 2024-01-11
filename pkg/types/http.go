package types

import (
	"encoding/json"
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

type HttpHeader struct {
	Code    int    `json:"code"`    //response code of business (0=OK, other fail)
	Message string `json:"message"` //error message
	Total   int64  `json:"total"`   //result total
	Count   int    `json:"count"`   //result count (single page)
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
	Code    int         `json:"code"`    //response code of business (0=OK, other fail)
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

func NewRpcError(id interface{}, code int, strError string) *RpcResponse {
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
