package proto

type PlatformUser struct {
	UserId      int32  `json:"user_id" db:"user_id" bson:"user_id"`                                 //用户ID(自增)
	UserName    string `json:"user_name" db:"user_name" bson:"user_name"`                           //登录名称
	UserAlias   string `json:"user_alias" db:"user_alias" bson:"user_alias"`                        //真实姓名
	PhoneNumber string `json:"phone_number" db:"phone_number" bson:"phone_number"`                  //联系手机号
	IsAdmin     bool   `json:"is_admin" db:"is_admin" bson:"is_admin"`                              //是否为超级管理员(0=普通账户 1=超级管理员)
	Email       string `json:"email" db:"email" bson:"email"`                                       //邮箱地址
	Address     string `json:"address" db:"address" bson:"address"`                                 //家庭住址/公司地址
	UserRemark  string `json:"user_remark" db:"user_remark" bson:"user_remark"`                     //备注
	State       int    `json:"state" db:"state" bson:"state"`                                       //是否已冻结(1=正常 2=已冻结)
	LoginIp     string `json:"login_ip" db:"login_ip" bson:"login_ip"`                              //最近登录IP
	LoginTime   int64  `json:"login_time" db:"login_time" bson:"login_time"`                        //最近登录时间
	RoleName    string `json:"role_name" db:"role_name" bson:"role_name"`                           //角色名称
	RoleAlias   string `json:"role_alias" db:"role_alias" bson:"role_alias"`                        //角色别名
	CreateUser  string `json:"create_user" db:"create_user" bson:"create_user"`                     //创建人
	EditUser    string `json:"edit_user" db:"edit_user" bson:"edit_user"`                           //最近编辑人
	Password    string `json:"password" db:"password" bson:"password"`                              //密码
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间
}

type PlatformRole struct {
	Id          int32  `json:"id" db:"id" bson:"id"`                                                //角色ID(自增)
	RoleName    string `json:"role_name" db:"role_name" bson:"role_name"`                           //角色名称
	RoleAlias   string `json:"role_alias" db:"alias" bson:"alias"`                                  //角色别名
	CreateUser  string `json:"create_user" db:"create_user" bson:"create_user"`                     //创建人
	EditUser    string `json:"edit_user" db:"edit_user" bson:"edit_user"`                           //最近编辑人
	Remark      string `json:"remark" db:"remark" bson:"remark"`                                    //备注
	IsInherent  bool   `json:"is_inherent" db:"is_inherent" bson:"is_inherent"`                     //是否固有角色(false=自定义角色 true=平台固有角色)
	Deleted     bool   `json:"deleted" db:"deleted" bson:"deleted"`                                 //是否已删除(false=未删除 true=已删除)
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间
}

type PlatformTotalUser struct {
	UserId      int32  `json:"user_id" db:"user_id" bson:"user_id"`                                 //用户ID(自增)
	UserName    string `json:"user_name" db:"user_name" bson:"user_name"`                           //登录名称
	UserAlias   string `json:"user_alias" db:"user_alias" bson:"user_alias"`                        //真实姓名
	PhoneNumber string `json:"phone_number" db:"phone_number" bson:"phone_number"`                  //联系手机号
	Email       string `json:"email" db:"email" bson:"email"`                                       //邮箱
	Remark      string `json:"remark" db:"remark" bson:"remark"`                                    //备注
	Password    string `json:"password" db:"password" bson:"password"`                              // 密码
	LoginTime   int64  `json:"login_time" db:"login_time" bson:"login_time"`                        //最近登录时间
	RoleName    string `json:"role_name" db:"role_name" bson:"role_name"`                           //角色名称
	State       int    `json:"state" db:"state" bson:"state"`                                       //是否已冻结(1=正常 2=已冻结)
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间
	CreateUser  string `json:"create_user" db:"create_user" bson:"create_user"`                     //创建人
}

type PlatformSysRole struct {
	Id          int32    `json:"id" db:"id" bson:"id"`                                                //角色ID(自增)
	RoleName    string   `json:"role_name" db:"role_name" bson:"role_name"`                           //角色名称
	CreateUser  string   `json:"create_user" db:"create_user" bson:"create_user"`                     //创建人
	Remark      string   `json:"remark" db:"remark" bson:"remark"`                                    //备注
	CreatedTime string   `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间
	Role        []string `json:"role"`                                                                //角色权限
}
