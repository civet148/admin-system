package api

import "github.com/gin-gonic/gin"

type PlatformApi interface {
	PlatformLogin(c *gin.Context)            // 用户登录
	PlatformLogout(c *gin.Context)           // 用户登出
	PlatformCheckExist(c *gin.Context)       // 检查用户名/邮箱是否存在
	PlatformListUser(c *gin.Context)         // 列出所有系统用户
	PlatformCreateUser(c *gin.Context)       // 创建系统用户
	PlatformEditUser(c *gin.Context)         // 编辑系统用户
	PlatformEnableUser(c *gin.Context)       // 启用系统用户
	PlatformDisableUser(c *gin.Context)      // 禁用系统用户
	PlatformDeleteUser(c *gin.Context)       // 删除系统用户
	PlatformListRole(c *gin.Context)         // 列出所有角色
	PlatformCreateRole(c *gin.Context)       // 创建角色
	PlatformEditRole(c *gin.Context)         // 编辑角色
	PlatformDeleteRole(c *gin.Context)       // 删除角色
	PlatformAuthRole(c *gin.Context)         // 角色授权
	PlatformInquireAuth(c *gin.Context)      // 查询权限
	PlatformPrivilegeLevel(c *gin.Context)   // 权限结构
	PlatformResetPassword(c *gin.Context)    // 管理员充值系统用户密码
	PlatformChangePassword(c *gin.Context)   // 系统用户自行修改密码
	PlatformListRoleUser(c *gin.Context)     // 列出角色对应用户
	PlatformRefreshAuthToken(c *gin.Context) // 刷新访问TOKEN
	PlatformListOperLog(c *gin.Context)      // 操作日志列表
}
