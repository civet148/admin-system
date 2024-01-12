package controllers

import (
	"admin-system/pkg/dal/models"
	"admin-system/pkg/middleware"
	"admin-system/pkg/privilege"
	"admin-system/pkg/proto"
	"admin-system/pkg/sessions"
	"admin-system/pkg/types"
	"admin-system/pkg/utils"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
)

func (m *Controller) PlatformLogin(c *gin.Context) { //user login
	var err error
	var req proto.PlatformLoginReq

	var ctx *types.Context
	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf("%s", err)
		return
	}

	var do *models.UserDO
	var strIP = m.GetClientIP(c)
	var code types.BizCode
	if do, code = m.PlatformCore.UserLogin(req.UserName, req.Password, strIP); !code.Ok() {
		m.Error(c, code)
		return
	}
	s := &types.Session{
		UserId:      do.GetId(),
		UserName:    do.GetUserName(),
		Alias:       do.GetUserAlias(),
		PhoneNumber: do.GetPhoneNumber(),
		IsAdmin:     do.GetIsAdmin(),
		LoginIP:     strIP,
	}

	if s.AuthToken, err = middleware.GenerateToken(s); err != nil {
		err = log.Errorf("generate token error [%s]", err.Error())
		m.Error(c, types.NewBizCode(types.CODE_INVALID_PARAMS, err.Error()))
		return
	}

	ctx = sessions.NewContext(s)
	log.Debugf("user [%v] login successful, user id [%v] is admin [%v] token [%s]", s.UserName, s.UserId, s.IsAdmin, s.AuthToken)

	role := m.PlatformCore.GetUserRole(ctx, do.GetUserName())
	if role == nil {
		err = log.Errorf("user [%s] role not found", req.UserName)
		m.Error(c, types.NewBizCode(types.CODE_NOT_FOUND, err.Error()))
		return
	}

	auth := privilege.GetUserRoleList(do.UserName)
	var resp = proto.PlatformLoginResp{
		Version:   m.cfg.Version,
		UserName:  do.UserName,
		AuthToken: s.AuthToken,
		LoginTime: do.LoginTime,
		LoginIp:   do.LoginIp,
		Role:      role.RoleName,
		Privilege: auth,
	}
	m.OK(c, &resp, 1, 1)
}

func (m *Controller) PlatformLogout(c *gin.Context) { //user logout
	sessions.RemoveContext(c)
	m.OK(c, nil, 0, 0)
}

func (m *Controller) PlatformCheckExist(c *gin.Context) { //check user account or email exist
	var err error
	var req proto.PlatformCheckExistReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	code := m.PlatformCore.CheckExist(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}
	m.OK(c, &proto.PlatformCheckExistResp{}, 1, 1)
}

func (m *Controller) PlatformListUser(c *gin.Context) { //list platform users

	var err error
	var req proto.PlatformListUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserAccess)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	users, total, code := m.PlatformCore.ListUser(ctx, &req)
	if !code.Ok() {
		log.Errorf("list user code [%s]", code.String())
		m.Error(c, code)
		return
	}
	m.OK(c, proto.PlatformListUserResp{Users: users}, len(users), total)
}

func (m *Controller) PlatformCreateUser(c *gin.Context) { //create user account

	var err error
	var req proto.PlatformCreateUserReq
	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserAdd)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.CheckUserNameExist(ctx, req.UserName); !code.Ok() {
		m.Error(c, code)
		return
	}

	if req.Email != "" {
		if code := m.PlatformCore.CheckUserEmailExist(ctx, req.Email); !code.Ok() {
			m.Error(c, code)
			return
		}
		if !utils.VerifyEmailFormat(req.Email) {
			err = log.Errorf("email [%s] format error", req.Email)
			m.Error(c, types.NewBizCode(types.CODE_INVALID_PARAMS, err.Error()))
			return
		}
	}
	user, code := m.PlatformCore.CreateUser(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}

	var resp = proto.PlatformCreateUserResp{
		UserId: user.GetId(),
	}

	m.OK(c, &resp, 1, 1)
}

func (m *Controller) PlatformEditUser(c *gin.Context) { //edit user information
	var err error
	var req proto.PlatformEditUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserEdit)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	var code types.BizCode
	if code = m.PlatformCore.EditUser(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}
	m.OK(c, &proto.PlatformEditUserResp{}, 1, 1)
}

func (m *Controller) PlatformEnableUser(c *gin.Context) {
	var err error
	var req proto.PlatformEnableUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserAdd)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}
	r, code := m.PlatformCore.EnableUser(ctx, &req)
	if !code.Ok() {
		log.Warnf("name [%s] id [%v]  operator user failed", ctx.UserName(), ctx.UserId())
		m.Error(c, code)
		return
	}

	m.OK(c, r, 1, 1)
}

func (m *Controller) PlatformDisableUser(c *gin.Context) {
	var err error
	var req proto.PlatformDisableUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserAdd)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}
	r, code := m.PlatformCore.DisableUser(ctx, &req)
	if !code.Ok() {
		log.Warnf("name [%s] id [%v]  operator user failed", ctx.UserName(), ctx.UserId())
		m.Error(c, code)
		return
	}

	m.OK(c, r, 1, 1)
}

func (m *Controller) PlatformDeleteUser(c *gin.Context) { //delete user account

	var err error
	var req proto.PlatformDeleteUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.UserDelete)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}
	if req.UserName == ctx.UserName() {
		err = log.Errorf("can't delete self")
		m.Error(c, types.NewBizCode(types.CODE_ACCESS_DENY, err.Error()))
		return
	}

	if code := m.PlatformCore.DeleteUser(ctx, &req); !code.Ok() {
		log.Warnf("operator name [%s] id [%v]  delete user failed", ctx.UserName(), ctx.UserId())
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformDeleteUserResp{}, 1, 1)
}

func (m *Controller) PlatformListRole(c *gin.Context) { //list platform roles
	var err error
	var req proto.PlatformListRoleReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.RoleAccess)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	roles, total, code := m.PlatformCore.ListRole(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}
	count := len(roles)
	m.OK(c, &proto.PlatformListRoleResp{Roles: roles}, count, total)
}

func (m *Controller) PlatformCreateRole(c *gin.Context) { //create a custom platform role
	var err error
	var req proto.PlatformCreateRoleReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.RoleAdd)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.CreateRole(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformCreateRoleResp{}, 1, 1)
}

func (m *Controller) PlatformEditRole(c *gin.Context) { //edit custom platform role
	var err error
	var req proto.PlatformEditRoleReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.RoleEdit)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.EditRole(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformEditRoleResp{}, 1, 1)
}

func (m *Controller) PlatformDeleteRole(c *gin.Context) { //delete custom platform role
	var err error
	var req proto.PlatformDeleteRoleReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.RoleDelete)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.DeleteRole(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformDeleteRoleResp{}, 1, 1)
}

func (m *Controller) PlatformAuthRole(c *gin.Context) {
	var err error
	var req proto.PlatformAuthRoleReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.RoleAuthority)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.AuthRole(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformAuthRoleResp{}, 1, 1)
}

func (m *Controller) PlatformInquireAuth(c *gin.Context) {
	var err error
	var req proto.PlatformInquireAuthReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	authority, code := m.PlatformCore.InquireAuth(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformInquireAuthResp{Privilege: authority}, 1, 1)
}

func (m *Controller) PlatformPrivilegeLevel(c *gin.Context) {

	var err error
	var req proto.PlatformPrivilegeLevelReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	PlatformPrivilegeLevelResp, code := m.PlatformCore.PrivilegeLevel(ctx, &req)
	if !code.Ok() {
		log.Errorf("list device type code [%s]")
		m.Error(c, code)
		return
	}
	m.OK(c, PlatformPrivilegeLevelResp, 1, 1)
}

func (m *Controller) PlatformResetPassword(c *gin.Context) { //platform administrator reset other user's password
	var err error
	var req proto.PlatformResetPasswordReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if code := m.PlatformCore.ResetUserPassword(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}

	m.OK(c, &proto.PlatformResetPasswordResp{}, 1, 1)
}

func (m *Controller) PlatformChangePassword(c *gin.Context) { //platform user change password by self
	var err error
	var req proto.PlatformResetPasswordReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}
	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	if m.isNilString(req.OldPassword) {
		err = log.Errorf("request body [%+v] old password is nil or ", req)
		m.Error(c, types.NewBizCode(types.CODE_INVALID_PARAMS, err.Error()))
		return
	}

	ok, code := m.PlatformCore.CheckUserPassword(ctx, ctx.UserName(), req.OldPassword)
	if !ok {
		if code.Ok() {
			m.Error(c, code)
			return
		}
		log.Error("user [%s] old password [%s] not match when change password by self", ctx.UserName(), req.OldPassword)
		m.Error(c, types.NewBizCode(types.CODE_INVALID_PASSWORD))
		return
	}
	req.UserName = ctx.UserName() //user change password by self (so the user name must be self name)
	if code = m.PlatformCore.ResetUserPassword(ctx, &req); !code.Ok() {
		m.Error(c, code)
		return
	}
	m.OK(c, &proto.PlatformResetPasswordResp{}, 1, 1)
}

func (m *Controller) PlatformListRoleUser(c *gin.Context) { //list role user
	var err error
	var req proto.PlatformListRoleUserReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	users, total, code := m.PlatformCore.ListRoleUser(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}
	userCount := len(users)
	m.OK(c, &proto.PlatformListRoleUserResp{
		RoleName:  req.RoleName,
		UserCount: userCount,
		Users:     users,
	}, userCount, total)
}

func (m *Controller) PlatformRefreshAuthToken(c *gin.Context) {
	var err error
	var req proto.PlatformRefreshAuthTokenReq

	if err = c.BindJSON(&req); err != nil {
		log.Errorf("%s", err)
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	var IP = m.GetClientIP(c)

	s := &types.Session{
		UserId:      ctx.UserId(),
		UserName:    ctx.UserName(),
		Alias:       ctx.Alias(),
		PhoneNumber: ctx.PhoneNumber(),
		IsAdmin:     ctx.IsAdmin(),
		LoginIP:     IP,
	}

	if s.AuthToken, err = middleware.GenerateToken(s); err != nil {
		err = log.Errorf("generate token error [%s]", err.Error())
		m.Error(c, types.NewBizCode(types.CODE_ERROR, err.Error()))
		return
	}
	_ = sessions.NewContext(s)
	var resp = proto.PlatformRefreshAuthTokenResp{
		AuthToken: s.AuthToken,
	}
	m.OK(c, &resp, 1, 1)
}

func (m *Controller) PlatformListOperLog(c *gin.Context) {
	var err error
	var req proto.PlatformListOperLogReq

	if err = m.bindJSON(c, &req); err != nil {
		log.Errorf(err.Error())
		return
	}

	ctx, ok := m.ContextPlatformPrivilege(c, privilege.Null)
	if !ok {
		log.Errorf("user authentication context is nil or privilege check failed")
		return
	}

	list, total, code := m.PlatformCore.ListOperLog(ctx, &req)
	if !code.Ok() {
		m.Error(c, code)
		return
	}
	count := len(list)
	m.OK(c, &proto.PlatformListOperLogResp{
		List: list,
	}, count, total)
}
