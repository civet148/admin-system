// Code generated by db2go. DO NOT EDIT.
// https://github.com/civet148/sqlca

package models

const TableNameLogin = "login" //登录记录表

const (
	LOGIN_COLUMN_ID           = "id"
	LOGIN_COLUMN_USER_ID      = "user_id"
	LOGIN_COLUMN_LOGIN_IP     = "login_ip"
	LOGIN_COLUMN_LOGIN_ADDR   = "login_addr"
	LOGIN_COLUMN_CREATED_TIME = "created_time"
	LOGIN_COLUMN_UPDATED_TIME = "updated_time"
)

type LoginDO struct {
	Id          int32  `json:"id" db:"id" bson:"_id"`                                               //自增ID
	UserId      int32  `json:"user_id" db:"user_id" bson:"user_id"`                                 //登录用户ID
	LoginIp     string `json:"login_ip" db:"login_ip" bson:"login_ip"`                              //登录IP
	LoginAddr   string `json:"login_addr" db:"login_addr" bson:"login_addr"`                        //登录地址
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间
}

func (do *LoginDO) GetId() int32            { return do.Id }
func (do *LoginDO) SetId(v int32)           { do.Id = v }
func (do *LoginDO) GetUserId() int32        { return do.UserId }
func (do *LoginDO) SetUserId(v int32)       { do.UserId = v }
func (do *LoginDO) GetLoginIp() string      { return do.LoginIp }
func (do *LoginDO) SetLoginIp(v string)     { do.LoginIp = v }
func (do *LoginDO) GetLoginAddr() string    { return do.LoginAddr }
func (do *LoginDO) SetLoginAddr(v string)   { do.LoginAddr = v }
func (do *LoginDO) GetCreatedTime() string  { return do.CreatedTime }
func (do *LoginDO) SetCreatedTime(v string) { do.CreatedTime = v }
func (do *LoginDO) GetUpdatedTime() string  { return do.UpdatedTime }
func (do *LoginDO) SetUpdatedTime(v string) { do.UpdatedTime = v }

/*
CREATE TABLE `login` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` int NOT NULL COMMENT '登录用户ID',
  `login_ip` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录IP',
  `login_addr` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '登录地址',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='登录记录表';
*/
