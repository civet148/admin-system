package core

import (
	"admin-system/pkg/config"
	"admin-system/pkg/dal/dao"
	"admin-system/pkg/dal/models"
	"admin-system/pkg/privilege"
	"admin-system/pkg/proto"
	"admin-system/pkg/types"
	"admin-system/pkg/utils"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const (
	InherentAdminName       = "admin"
	InherentAdminRole       = "admin"
	InherentAdminPassword   = "e10adc3949ba59abbe56e057f20f883e"
	InherentAdminNameRemark = ""
	InherentAdminRoleRemark = "supper administrator role"
)

var isInitialized bool

type PlatformCore struct {
	db            *sqlca.Engine
	cfg           *config.Config
	userDAO       *dao.UserDAO
	roleDAO       *dao.RoleDAO
	loginDAO      *dao.LoginDAO
	userRoleDAO   *dao.UserRoleDAO
	privilegeDAO  *dao.PrivilegeDAO
	dictionaryDAO *dao.DictionaryDAO
}

func NewPlatformCore(cfg *config.Config, db *sqlca.Engine) *PlatformCore {

	m := &PlatformCore{
		db:            db,
		cfg:           cfg,
		userDAO:       dao.NewUserDAO(db),
		roleDAO:       dao.NewRoleDAO(db),
		loginDAO:      dao.NewLoginDAO(db),
		userRoleDAO:   dao.NewUserRoleDAO(db),
		privilegeDAO:  dao.NewPrivilegeDAO(db),
		dictionaryDAO: dao.NewDictionaryDAO(db),
	}
	return m.initialize()
}

// initialize role and privilege of platform inherent
func (m *PlatformCore) initialize() *PlatformCore {
	if !isInitialized {
		privilege.InitCasbin(m.cfg.DSN)
		m.initializeInherentRoles()
		m.initializeInherentAccount()
	}
	isInitialized = true
	return m
}

func (m *PlatformCore) initializeInherentRoles() {
	var err error
	var exist bool

	if exist, err = m.roleDAO.CheckRoleExistByName(InherentAdminRole); err != nil {
		log.Errorf(err.Error())
		return
	}

	if exist {
		log.Debugf("role name [%s] already exist", "root")
		return
	}

	if _, err = m.roleDAO.Insert(&models.RoleDO{
		RoleName:   InherentAdminRole,
		CreateUser: InherentAdminName,
		EditUser:   InherentAdminName,
		Remark:     InherentAdminRoleRemark,
	}); err != nil {

		log.Errorf(err.Error())
		return
	}
	m.initializeInherentPrivileges()
}

func (m *PlatformCore) initializeInherentPrivileges() {
	log.Debugf("in")
	for _, authority := range privilege.TotalPrivilege {
		log.Debugf("init authority [%s]", authority)
		role := InherentAdminRole
		path := privilege.GetKeyMatchPath(authority)
		if strings.EqualFold(path, "") {
			log.Errorf("unknown authority : %s", authority)
			continue
		}
		// 给角色授予权限
		privilege.AddRoleAuthority(role, path, authority)
	}
}

func (m *PlatformCore) initializeInherentAccount() {
	if ok, err := m.userDAO.CheckActiveUserByUserName(InherentAdminName); err != nil {
		log.Error("query by user name error [%s]", err.Error())
		return
	} else if ok {
		return
	}

	privilege.AddUserRole(InherentAdminName, InherentAdminRole)

	var lastId int64
	var err error
	var strSalt = utils.GenerateSalt()
	do := &models.UserDO{
		UserName:   InherentAdminName,
		Password:   InherentAdminPassword,
		Salt:       strSalt,
		UserAlias:  InherentAdminName,
		IsAdmin:    false,
		Remark:     InherentAdminNameRemark,
		CreateUser: InherentAdminName,
		EditUser:   InherentAdminName,
		LoginTime:  utils.Now64(),
		Deleted:    false,
		State:      dao.UserState_Enabled,
	}
	if lastId, err = m.userDAO.Insert(do); err != nil {
		log.Errorf(err.Error())
		return
	}
	do.SetId(int32(lastId))
	if err = m.userRoleDAO.Insert(&models.UserRoleDO{
		UserName:   InherentAdminName,
		RoleName:   InherentAdminRole,
		CreateUser: InherentAdminName,
		EditUser:   InherentAdminName,
		Deleted:    false,
	}); err != nil {
		log.Errorf(err.Error())
		return
	}
}

func (m *PlatformCore) UserLogin(strUserName, strPassword, strIP string) (user *models.UserDO, code types.BizCode) {

	var err error
	if user, err = m.userDAO.SelectUserByName(strUserName); err != nil {
		log.Errorf("select [user name] from user table error [%s]", err.Error())
		return nil, types.NewBizCode(types.CODE_INVALID_USER_OR_PASSWORD)
	}

	if user == nil || user.GetId() == 0 {
		if user, err = m.userDAO.SelectUserByPhone(strUserName); err != nil {
			log.Errorf("select [phone] from user table error [%s]", err.Error())
			return nil, types.NewBizCode(types.CODE_DATABASE_ERROR)
		}
		if user == nil || user.GetId() == 0 {
			log.Errorf("user name/email [%s] data not found in db", strUserName)
			return nil, types.NewBizCode(types.CODE_INVALID_USER_OR_PASSWORD)
		}
	}

	if user.State == dao.UserState_Disabled {
		log.Errorf("user name/email [%s] account was disabled", strUserName)
		return nil, types.NewBizCode(types.CODE_ACCOUNT_BANNED)
	}

	user.LoginIp = strIP
	user.LoginTime = time.Now().Unix()
	if err = m.userDAO.UpdateByName(user, models.USER_COLUMN_LOGIN_IP, models.USER_COLUMN_LOGIN_TIME); err != nil {
		log.Errorf("update user [%s] login ip error [%s]", strUserName, strIP)
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}

	if strPassword != user.Password {
		err = fmt.Errorf("user name [%s] password verify failed, password [%s] not match", strUserName, strPassword)
		log.Errorf(err.Error())
		return nil, types.NewBizCode(types.CODE_INVALID_USER_OR_PASSWORD)
	}
	_, _ = m.loginDAO.Insert(&models.LoginDO{
		UserId:    user.GetId(),
		LoginIp:   strIP,
		LoginAddr: "",
	})
	return
}

func (m *PlatformCore) GetUserRole(ctx *types.Context, strUserName string) (role *models.RoleDO) {
	var err error

	fmt.Println(ctx.UserName())
	if role, err = m.roleDAO.SelectUserRole(strUserName); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (m *PlatformCore) CheckPrivilege(c *gin.Context, ctx *types.Context, strUserName string, authority privilege.Privilege) (ok bool) {
	if authority == privilege.Null {
		return true
	}
	var err error
	//获取请求的URI
	obj := c.Request.URL.RequestURI()
	//获取权限
	act := string(authority)
	//获取用户名
	sub := strUserName

	if ok, err = privilege.CasRule.Enforce(sub, obj, act); ok {
		return true
	}
	if err != nil {
		log.Errorf("user id [%d] privilege [%v] error:%s", ctx.UserId(), authority, err.Error())
	}

	log.Warnf("user id [%d] have no privilege [%v]", ctx.UserId(), authority)
	return
}

func (m *PlatformCore) CheckUserNameExist(ctx *types.Context, strUserName string) (code types.BizCode) {
	if ok, err := m.userDAO.CheckActiveUserByUserName(strUserName); err != nil {
		log.Error("query by user name %s error [%s]", strUserName, err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	} else if ok {
		return types.NewBizCode(types.CODE_ALREADY_EXIST)
	}
	return types.BizOK
}

func (m *PlatformCore) CheckUserEmailExist(ctx *types.Context, strEmail string) (code types.BizCode) {

	if ok, err := m.userDAO.CheckActiveUserByEmail(strEmail); err != nil {
		log.Error("query by email error [%s]", err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	} else if ok {
		return types.NewBizCode(types.CODE_ALREADY_EXIST, "email already exists")
	}
	return types.BizOK
}

func (m *PlatformCore) CheckUserPhoneExist(ctx *types.Context, strPhone string) (code types.BizCode) {

	if ok, err := m.userDAO.CheckActiveUserByPhone(strPhone); err != nil {
		log.Error(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	} else if ok {
		return types.NewBizCode(types.CODE_ALREADY_EXIST)
	}
	return types.BizOK
}

func (m *PlatformCore) ListUser(ctx *types.Context, req *proto.PlatformListUserReq) (totalUsers []*proto.PlatformTotalUser, total int64, code types.BizCode) {
	var err error
	var users = make([]*proto.PlatformUser, 0)
	totalUsers = make([]*proto.PlatformTotalUser, 0)
	if users, total, err = m.userRoleDAO.SelectUsers(req); err != nil {
		log.Errorf(err.Error())
		return nil, 0, types.NewBizCodeDatabaseError(err.Error())
	}
	for _, user := range users {
		sysUser := &proto.PlatformTotalUser{
			UserId:      user.UserId,
			UserName:    user.UserName,
			UserAlias:   user.UserAlias,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Remark:      user.UserRemark,
			LoginTime:   user.LoginTime,
			RoleName:    user.RoleName,
			Password:    user.Password,
			CreateUser:  user.CreateUser,
			State:       user.State,
			CreatedTime: user.CreatedTime,
		}
		totalUsers = append(totalUsers, sysUser)
	}
	return
}

func (m *PlatformCore) CreateUser(ctx *types.Context, req *proto.PlatformCreateUserReq) (do *models.UserDO, code types.BizCode) {

	var err error
	var role *models.RoleDO

	if req.PhoneNumber != "" && !utils.VerifyMobileFormat(req.PhoneNumber) {
		return nil, types.NewBizCode(types.CODE_INVALID_PARAMS, "invalid phone number")
	}

	if req.RoleName != "" {
		if role, err = m.roleDAO.SelectRoleByName(req.RoleName); err != nil {
			log.Errorf(err.Error())
			return nil, types.NewBizCodeDatabaseError(err.Error())
		}
		if role == nil || role.GetId() == 0 {
			log.Errorf("role name [%s] not found", req.RoleName)
			return nil, types.NewBizCode(types.CODE_NOT_FOUND)
		}

		//TODO: merge two operations into one transaction later

		// 添加用户角色
		privilege.AddUserRole(req.UserName, req.RoleName)
	}

	var lastId int64
	var strSalt = utils.GenerateSalt()
	do = &models.UserDO{
		UserName:    req.UserName,
		Password:    req.Password,
		Salt:        strSalt,
		UserAlias:   req.UserAlias,
		IsAdmin:     false,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Remark:      req.Remark,
		CreateUser:  ctx.UserName(),
		EditUser:    ctx.UserName(),
		LoginTime:   utils.Now64(),
		Deleted:     false,
		State:       dao.UserState_Enabled,
	}
	if lastId, err = m.userDAO.Insert(do); err != nil {
		log.Errorf(err.Error())
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}
	do.SetId(int32(lastId))
	if err = m.userRoleDAO.Insert(&models.UserRoleDO{
		UserName:   req.UserName,
		RoleName:   req.RoleName,
		CreateUser: ctx.UserName(),
		EditUser:   ctx.UserName(),
		Deleted:    false,
	}); err != nil {
		log.Errorf(err.Error())
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}
	return do, types.BizOK
}

func (m *PlatformCore) EditUser(ctx *types.Context, req *proto.PlatformEditUserReq) (code types.BizCode) {
	var err error
	var user *models.UserDO
	var userRoleDo = make([]*models.UserRoleDO, 0)
	if user, code = m.GetUserByName(ctx, req.UserName); !code.Ok() {
		log.Errorf("edit user return code [%v]", code)
		return code
	}

	if req.PhoneNumber != "" && user.PhoneNumber != req.PhoneNumber {
		if !utils.VerifyMobileFormat(req.PhoneNumber) {
			return types.NewBizCode(types.CODE_INVALID_PARAMS, "invalid phone number")
		}
	}

	if user.Email != req.Email {
		if !utils.VerifyEmailFormat(req.Email) {
			return types.NewBizCode(types.CODE_INVALID_PARAMS, "invalid email")
		}
		var ok bool
		if ok, err = m.userDAO.CheckActiveUserByEmail(req.Email); err != nil {
			log.Errorf(err.Error())
			return types.NewBizCodeDatabaseError(err.Error())
		} else if ok {
			return types.NewBizCode(types.CODE_ALREADY_EXIST)
		}
	}

	if userRoleDo, err = m.userRoleDAO.SelectUserByName(req.UserName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}

	user.PhoneNumber = req.PhoneNumber
	user.Remark = req.Remark
	user.UserAlias = req.UserAlias
	user.Password = req.Password
	user.Email = req.Email
	user.EditUser = ctx.UserName()
	if err = m.userDAO.UpdateByName(user,
		models.USER_COLUMN_USER_ALIAS,
		models.USER_COLUMN_REMARK,
		models.USER_COLUMN_PASSWORD,
		models.USER_COLUMN_EMAIL,
		models.USER_COLUMN_PHONE_NUMBER,
		models.USER_COLUMN_EDIT_USER); err != nil {

		log.Errorf("%s", err)
		return types.NewBizCodeDatabaseError(err.Error())
	}

	// 用户角色更新
	for _, userRole := range userRoleDo {
		privilege.DeleteUserRole(userRole.UserName, userRole.RoleName)
		privilege.AddUserRole(userRole.UserName, req.RoleName)
	}
	//upsert a new role of platform
	if err = m.userRoleDAO.Upsert(&models.UserRoleDO{
		UserName:   user.GetUserName(),
		RoleName:   req.RoleName,
		Deleted:    false,
		CreateUser: ctx.UserName(),
		EditUser:   ctx.UserName(),
	}, models.USER_ROLE_COLUMN_ROLE_NAME, models.USER_ROLE_COLUMN_EDIT_USER); err != nil {

		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) ResetUserPassword(ctx *types.Context, req *proto.PlatformResetPasswordReq) (code types.BizCode) {
	var err error
	var user *models.UserDO

	if user, code = m.GetUserByName(ctx, req.UserName); !code.Ok() {
		log.Errorf("reset user password return code [%v]", code)
		return code
	}
	user.EditUser = ctx.UserName()
	user.Password = req.NewPassword
	if err = m.userDAO.UpdateByName(user, models.USER_COLUMN_PASSWORD, models.USER_COLUMN_EDIT_USER); err != nil {
		log.Error(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) EnableUser(ctx *types.Context, req *proto.PlatformEnableUserReq) (r *proto.PlatformEnableUserResp, code types.BizCode) {
	if err := m.userDAO.UpdateUserState(req.UserName, dao.UserState_Enabled); err != nil {
		log.Errorf(err.Error())
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}
	return &proto.PlatformEnableUserResp{}, types.BizOK
}

func (m *PlatformCore) DisableUser(ctx *types.Context, req *proto.PlatformDisableUserReq) (r *proto.PlatformDisableUserResp, code types.BizCode) {
	if err := m.userDAO.UpdateUserState(req.UserName, dao.UserState_Disabled); err != nil {
		log.Errorf(err.Error())
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}
	return &proto.PlatformDisableUserResp{}, types.BizOK
}

func (m *PlatformCore) DeleteUser(ctx *types.Context, req *proto.PlatformDeleteUserReq) (code types.BizCode) {
	var err error
	var user *models.UserDO
	if req.UserName == "" {
		err = log.Errorf("user name to delete is nil")
		return types.NewBizCode(types.CODE_INVALID_PARAMS, err.Error())
	}

	if user, code = m.GetUserByName(ctx, req.UserName); !code.Ok() {
		log.Errorf("delete user return code [%v] ", code)
		return code
	}

	if user.GetId() == 0 {
		err = log.Errorf("user %s not found", req.UserName)
		return types.NewBizCode(types.CODE_INVALID_PARAMS, err.Error())
	}
	if err = m.userRoleDAO.Delete(&models.UserRoleDO{
		UserName: req.UserName,
		EditUser: ctx.UserName(),
		Deleted:  true,
	}); err != nil {
		log.Error("delete role by user name error [%s]", err)
		return types.NewBizCodeDatabaseError(err.Error())
	}
	privilege.DeleteUser(req.UserName)
	user.Deleted = true
	user.EditUser = ctx.UserName()
	if err = m.userDAO.DeleteUser(user); err != nil {
		log.Errorf("delete user error [%s]", err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) GetUserByName(ctx *types.Context, strUserName string) (do *models.UserDO, code types.BizCode) {

	var err error
	do, err = m.userDAO.SelectUserByName(strUserName)
	if err != nil {
		log.Errorf(err.Error())
		return nil, types.NewBizCodeDatabaseError(err.Error())
	}

	if do == nil || do.GetId() == 0 {
		log.Errorf("user [%s] not found", strUserName)
		return nil, types.NewBizCode(types.CODE_NOT_FOUND)
	}
	return do, types.BizOK
}

func (m *PlatformCore) ListRole(ctx *types.Context, req *proto.PlatformListRoleReq) (roleLists []*proto.PlatformSysRole, total int64, code types.BizCode) {
	var err error
	log.Debugf("request processing...")
	var roles = make([]*proto.PlatformRole, 0)
	roleLists = make([]*proto.PlatformSysRole, 0)
	if roles, total, err = m.roleDAO.SelectPlatformRoles(req.PageNo, req.PageSize, req.RoleName); err != nil {
		log.Errorf(err.Error())

		return
	}
	for _, role := range roles {
		auth := privilege.GetRoleAuthority(role.RoleName)
		roleLists = append(roleLists, &proto.PlatformSysRole{
			Id:          role.Id,
			RoleName:    role.RoleName,
			CreateUser:  role.CreateUser,
			Remark:      role.Remark,
			CreatedTime: role.CreatedTime,
			Role:        auth,
		})
	}
	return
}

func (m *PlatformCore) CheckUserPassword(ctx *types.Context, strUserName, strPassword string) (ok bool, code types.BizCode) {
	var do *models.UserDO

	if do, code = m.GetUserByName(ctx, strUserName); !code.Ok() {
		log.Errorf("check user password return code [%v]", code)
		return
	}
	if strPassword != do.Password {
		return false, types.BizOK
	}
	return true, types.BizOK
}

func (m *PlatformCore) CreateRole(ctx *types.Context, req *proto.PlatformCreateRoleReq) (code types.BizCode) {
	var err error
	var exist bool

	if exist, err = m.roleDAO.CheckRoleExistByName(req.RoleName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}

	if exist {
		log.Errorf("role name [%s] already exist", req.RoleName)
		return types.NewBizCode(types.CODE_ALREADY_EXIST)
	}

	if _, err = m.roleDAO.Insert(&models.RoleDO{
		RoleName:   req.RoleName,
		CreateUser: ctx.UserName(),
		EditUser:   ctx.UserName(),
		Remark:     req.Remark,
	}); err != nil {

		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) EditRole(ctx *types.Context, req *proto.PlatformEditRoleReq) (code types.BizCode) {
	var err error
	var ok bool
	var do *models.RoleDO
	// 查询role表获取信息
	if do, err = m.roleDAO.SelectRoleById(req.Id); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	// 若角色名称相同则直接更新描述
	if do.RoleName == req.RoleName {
		do.Remark = req.Remark
		if err = m.roleDAO.Update(do,
			models.ROLE_COLUMN_ROLE_NAME,
			models.ROLE_COLUMN_REMARK); err != nil {
			log.Errorf(err.Error())
			return types.NewBizCodeDatabaseError(err.Error())
		}
		return types.BizOK
	}
	// 检查名称是否存在
	if ok, err = m.roleDAO.CheckRoleExistByName(req.RoleName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	if ok {
		log.Errorf("role name [%s] already exist", req.RoleName)
		return types.NewBizCode(types.CODE_ALREADY_EXIST, "role name already exist")
	}
	//角色权限继承 roleA 继承/获取 roleB权限
	privilege.InheritRoleAuthority(req.RoleName, do.RoleName)
	// 角色更名时，用户继承新角色,并删除旧角色
	privilege.InheritUserRole(req.RoleName, do.RoleName)
	// 更新用户角色表
	var dos = make([]*models.UserRoleDO, 0)
	if dos, err = m.userRoleDAO.SelectUserByRole(do.RoleName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	for _, doRole := range dos {
		doRole.RoleName = req.RoleName
		err = m.userRoleDAO.UpdateUserById(doRole)
		if err != nil {
			log.Errorf(err.Error())
			continue
		}
	}

	do.RoleName = req.RoleName
	do.Remark = req.Remark
	// 更新role表
	if err = m.roleDAO.Update(do,
		models.ROLE_COLUMN_ROLE_NAME,
		models.ROLE_COLUMN_REMARK,
	); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) DeleteRole(ctx *types.Context, req *proto.PlatformDeleteRoleReq) (code types.BizCode) {
	var exist bool
	var err error

	if exist, err = m.roleDAO.CheckRoleExistByName(req.RoleName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	if !exist {
		log.Warnf("role name [%s] not found", req.RoleName)
		return types.NewBizCode(types.CODE_NOT_FOUND, "role name not found")
	}

	var dos = make([]*models.UserRoleDO, 0)
	if dos, err = m.userRoleDAO.SelectUserByRole(req.RoleName); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	if len(dos) > 0 {
		log.Warnf("role: %s had %d user", req.RoleName, len(dos))
		return types.NewBizCode(types.CODE_ACCESS_VIOLATE, "role has too many users to delete")
	}
	// 删除角色
	privilege.DeleteRole(req.RoleName)

	if err = m.roleDAO.Delete(&models.RoleDO{
		RoleName: req.RoleName,
		Deleted:  true,
		EditUser: ctx.UserName(),
	}); err != nil {
		log.Errorf(err.Error())
		return types.NewBizCodeDatabaseError(err.Error())
	}
	return types.BizOK
}

func (m *PlatformCore) AuthRole(ctx *types.Context, req *proto.PlatformAuthRoleReq) (code types.BizCode) {
	if len(req.Privilege) == 0 {
		privilege.DeleteRole(req.RoleName)
		return types.BizOK
	}
	// 获取具有角色的用户
	res, err := privilege.CasRule.GetUsersForRole(req.RoleName)
	if err != nil {
		log.Errorf("get users for role:%s error:%s", req.RoleName, err.Error())
	}
	privilege.DeleteRole(req.RoleName)
	for _, authority := range req.Privilege {
		role := req.RoleName
		path := privilege.GetKeyMatchPath(authority)
		if strings.EqualFold(path, "") {
			log.Errorf("unknown authority : %s", authority)
			continue
		}
		// 给角色授予权限
		privilege.AddRoleAuthority(role, path, authority)
		if authority == privilege.UserAdd || authority == privilege.UserEdit {
			strRolePath := privilege.GetKeyMatchPath(privilege.RoleAccess)
			privilege.AddRoleAuthority(role, strRolePath, privilege.RoleAccess) //角色拥有用户添加和编辑权限默认附带角色查看权限
		}
	}
	for _, user := range res {
		// 给用户添加角色
		privilege.AddUserRole(user, req.RoleName)
	}
	return types.BizOK
}

func (m *PlatformCore) InquireAuth(ctx *types.Context, req *proto.PlatformInquireAuthReq) (auth []string, code types.BizCode) {
	auth = make([]string, 0)
	switch req.NameType {
	case proto.TypeUser:
		auth = privilege.GetUserRoleList(req.Name)
	case proto.TypeRole:
		auth = privilege.GetRoleAuthority(req.Name)
	default:
		return auth, types.NewBizCode(types.CODE_TYPE_UNDEFINED, "auth type not defined")
	}
	return auth, types.BizOK
}

func (m *PlatformCore) UserQuery(ctx *types.Context, req *proto.PlatformUserQueryReq) (userList proto.PlatformUserQueryResp, total int64, code types.BizCode) {
	var err error
	var userName = make([]*models.UserDO, 0)
	userList = proto.PlatformUserQueryResp{
		NameList: make([]string, 0),
	}
	if userName, total, err = m.userDAO.SelectUserName(req); err != nil {
		log.Errorf("select user name list error:%s\n", err.Error())
		return
	}
	for _, name := range userName {
		userList.NameList = append(userList.NameList, name.UserName)
	}
	return
}

func (m *PlatformCore) PrivilegeLevel(ctx *types.Context, req *proto.PlatformPrivilegeLevelReq) (typeList proto.PlatformPrivilegeLevelResp, code types.BizCode) {
	typeList = proto.PlatformPrivilegeLevelResp{
		TreeList: make(privilege.TreePrivilege, 0),
	}
	typeList.TreeList = privilege.Tree()
	return
}

func (m *PlatformCore) EmailConfigList(ctx *types.Context, req *proto.PlatformEmailConfigListReq) (emailConfig proto.PlatformEmailConfigListResp, code types.BizCode) {
	emailConfig = m.dictionaryDAO.SelectEmailConfig()
	return
}

func (m *PlatformCore) EmailConfig(ctx *types.Context, req *proto.PlatformEmailConfigReq) (emailConfig proto.PlatformEmailConfigResp, code types.BizCode) {
	if !utils.VerifyEmailFormat(req.EmailName) {
		return emailConfig, types.NewBizCode(types.CODE_INVALID_PARAMS, "malformed email address")
	}
	var err error
	emailServerDo := &models.DictionaryDO{
		Name:      dao.Dictionary_Name_Email_Server,
		ConfigKey: dao.Dictionary_Key_Email_Server,
		Value:     req.EmailServer,
		Remark:    dao.Dictionary_Remark_Email_Server,
		Deleted:   false,
	}
	err = m.dictionaryDAO.Upsert(emailServerDo)
	if err != nil {
		log.Errorf("emailServerDo error:%s", err.Error())
		return emailConfig, types.NewBizCodeDatabaseError(err.Error())
	}
	portDo := &models.DictionaryDO{
		Name:      dao.Dictionary_Name_Email_Port,
		ConfigKey: dao.Dictionary_Key_Email_Port,
		Value:     req.Port,
		Remark:    dao.Dictionary_Remark_Email_Port,
		Deleted:   false,
	}
	err = m.dictionaryDAO.Upsert(portDo)
	if err != nil {
		log.Errorf("portDo error:%s", err.Error())
		return emailConfig, types.NewBizCodeDatabaseError(err.Error())
	}
	emailNameDo := &models.DictionaryDO{
		Name:      dao.Dictionary_Name_Email_Name,
		ConfigKey: dao.Dictionary_Key_Email_Name,
		Value:     req.EmailName,
		Remark:    dao.Dictionary_Remark_Email_Name,
		Deleted:   false,
	}
	err = m.dictionaryDAO.Upsert(emailNameDo)
	if err != nil {
		log.Errorf("emailServerDo error:%s", err.Error())
		return emailConfig, types.NewBizCodeDatabaseError(err.Error())
	}
	autoCodeDo := &models.DictionaryDO{
		Name:      dao.Dictionary_Name_Email_Auth_Code,
		ConfigKey: dao.Dictionary_Key_Email_Auth_Code,
		Value:     req.AuthCode,
		Remark:    dao.Dictionary_Remark_Email_Auth_Code,
		Deleted:   false,
	}
	err = m.dictionaryDAO.Upsert(autoCodeDo)
	if err != nil {
		log.Errorf("autoCodeDo error:%s\n", err.Error())
		return emailConfig, types.NewBizCodeDatabaseError(err.Error())
	}
	SendNameDo := &models.DictionaryDO{
		Name:      dao.Dictionary_Name_Email_Send_Name,
		ConfigKey: dao.Dictionary_Key_Email_Send_Name,
		Value:     req.SendName,
		Remark:    dao.Dictionary_Remark_Email_Send_Name,
		Deleted:   false,
	}
	err = m.dictionaryDAO.Upsert(SendNameDo)
	if err != nil {
		log.Errorf("emailServerDo error:%s\n", err.Error())
		return emailConfig, types.NewBizCodeDatabaseError(err.Error())
	}
	return
}

func (m *PlatformCore) CheckExist(ctx *types.Context, req *proto.PlatformCheckExistReq) (code types.BizCode) {
	switch req.CheckType {
	case types.CheckType_UserName:
		{
			if do, err := m.userDAO.SelectUserByName(req.Name); err != nil {
				log.Errorf(err.Error())
				return types.NewBizCodeDatabaseError(err.Error())
			} else if do != nil && do.GetId() != 0 {
				return types.NewBizCode(types.CODE_ALREADY_EXIST)
			}
		}
	case types.CheckType_UserPhoneNumber:
		{
			if do, err := m.userDAO.SelectUserByPhone(req.Name); err != nil {
				log.Errorf(err.Error())
				return types.NewBizCodeDatabaseError(err.Error())
			} else if do != nil && do.GetId() != 0 {
				return types.NewBizCode(types.CODE_ALREADY_EXIST)
			}
		}
	case types.CheckType_UserEmail:
		{
			if do, err := m.userDAO.SelectUserByEmail(req.Name); err != nil {
				log.Errorf(err.Error())
				return types.NewBizCodeDatabaseError(err.Error())
			} else if do != nil && do.GetId() != 0 {
				return types.NewBizCode(types.CODE_ALREADY_EXIST)
			}
		}
	case types.CheckType_RoleName:
		{
			if do, err := m.roleDAO.SelectRoleByName(req.Name); err != nil {
				log.Errorf(err.Error())
				return types.NewBizCodeDatabaseError(err.Error())
			} else if do != nil && do.GetId() != 0 {
				return types.NewBizCode(types.CODE_ALREADY_EXIST)
			}
		}
	}
	return types.NewBizCode(types.CODE_NOT_FOUND)
}

func (m *PlatformCore) ListRoleUser(ctx *types.Context, req *proto.PlatformListRoleUserReq) (users []*proto.PlatformUser, total int64, code types.BizCode) {
	var err error

	if users, total, err = m.userRoleDAO.SelectRoleUsers(req.RoleName, req.PageNo, req.PageSize); err != nil {
		log.Errorf(err.Error())
		return nil, 0, types.NewBizCodeDatabaseError(err.Error())
	}
	return
}

func (m *PlatformCore) ListOperLog(ctx *types.Context, req *proto.PlatformListOperLogReq) (list []*proto.OperLog, total int64, code types.BizCode) {
	var err error
	_ = err
	return
}
