// Code generated by db2go. DO NOT EDIT.
// https://github.com/civet148/sqlca

package models

const TableNameUserRole = "user_role" //用户角色关系表 

const (
USER_ROLE_COLUMN_ID = "id"
USER_ROLE_COLUMN_USER_NAME = "user_name"
USER_ROLE_COLUMN_ROLE_NAME = "role_name"
USER_ROLE_COLUMN_CREATE_USER = "create_user"
USER_ROLE_COLUMN_EDIT_USER = "edit_user"
USER_ROLE_COLUMN_DELETED = "deleted"
USER_ROLE_COLUMN_CREATED_TIME = "created_time"
USER_ROLE_COLUMN_UPDATED_TIME = "updated_time"
)

type UserRoleDO struct { 
	Id int32 `json:"id" db:"id" bson:"_id"` //自增ID 
	UserName string `json:"user_name" db:"user_name" bson:"user_name"` //用户名 
	RoleName string `json:"role_name" db:"role_name" bson:"role_name"` //角色名 
	CreateUser string `json:"create_user" db:"create_user" bson:"create_user"` //创建人 
	EditUser string `json:"edit_user" db:"edit_user" bson:"edit_user"` //最近编辑人 
	Deleted bool `json:"deleted" db:"deleted" bson:"deleted"` //是否已删除(0=未删除 1=已删除) 
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间 
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间 
}

func (do *UserRoleDO) GetId() int32 { return do.Id } 
func (do *UserRoleDO) SetId(v int32) { do.Id = v } 
func (do *UserRoleDO) GetUserName() string { return do.UserName } 
func (do *UserRoleDO) SetUserName(v string) { do.UserName = v } 
func (do *UserRoleDO) GetRoleName() string { return do.RoleName } 
func (do *UserRoleDO) SetRoleName(v string) { do.RoleName = v } 
func (do *UserRoleDO) GetCreateUser() string { return do.CreateUser } 
func (do *UserRoleDO) SetCreateUser(v string) { do.CreateUser = v } 
func (do *UserRoleDO) GetEditUser() string { return do.EditUser } 
func (do *UserRoleDO) SetEditUser(v string) { do.EditUser = v } 
func (do *UserRoleDO) GetDeleted() bool { return do.Deleted } 
func (do *UserRoleDO) SetDeleted(v bool) { do.Deleted = v } 
func (do *UserRoleDO) GetCreatedTime() string { return do.CreatedTime } 
func (do *UserRoleDO) SetCreatedTime(v string) { do.CreatedTime = v } 
func (do *UserRoleDO) GetUpdatedTime() string { return do.UpdatedTime } 
func (do *UserRoleDO) SetUpdatedTime(v string) { do.UpdatedTime = v } 
/*
CREATE TABLE `user_role` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `role_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名',
  `create_user` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `edit_user` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '最近编辑人',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已删除(0=未删除 1=已删除)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UNIQ_USER_NAME` (`user_name`) COMMENT '用户、角色类型唯一约束'
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='用户角色关系表';
*/
