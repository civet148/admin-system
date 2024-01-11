package routers

import (
	"admin-system/pkg/api"
	"admin-system/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const (
	GroupRouterPlatform          = "/api/v1/platform"
	RouterSubPathPlatformLoginV1 = "/api/v1/login"
)

const ( //prefix http://localhost:port/api/v1/platform

	RouterSubPathPlatformLogoutV1           = "/logout"
	RouterSubPathPlatformCheckExistV1       = "/check/exist"
	RouterSubPathPlatformListUserV1         = "/list/user"
	RouterSubPathPlatformCreateUserV1       = "/create/user"
	RouterSubPathPlatformEditUserV1         = "/edit/user"
	RouterSubPathPlatformEnableUserV1       = "/enable/user"
	RouterSubPathPlatformDisableUserV1      = "/disable/user"
	RouterSubPathPlatformDeleteUserV1       = "/delete/user"
	RouterSubPathPlatformListRoleV1         = "/list/role"
	RouterSubPathPlatformCreateRoleV1       = "/create/role"
	RouterSubPathPlatformEditRoleV1         = "/edit/role"
	RouterSubPathPlatformDeleteRoleV1       = "/delete/role"
	RouterSubPathPlatformAuthRoleV1         = "/auth/role"
	RouterSubPathPlatformInquireAuthV1      = "/inquire/auth"
	RouterSubPathPlatformPrivilegeLevelV1   = "/privilege/level"
	RouterSubPathPlatformResetPasswordV1    = "/reset/password"
	RouterSubPathPlatformChangePasswordV1   = "/change/password"
	RouterSubPathPlatformListRoleUserV1     = "/list/role-user"
	RouterSubPathPlatformRefreshAuthTokenV1 = "/refresh/token"
	RouterSubPathPlatformListOperLogV1      = "/list/oper-log"
)

func InitRouterGroupPlatform(r *gin.Engine, handlers api.PlatformApi) {

	r.POST(RouterSubPathPlatformLoginV1, handlers.PlatformLoginV1) //do not need JWT authentication

	g := r.Group(GroupRouterPlatform)
	g.Use(middleware.JWT()) //use JWT token middleware
	{
		g.POST(RouterSubPathPlatformLogoutV1, handlers.PlatformLogoutV1)
		g.POST(RouterSubPathPlatformCheckExistV1, handlers.PlatformCheckExistV1)
		g.POST(RouterSubPathPlatformListUserV1, handlers.PlatformListUserV1)
		g.POST(RouterSubPathPlatformCreateUserV1, handlers.PlatformCreateUserV1)
		g.POST(RouterSubPathPlatformEditUserV1, handlers.PlatformEditUserV1)
		g.POST(RouterSubPathPlatformEnableUserV1, handlers.PlatformEnableUserV1)
		g.POST(RouterSubPathPlatformDisableUserV1, handlers.PlatformDisableUserV1)
		g.POST(RouterSubPathPlatformDeleteUserV1, handlers.PlatformDeleteUserV1)
		g.POST(RouterSubPathPlatformListRoleV1, handlers.PlatformListRoleV1)
		g.POST(RouterSubPathPlatformCreateRoleV1, handlers.PlatformCreateRoleV1)
		g.POST(RouterSubPathPlatformEditRoleV1, handlers.PlatformEditRoleV1)
		g.POST(RouterSubPathPlatformDeleteRoleV1, handlers.PlatformDeleteRoleV1)
		g.POST(RouterSubPathPlatformAuthRoleV1, handlers.PlatformAuthRoleV1)
		g.POST(RouterSubPathPlatformInquireAuthV1, handlers.PlatformInquireAuthV1)
		g.POST(RouterSubPathPlatformPrivilegeLevelV1, handlers.PlatformPrivilegeLevelV1)
		g.POST(RouterSubPathPlatformResetPasswordV1, handlers.PlatformResetPasswordV1)
		g.POST(RouterSubPathPlatformChangePasswordV1, handlers.PlatformChangePasswordV1)
		g.POST(RouterSubPathPlatformListRoleUserV1, handlers.PlatformListRoleUserV1)
		g.POST(RouterSubPathPlatformRefreshAuthTokenV1, handlers.PlatformRefreshAuthTokenV1)
		g.POST(RouterSubPathPlatformListOperLogV1, handlers.PlatformListOperLogV1)
	}
	return
}
