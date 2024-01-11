package routers

import (
	"admin-system/pkg/api"
	"admin-system/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const (
	GroupRouterPlatformV1      = "/api/v1/platform"
	RouterSubPathPlatformLogin = "/api/v1/login"
)

const ( //prefix http://localhost:port/api/v1/platform

	RouterSubPathPlatformLogout           = "/logout"
	RouterSubPathPlatformCheckExist       = "/check/exist"
	RouterSubPathPlatformListUser         = "/list/user"
	RouterSubPathPlatformCreateUser       = "/create/user"
	RouterSubPathPlatformEditUser         = "/edit/user"
	RouterSubPathPlatformEnableUser       = "/enable/user"
	RouterSubPathPlatformDisableUser      = "/disable/user"
	RouterSubPathPlatformDeleteUser       = "/delete/user"
	RouterSubPathPlatformListRole         = "/list/role"
	RouterSubPathPlatformCreateRole       = "/create/role"
	RouterSubPathPlatformEditRole         = "/edit/role"
	RouterSubPathPlatformDeleteRole       = "/delete/role"
	RouterSubPathPlatformAuthRole         = "/auth/role"
	RouterSubPathPlatformInquireAuth      = "/inquire/auth"
	RouterSubPathPlatformPrivilegeLevel   = "/privilege/level"
	RouterSubPathPlatformResetPassword    = "/reset/password"
	RouterSubPathPlatformChangePassword   = "/change/password"
	RouterSubPathPlatformListRoleUser     = "/list/role-user"
	RouterSubPathPlatformRefreshAuthToken = "/refresh/token"
	RouterSubPathPlatformListOperLog      = "/list/oper-log"
)

func InitRouterGroupPlatform(r *gin.Engine, handlers api.PlatformApi) {

	r.POST(RouterSubPathPlatformLogin, handlers.PlatformLogin) //do not need JWT authentication

	g := r.Group(GroupRouterPlatformV1)
	g.Use(middleware.JWT()) //use JWT token middleware
	{
		g.POST(RouterSubPathPlatformLogout, handlers.PlatformLogout)
		g.POST(RouterSubPathPlatformCheckExist, handlers.PlatformCheckExist)
		g.POST(RouterSubPathPlatformListUser, handlers.PlatformListUser)
		g.POST(RouterSubPathPlatformCreateUser, handlers.PlatformCreateUser)
		g.POST(RouterSubPathPlatformEditUser, handlers.PlatformEditUser)
		g.POST(RouterSubPathPlatformEnableUser, handlers.PlatformEnableUser)
		g.POST(RouterSubPathPlatformDisableUser, handlers.PlatformDisableUser)
		g.POST(RouterSubPathPlatformDeleteUser, handlers.PlatformDeleteUser)
		g.POST(RouterSubPathPlatformListRole, handlers.PlatformListRole)
		g.POST(RouterSubPathPlatformCreateRole, handlers.PlatformCreateRole)
		g.POST(RouterSubPathPlatformEditRole, handlers.PlatformEditRole)
		g.POST(RouterSubPathPlatformDeleteRole, handlers.PlatformDeleteRole)
		g.POST(RouterSubPathPlatformAuthRole, handlers.PlatformAuthRole)
		g.POST(RouterSubPathPlatformInquireAuth, handlers.PlatformInquireAuth)
		g.POST(RouterSubPathPlatformPrivilegeLevel, handlers.PlatformPrivilegeLevel)
		g.POST(RouterSubPathPlatformResetPassword, handlers.PlatformResetPassword)
		g.POST(RouterSubPathPlatformChangePassword, handlers.PlatformChangePassword)
		g.POST(RouterSubPathPlatformListRoleUser, handlers.PlatformListRoleUser)
		g.POST(RouterSubPathPlatformRefreshAuthToken, handlers.PlatformRefreshAuthToken)
		g.POST(RouterSubPathPlatformListOperLog, handlers.PlatformListOperLog)
	}
	return
}
