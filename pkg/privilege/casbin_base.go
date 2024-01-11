package privilege

import (
	"admin-system/pkg/routers"
	"fmt"

	"github.com/civet148/log"
)

type Privilege string

// 权限列表
const (
	Null                Privilege = "Null"                // 无权限校验
	UserAccess          Privilege = "UserAccess"          // 系统管理-用户管理-访问
	UserAdd             Privilege = "UserAdd"             // 系统管理-用户管理-添加
	UserEdit            Privilege = "UserEdit"            // 系统管理-用户管理-编辑
	UserDelete          Privilege = "UserDelete"          // 系统管理-用户管理-删除
	UserEnableOrDisable Privilege = "UserEnableOrDisable" // 系统管理-用户管理-禁用/启用
	RoleAccess          Privilege = "RoleAccess"          // 系统管理-角色管理-访问
	RoleAdd             Privilege = "RoleAdd"             // 系统管理-角色管理-添加
	RoleDelete          Privilege = "RoleDelete"          // 系统管理-角色管理-删除
	RoleEdit            Privilege = "RoleEdit"            // 系统管理-角色管理-编辑
	RoleAuthority       Privilege = "RoleAuthority"       // 系统管理-角色管理-权限授权
	RoleEnableOrDisable Privilege = "RoleEnableOrDisable" // 系统管理-角色管理-禁用/启用
	OperLogAccess       Privilege = "OperLogAccess"       // 系统管理-操作日志访问
)

var TotalPrivilege = []Privilege{
	UserAccess,
	UserAdd,
	UserEdit,
	UserDelete,
	UserEnableOrDisable,
	RoleAccess,
	RoleAdd,
	RoleDelete,
	RoleEdit,
	RoleAuthority,
	RoleEnableOrDisable,
}

func GetKeyMatchPath(auth Privilege) (path string) {
	switch auth {
	case UserAccess, UserAdd, UserEdit, UserDelete, UserEnableOrDisable, RoleAccess, RoleAdd, RoleDelete, RoleEdit, RoleAuthority, OperLogAccess, RoleEnableOrDisable:
		path = routers.GroupRouterPlatformV1 + "/*"
	default:
		path = ""
	}
	return path
}

// 树结构
type TreePrivilege []struct {
	Id       int           `json:"id"`
	Label    string        `json:"label"`
	Name     Privilege     `json:"name,omitempty"`
	Children TreePrivilege `json:"children,omitempty"`
}

func Tree() (tree TreePrivilege) {
	tree = TreePrivilege{
		{
			Id:    2,
			Label: "系统管理",
			Children: TreePrivilege{
				{
					Id:    20,
					Label: "用户管理",
					Children: TreePrivilege{
						{Id: 21, Label: "访问", Name: UserAccess},
						{Id: 22, Label: "添加", Name: UserAdd},
						{Id: 23, Label: "删除", Name: UserDelete},
						{Id: 24, Label: "编辑", Name: UserEdit},
						{Id: 25, Label: "禁用/启用", Name: UserEnableOrDisable},
					},
				},
				{
					Id:    30,
					Label: "角色管理",
					Children: TreePrivilege{
						{Id: 31, Label: "访问", Name: RoleAccess},
						{Id: 32, Label: "添加", Name: RoleAdd},
						{Id: 33, Label: "删除", Name: RoleDelete},
						{Id: 34, Label: "编辑", Name: RoleEdit},
						{Id: 35, Label: "禁用/启用", Name: RoleEnableOrDisable},
						{Id: 36, Label: "权限授权", Name: RoleAuthority},
					},
				},
			},
		},
	}
	return tree
}

// 路由

// [][]string默认位数
const (
	roleDigits = 0
	pathDigits = 1
	authDigits = 2
)

// 添加角色权限
func AddRoleAuthority(role string, accessPath string, authority Privilege) {
	if ok, _ := CasRule.AddPolicy(role, accessPath, string(authority)); ok {
		log.Infof("add role authority %s path %s ok", role, accessPath)
	}
}

// 获取角色权限
func GetRoleAuthority(role string) (authority []string) {
	var authorityList = make([]string, 0)
	list := CasRule.GetPermissionsForUser(role)
	for _, vlist := range list {
		if len(vlist) >= authDigits {
			authorityList = append(authorityList, vlist[authDigits])
		}
	}
	return authorityList
}

// 角色权限继承 roleA 继承/获取 roleB权限
func InheritRoleAuthority(roleA, roleB string) {
	// 获取roleB权限
	list := CasRule.GetPermissionsForUser(roleB)
	for _, listV := range list {
		if len(listV) >= authDigits {
			CasRule.AddPolicy(roleA, listV[pathDigits], listV[authDigits])
		}
	}
}

// 角色更名时，用户继承新角色,并删除旧角色
func InheritUserRole(roleA, roleB string) {
	// 获取具有角色的用户
	res, err := CasRule.GetUsersForRole(roleB)
	if err != nil {
		log.Errorf("get users for role:%s error:%s", roleB, err.Error())
		return
	}
	for _, user := range res {
		AddUserRole(user, roleA)
		DeleteUserRole(user, roleB)
	}
}

// 删除角色权限
func DeleteRoleAuthority(role string, accessPath string, authority string) {
	if ok, _ := CasRule.RemovePolicy(role, accessPath, authority); !ok {
		fmt.Println("role authority doesn't exit!")
	} else {
		fmt.Println("role authority delete success")
	}
}

// 删除一个角色
func DeleteRole(role string) {
	CasRule.DeleteRole(role)
}

// 获取用户权限列表
func GetUserRoleList(userName string) (roleList []string) {
	roleList = make([]string, 0)
	// 获取用户角色
	res, err := CasRule.GetRolesForUser(userName)
	if err != nil {
		log.Error(err.Error())
		return roleList
	}
	for _, role := range res {
		// 获取角色权限
		list := CasRule.GetPermissionsForUser(role)
		for _, vlist := range list {
			if len(vlist) >= authDigits {
				roleList = append(roleList, vlist[authDigits])
			}
		}
	}
	return removeDuplicateElement(roleList)
}

// 为用户添加角色
func AddUserRole(userName string, role string) {
	if ok, _ := CasRule.AddRoleForUser(userName, role); !ok {
		log.Infof("role hadn't exit:%s", role)
	} else {
		log.Infof("add role success!")
	}
}

// 删除用户角色
func DeleteUserRole(userName, role string) {
	ok, _ := CasRule.DeleteRoleForUser(userName, role)
	if ok {
		log.Infof("delete role:%s for user:%s success", role, userName)
	} else {
		log.Infof("delete role:%s for user:%s failed!", role, userName)
	}
}

// 删除用户所有角色
func DeleteAllRoleForUser(userName string) {
	if ok, _ := CasRule.DeleteRolesForUser(userName); !ok {
		log.Infof("delete all user role failed:%s", userName)
	} else {
		log.Infof("delete all user role success!")
	}
}

// 删除一个用户
func DeleteUser(userName string) {
	CasRule.DeleteUser(userName)
}

// 权限列表去重
func removeDuplicateElement(authList []string) []string {
	result := make([]string, 0, len(authList))
	temp := map[string]struct{}{}
	for _, item := range authList {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
