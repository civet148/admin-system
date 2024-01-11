package api

import "github.com/gin-gonic/gin"

type PlatformApi interface {
	PlatformLoginV1(c *gin.Context)            // 用户登录
	PlatformLogoutV1(c *gin.Context)           // 用户登出
	PlatformCheckExistV1(c *gin.Context)       // 检查用户名/邮箱是否存在
	PlatformListUserV1(c *gin.Context)         // 列出所有系统用户
	PlatformCreateUserV1(c *gin.Context)       // 创建系统用户
	PlatformEditUserV1(c *gin.Context)         // 编辑系统用户
	PlatformEnableUserV1(c *gin.Context)       // 启用系统用户
	PlatformDisableUserV1(c *gin.Context)      // 禁用系统用户
	PlatformDeleteUserV1(c *gin.Context)       // 删除系统用户
	PlatformListRoleV1(c *gin.Context)         // 列出所有角色
	PlatformCreateRoleV1(c *gin.Context)       // 创建角色
	PlatformEditRoleV1(c *gin.Context)         // 编辑角色
	PlatformDeleteRoleV1(c *gin.Context)       // 删除角色
	PlatformAuthRoleV1(c *gin.Context)         // 角色授权
	PlatformInquireAuthV1(c *gin.Context)      // 查询权限
	PlatformPrivilegeLevelV1(c *gin.Context)   // 权限结构
	PlatformResetPasswordV1(c *gin.Context)    // 管理员充值系统用户密码
	PlatformChangePasswordV1(c *gin.Context)   // 系统用户自行修改密码
	PlatformListRoleUserV1(c *gin.Context)     // 列出角色对应用户
	PlatformRefreshAuthTokenV1(c *gin.Context) // 刷新访问TOKEN
	PlatformListOperLogV1(c *gin.Context)      // 操作日志列表
}
