package proto

import (
	"admin-system/pkg/privilege"
	"admin-system/pkg/types"
)

type PlatformLoginReq struct {
	UserName string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PlatformLoginResp struct {
	Version   string   `json:"version"`
	UserName  string   `json:"user_name" db:"user_name" bson:"user_name"`
	AuthToken string   `json:"auth_token" db:"auth_token" bson:"auth_token"`
	LoginIp   string   `json:"login_ip" db:"login_ip" bson:"login_ip"`       //最近登录IP
	LoginTime int64    `json:"login_time" db:"login_time" bson:"login_time"` //最近登录时间
	Role      string   `json:"role" db:"role" bson:"role"`
	Privilege []string `json:"privilege" db:"privilege" bson:"privilege"`
}

type PlatformLogoutReq struct {
}

type PlatformLogoutResp struct {
}

type PlatformWorkspaceReq struct {
}

type PlatformWorkspaceResp struct {
}

type PlatformSummaryReq struct {
}

type PlatformSummaryResp struct {
}

type PlatformCheckExistReq struct {
	Name      string          `json:"name" db:"name" bson:"name" binding:"required"` //用户/矿池/角色名称或邮箱地址
	CheckType types.CheckType `json:"check_type" db:"check_type" bson:"check_type"`  //检查类型(0=用户名 1=邮箱 2=角色名称)
}

type PlatformCheckExistResp struct {
}

type PlatformListUserReq struct {
	UserName string `json:"user_name" db:"user_name" bson:"user_name"`
	PageNo   int    `json:"page_no" db:"page_no" bson:"page_no"` //page_no must >= 0
	PageSize int    `json:"page_size" db:"page_size" bson:"page_size"`
}

type PlatformListUserResp struct {
	Users []*PlatformTotalUser `json:"users" db:"users" bson:"users"`
}

type PlatformCreateUserReq struct {
	UserName    string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
	UserAlias   string `json:"user_alias" db:"user_alias" bson:"user_alias" binding:"required"`
	PhoneNumber string `json:"phone_number" db:"phone_number" bson:"phone_number"`
	Email       string `json:"email" db:"email" bson:"email"`
	Password    string `json:"password" db:"password" bson:"password"`
	Remark      string `json:"remark" db:"remark" bson:"remark"`
	RoleName    string `json:"role_name" db:"role_name" bson:"role_name"`
}

type PlatformCreateUserResp struct {
	UserId int32 `json:"user_id" db:"user_id" bson:"user_id"`
}

type PlatformEditUserReq struct {
	UserName    string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
	Password    string `json:"password" db:"password" bson:"password"`
	UserAlias   string `json:"user_alias" db:"user_alias" bson:"user_alias"`
	PhoneNumber string `json:"phone_number" db:"phone_number" bson:"phone_number"` //联系手机号
	Email       string `json:"email" db:"email" bson:"email"`
	Remark      string `json:"remark" db:"remark" bson:"remark"`
	RoleName    string `json:"role_name" db:"role_name" bson:"role_name"`
}

type PlatformEditUserResp struct {
}

type PlatformEnableUserReq struct {
	UserName string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
}

type PlatformEnableUserResp struct {
}

type PlatformDisableUserReq struct {
	UserName string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
}

type PlatformDisableUserResp struct {
}

type PlatformDeleteUserReq struct {
	UserName string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
}

type PlatformDeleteUserResp struct {
}

type PlatformEditUserRoleReq struct {
	UserName string `json:"user_name" db:"user_name" bson:"user_name" binding:"required"`
	RoleName string `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
}

type PlatformEditUserRoleResp struct {
}

type PlatformListRoleReq struct {
	PageNo   int    `json:"page_no" db:"page_no" bson:"page_no"` //page_no must >= 0
	PageSize int    `json:"page_size" db:"page_size" bson:"page_size"`
	RoleName string `json:"role_name" db:"role_name" bson:"role_name"`
}

type PlatformListRoleResp struct {
	Roles []*PlatformSysRole `json:"roles" db:"roles" bson:"roles"`
}

type PlatformCreateRoleReq struct {
	RoleName string `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
	Remark   string `json:"remark" db:"remark" bson:"remark"`
}

type PlatformCreateRoleResp struct {
}

type PlatformEditRoleReq struct {
	Id       int32  `json:"id" db:"id" bson:"id"` //角色ID(自增)
	RoleName string `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
	Remark   string `json:"remark" db:"remark" bson:"remark"`
}

type PlatformEditRoleResp struct {
}

type PlatformDeleteRoleReq struct {
	RoleName string `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
}

type PlatformDeleteRoleResp struct {
}

type PlatformAuthRoleReq struct {
	RoleName  string                `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
	Privilege []privilege.Privilege `json:"privilege" db:"privilege" bson:"privilege" binding:"required"` // 权限列表
}

type PlatformAuthRoleResp struct {
}

const (
	TypeUser = iota // 用户
	TypeRole        // 角色
)

type PlatformInquireAuthReq struct {
	Name     string `json:"name" db:"name" bson:"name" binding:"required"`                //名称
	NameType int    `json:"name_type" db:"name_type" bson:"name_type" binding:"required"` // 类型 0：用户 1：角色
}

type PlatformInquireAuthResp struct {
	Privilege []string `json:"privilege" db:"privilege" bson:"privilege" binding:"required"` // 权限列表
}

type PlatformUserQueryReq struct {
	Name string `json:"name"`
}

type PlatformUserQueryResp struct {
	NameList []string `json:"name_list"`
}

type PlatformPrivilegeLevelReq struct {
}

type PlatformPrivilegeLevelResp struct {
	TreeList privilege.TreePrivilege `json:"tree_list"`
}

type PlatformEmailConfigListReq struct {
}

type PlatformEmailConfigListResp struct {
	EmailServer string `json:"email_server"` // 邮箱服务器
	Port        string `json:"port"`         // 端口
	EmailName   string `json:"email_name"`   // 邮箱名
	AuthCode    string `json:"auth_code"`    // 授权码
	SendName    string `json:"send_name"`    // 发件人名称
}

type PlatformEmailConfigReq struct {
	EmailServer string `json:"email_server"` // 邮箱服务器
	Port        string `json:"port"`         // 端口
	EmailName   string `json:"email_name"`   // 邮箱名
	AuthCode    string `json:"auth_code"`    // 授权码
	SendName    string `json:"send_name"`    // 发件人名称
}

type PlatformEmailConfigResp struct {
}

type PlatformResetPasswordReq struct {
	UserName    string `json:"user_name" db:"user_name" bson:"user_name"`
	OldPassword string `json:"old_password" db:"old_password" bson:"old_password"`
	NewPassword string `json:"new_password" db:"new_password" bson:"new_password" binding:"required"`
}

type PlatformResetPasswordResp struct {
}

type PlatformListRoleUserReq struct {
	RoleName string `json:"role_name" db:"role_name" bson:"role_name" binding:"required"`
	PageNo   int    `json:"page_no" db:"page_no" bson:"page_no"` //page_no must >= 0
	PageSize int    `json:"page_size" db:"page_size" bson:"page_size"`
}

type PlatformListRoleUserResp struct {
	RoleName  string          `json:"role_name" db:"role_name" bson:"role_name"`
	UserCount int             `json:"user_count" db:"user_count" bson:"user_count"`
	Users     []*PlatformUser `json:"users" db:"users" bson:"users"`
}

type PlatformRefreshAuthTokenReq struct {
}

type PlatformRefreshAuthTokenResp struct {
	AuthToken string `json:"auth_token" db:"auth_token"`
}

type OperLog struct {
	OperUser    string `json:"oper_user"`
	OperType    int    `json:"oper_type"`
	OperTime    string `json:"oper_time"`
	OperContent string `json:"oper_content"`
}

type PlatformListOperLogReq struct {
}

type PlatformListOperLogResp struct {
	List []*OperLog `json:"list"`
}
