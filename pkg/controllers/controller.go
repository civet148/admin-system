package controllers

import (
	"admin-system/pkg/config"
	"admin-system/pkg/dal"
	"admin-system/pkg/middleware"
	"admin-system/pkg/privilege"
	"admin-system/pkg/sessions"
	"admin-system/pkg/types"
	"encoding/json"
	"fmt"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	*dal.Dal
	cfg *config.Config
}

func NewController(cfg *config.Config) *Controller {
	return &Controller{
		cfg: cfg,
		Dal: dal.GetDal(cfg),
	}

}

func (m *Controller) OK(c *gin.Context, data interface{}, count int, total int64) {
	var bc = types.BizOK
	if data == nil {
		data = struct{}{}
	}
	var r = &types.HttpResponse{
		Header: types.HttpHeader{
			Code:    bc.Code,
			Message: bc.Message,
			Count:   count,
			Total:   total,
		},
		Data: data,
	}
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func (m *Controller) Error(c *gin.Context, bc types.BizCode) {
	var r = &types.HttpResponse{
		Header: types.HttpHeader{
			Code:    bc.Code,
			Message: bc.Message,
			Count:   0,
		},
		Data: struct{}{},
	}
	log.Errorf("[Controller] response error code [%d] message [%s]", bc.Code, bc.Message)
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func (m *Controller) ErrorStatus(c *gin.Context, status int, message string) {
	log.Errorf("[Controller] http status code [%d] message [%s]", status, message)
	c.String(status, message)
	c.Abort()
}

func (m *Controller) RpcResult(c *gin.Context, data interface{}, err error, id interface{}) {
	var status = http.StatusOK
	var strResp string
	if err != nil {
		status = http.StatusInternalServerError
		data = &types.RpcResponse{
			Id:      id,
			JsonRpc: "2.0",
			Error: types.RpcError{
				Code:    types.CODE_INTERNAL_SERVER_ERROR,
				Message: err.Error(),
			},
			Result: nil,
		}
	}
	switch data.(type) {
	case string:
		strResp = data.(string)
	default:
		{
			b, _ := json.Marshal(data)
			strResp = string(b)
		}
	}

	c.String(status, strResp)
	c.Abort()
}

func (m *Controller) GetClientIP(c *gin.Context) (strIP string) {
	return c.ClientIP()
}

func (m *Controller) ContextPlatformPrivilege(c *gin.Context, privileges ...privilege.Privilege) (ctx *types.Context, ok bool) {
	var err error
	ctx = sessions.GetContext(c)
	if ctx == nil {
		err = log.Errorf("user session context is nil, token [%s]", middleware.GetAuthToken(c))
		m.Error(c, types.NewBizCode(types.CODE_UNAUTHORIZED, err.Error()))
		return
	}
	for _, auth := range privileges {
		if m.PlatformCore.CheckPrivilege(c, ctx, ctx.UserName(), auth) {
			return ctx, true
		}
	}
	err = log.Errorf("operator name [%s] id [%v] have no privilege %+v", ctx.UserName(), ctx.UserId(), privileges)
	m.Error(c, types.NewBizCode(types.CODE_ACCESS_DENY, err.Error()))
	return ctx, false
}

func (m *Controller) bindJSON(c *gin.Context, req interface{}) (err error) {
	if err = c.ShouldBindJSON(req); err != nil {
		log.Errorf("%s", err)
		m.Error(c, types.NewBizCode(types.CODE_INVALID_JSON_OR_REQUIRED_PARAMS, err.Error()))
		c.Abort()
		return
	}
	body, _ := json.MarshalIndent(req, "", "\t")
	log.Debugf("request from [%s] body [%+v]", c.ClientIP(), string(body))
	return
}

func (m *Controller) isNilString(strIn string) bool {
	if strIn == "" {
		return true
	}
	return false
}

func (m *Controller) isZero(n interface{}) bool {
	strNumber := fmt.Sprintf("%v", n)
	if strNumber == "0" {
		return true
	}
	return false
}
