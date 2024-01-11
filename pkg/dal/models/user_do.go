// Code generated by db2go. DO NOT EDIT.
// https://github.com/civet148/sqlca

package models

const TableNameUser = "user" //登录账户信息表 

const (
USER_COLUMN_ID = "id"
USER_COLUMN_USER_NAME = "user_name"
USER_COLUMN_USER_ALIAS = "user_alias"
USER_COLUMN_PASSWORD = "password"
USER_COLUMN_SALT = "salt"
USER_COLUMN_PHONE_NUMBER = "phone_number"
USER_COLUMN_IS_ADMIN = "is_admin"
USER_COLUMN_EMAIL = "email"
USER_COLUMN_ADDRESS = "address"
USER_COLUMN_REMARK = "remark"
USER_COLUMN_DELETED = "deleted"
USER_COLUMN_STATE = "state"
USER_COLUMN_LOGIN_IP = "login_ip"
USER_COLUMN_LOGIN_TIME = "login_time"
USER_COLUMN_CREATE_USER = "create_user"
USER_COLUMN_EDIT_USER = "edit_user"
USER_COLUMN_CREATED_TIME = "created_time"
USER_COLUMN_UPDATED_TIME = "updated_time"
)

type UserDO struct { 
	Id int32 `json:"id" db:"id" bson:"_id"` //用户ID(自增) 
	UserName string `json:"user_name" db:"user_name" bson:"user_name"` //登录名称 
	UserAlias string `json:"user_alias" db:"user_alias" bson:"user_alias"` //账户别名 
	Password string `json:"password" db:"password" bson:"password"` //登录密码(MD5+SALT) 
	Salt string `json:"salt" db:"salt" bson:"salt"` //MD5加密盐 
	PhoneNumber string `json:"phone_number" db:"phone_number" bson:"phone_number"` //联系手机号 
	IsAdmin bool `json:"is_admin" db:"is_admin" bson:"is_admin"` //是否为超级管理员(0=普通账户 1=超级管理员) 
	Email string `json:"email" db:"email" sqlca:"isnull" bson:"email"` //邮箱地址 
	Address string `json:"address" db:"address" bson:"address"` //家庭住址/公司地址 
	Remark string `json:"remark" db:"remark" bson:"remark"` //备注 
	Deleted bool `json:"deleted" db:"deleted" bson:"deleted"` //是否已删除(0=未删除 1=已删除) 
	State int8 `json:"state" db:"state" bson:"state"` //是否已冻结(1=已启用 2=已冻结) 
	LoginIp string `json:"login_ip" db:"login_ip" bson:"login_ip"` //最近登录IP 
	LoginTime int64 `json:"login_time" db:"login_time" bson:"login_time"` //最近登录时间 
	CreateUser string `json:"create_user" db:"create_user" bson:"create_user"` //创建人 
	EditUser string `json:"edit_user" db:"edit_user" bson:"edit_user"` //最近编辑人 
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间 
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间 
}

func (do *UserDO) GetId() int32 { return do.Id } 
func (do *UserDO) SetId(v int32) { do.Id = v } 
func (do *UserDO) GetUserName() string { return do.UserName } 
func (do *UserDO) SetUserName(v string) { do.UserName = v } 
func (do *UserDO) GetUserAlias() string { return do.UserAlias } 
func (do *UserDO) SetUserAlias(v string) { do.UserAlias = v } 
func (do *UserDO) GetPassword() string { return do.Password } 
func (do *UserDO) SetPassword(v string) { do.Password = v } 
func (do *UserDO) GetSalt() string { return do.Salt } 
func (do *UserDO) SetSalt(v string) { do.Salt = v } 
func (do *UserDO) GetPhoneNumber() string { return do.PhoneNumber } 
func (do *UserDO) SetPhoneNumber(v string) { do.PhoneNumber = v } 
func (do *UserDO) GetIsAdmin() bool { return do.IsAdmin } 
func (do *UserDO) SetIsAdmin(v bool) { do.IsAdmin = v } 
func (do *UserDO) GetEmail() string { return do.Email } 
func (do *UserDO) SetEmail(v string) { do.Email = v } 
func (do *UserDO) GetAddress() string { return do.Address } 
func (do *UserDO) SetAddress(v string) { do.Address = v } 
func (do *UserDO) GetRemark() string { return do.Remark } 
func (do *UserDO) SetRemark(v string) { do.Remark = v } 
func (do *UserDO) GetDeleted() bool { return do.Deleted } 
func (do *UserDO) SetDeleted(v bool) { do.Deleted = v } 
func (do *UserDO) GetState() int8 { return do.State } 
func (do *UserDO) SetState(v int8) { do.State = v } 
func (do *UserDO) GetLoginIp() string { return do.LoginIp } 
func (do *UserDO) SetLoginIp(v string) { do.LoginIp = v } 
func (do *UserDO) GetLoginTime() int64 { return do.LoginTime } 
func (do *UserDO) SetLoginTime(v int64) { do.LoginTime = v } 
func (do *UserDO) GetCreateUser() string { return do.CreateUser } 
func (do *UserDO) SetCreateUser(v string) { do.CreateUser = v } 
func (do *UserDO) GetEditUser() string { return do.EditUser } 
func (do *UserDO) SetEditUser(v string) { do.EditUser = v } 
func (do *UserDO) GetCreatedTime() string { return do.CreatedTime } 
func (do *UserDO) SetCreatedTime(v string) { do.CreatedTime = v } 
func (do *UserDO) GetUpdatedTime() string { return do.UpdatedTime } 
func (do *UserDO) SetUpdatedTime(v string) { do.UpdatedTime = v } 
/*
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID(自增)',
  `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录名称',
  `user_alias` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账户别名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录密码(MD5+SALT)',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'MD5加密盐',
  `phone_number` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系手机号',
  `is_admin` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否为超级管理员(0=普通账户 1=超级管理员)',
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '邮箱地址',
  `address` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '家庭住址/公司地址',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已删除(0=未删除 1=已删除)',
  `state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已冻结(1=已启用 2=已冻结)',
  `login_ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '最近登录IP',
  `login_time` bigint NOT NULL DEFAULT '0' COMMENT '最近登录时间',
  `create_user` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `edit_user` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '最近编辑人',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `UNIQ_USER_NAME` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='登录账户信息表';
*/