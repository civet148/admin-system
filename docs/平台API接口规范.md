# 项目管理系统接口说明



# 1. 错误码定义



| 错误码 | 错误信息               | 备注     |
| ------ | ---------------------- | -------- |
| -1     | CODE_ERROR        | 未知错误 |
| 0      | CODE_OK           | 成功     |
| 500    | CODE_INTERNAL_SERVER_ERROR | 内部服务器错误 |
| 1000   | CODE_ACCESS_DENY  | 无访问权限 |
| 1001   | CODE_UNAUTHORIZED | TOKEN鉴权失败 |
| 1002   | CODE_INVALID_USER_OR_PASSWORD  | 无效用户名或密码 |
| 1003   | CODE_INVALID_PARAMS | 无效参数 |
| 1004   | CODE_INVALID_JSON_OR_REQUIRED_PARAMS | 请求数据JSON格式不正确或缺少必填参数|
| 1005   | CODE_ALREADY_EXIST | 数据已存在(用户名/邮箱/权限名称等等) |
| 1006   | CODE_NOT_FOUND | 没有找到数据 |
| 1007   | CODE_INVALID_PASSWORD | 无效密码 |
| 1008   | CODE_INVALID_AUTH_CODE | 无效验证码或验证码超时 |
| 1009   | CODE_ACCESS_VIOLATE | 访问违规 |
| 1010 | CODE_TYPE_UNDEFINED | 类型未定义 |
| 1011 | CODE_BAD_DID_OR_SIGNATURE | DID/签名错误 |
| 1012 | CODE_ACCOUNT_BANNED | 账户被禁用 |
| 1013 | CODE_EXPORT_FAILED | 导出失败 |
| 1014 | CODE_PRIVILEGE_EXIT | 用户已有权限 |
| 1015 | CODE_USERNAME_EXIT | 用户名已存在 |
| 1016 | CODE_PHONENUMBER_EXIT | 手机号已存在 |
| 1017 | CODE_PHONE_FORMAT_WRONG | 手机格式错误 |
| 1018 | CODE_PIVILEGE_EXIT | 权限名称已存在 |
| 1019 | CODE_SERVER_ERROR | 服务器错误 |
| 1020 | CODE_SESSION_CONTEXT | 用户一退出 |
| 1021 | CODE_NO_PRIVILEGE | 无权限 |
| 1022 | CODE_GET_PRIVILEGE_FAILED | 获取权限失败 |
| 1023 | CODE_CHAIN_VERFICATION_FAILED | 检验失败 |
| 1024 | CODE_EDIT_WINDING_CYCLE_FAILED | 编辑上链周期失败 |
| 1025 | CODE_EDIT_CONTROL_FAILED | 开启/禁用失败 |
| 1026 | CODE_MAX_DURATION_LIMIT | 不在周期范围内 |
| 1027 | CODE_EMAIL_FORMAT_WRONG | 邮箱格式错误 |
| 1028 | CODE_EMAIL_EXIT | 邮箱已存在1029 |
| 1029 | CODE_EMAIL_SEND_FAILED | 邮件发送失败 |
| 1030 | CODE_EMAIL_CODE_TIMEOUT | 验证码超时！ |
| 1031 | CODE_EMAIL_NOT_EXIT | 邮箱不存在 |
| 1032 | CODE_EMAIL_CODE_ERROR | 验证码错误！ |



# 2. 权限码定义

| 权限码                       | 取值                            | 业务码用途              |
| ---------------------------- | ------------------------------ | ---------------------- |
| Null                         | "Null"                         | 无权限校验              |
| UserAccess                   |   "UserAccess"                 | 系统管理-用户管理-访问 |
| UserAdd                      |   "UserAdd"                    | 系统管理-用户管理-添加 |
| UserEdit                     |   "UserEdit"                   | 系统管理-用户管理-编辑 |
| UserDelete                   |   "UserDelete"                 | 系统管理-用户管理-删除 |
| UserEnableOrDisable          |   "UserEnableOrDisable"        | 系统管理-用户管理-禁用/启用 |
| RoleAccess                   |   "RoleAccess"                 | 系统管理-角色管理-访问 |
| RoleAdd                      |   "RoleAdd"                    | 系统管理-角色管理-添加 |
| RoleDelete                   |   "RoleDelete"                 | 系统管理-角色管理-删除 |
| RoleEdit                     |   "RoleEdit"                   | 系统管理-角色管理-编辑 |
| RoleEnableOrDisable          |   "RoleEnableOrDisable"        | 系统管理-角色管理-禁用/启用 |
| RoleAuthority | "RoleAuthority" | 系统管理-角色管理-权限授权 |

# 3. 系统通用接口

## 3.1 平台登录

- **请求方式**   POST   
- **数据格式**   application/json
- **请求路径**   http://127.0.0.1:8088/api/v1/login
- **请求头部** 
- **请求数据** 

|  字段    |  类型    |   描述   |  必填    |   备注   |
| ---- | ---- | ---- | ---- | ---- |
| name | string | 登录用户名 | YES |  |
| password | string | 登录密码 | YES | 输入密码后前段直接进行md5加密 |

**请求示例:**

```json
{
    "name":"admin",
    "password":"admin" // MD5加密后的数据
}
```

- **响应数据**

```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {
        "user_name": "libin",
        "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MjE1NzgxODIsImNsYWltX2lhdCI6MTYyMDk3MzM4MiwiY2xhaW1faWQiOjEzLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibGliaW4ifQ.haUXz_hihQebreN54aulsxEadY7oXOEbcCVNLb23ujY",
        "role":"admin",
        "privilege":["Null","UserAccess"]
    }
}
```



## 3.2 平台退出

**请求方式** POST

- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/logout
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain



- **请求数据** 

| 字段 | 类型 | 描述 | 必填 | 备注 |
| ---- | ---- | ---- | ---- | ---- |
|      |      |      |      |      |


**请求示例:**

```json
{
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```

## 3.3  系统账户列表



- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/list/user
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段      | 类型   | 描述     | 必填 | 备注              |      |
| --------- | ------ | -------- | ---- | ----------------- | ---- |
| name      | string | 用户名   | NO   |                   |      |
| page_no   | int    | 分页页码 | NO   | 分页时从0开始递增 |      |
| page_size | int    | 分页条数 | NO   | 不分页填0即可     |      |

**请求示例:**

```json
{
    "page_no":0,
    "page_size": 50
}
```

- **响应数据**

```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 2,
        "count": 2
    },
    "data": {
        "users": [
            {
                "user_id": 11,
                "user_did":"65+4121212",//数字身份ID
                "user_name": "lory", // 用户名
                "user_alias": "李彬", // 真实姓名
                "phone_number": "137000000", // 电话号码, 
                "create_user": "admin", // 创建者
                "login_time": 1618798847, // 最近登陆时间
                "role_name": "platform-admin",// 角色名
                "remark": "平台超级管理员",// 描述
                "password":"454ad",// md5密码
            },
            {
                "user_id": 11,
                "user_did":"65+4121212",//数字身份ID
                "user_name": "lory", // 用户名
                "user_alias": "李彬", // 真实姓名
                "phone_number": "137000000", // 电话号码, 
                "create_user": "admin", // 创建者
                "login_time": 1618798847, // 最近登陆时间
                "role_name": "platform-admin",// 角色名
                "remark": "平台超级管理员",// 描述
                "password":"454ad",// md5密码
            }
        ]
    }
}
```
## 3.4 创建系统账户



- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/create/user
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段      | 类型   | 描述     | 必填 | 备注 |
| --------- | ------ | -------- | -- | ---- |
| user_name | string | 账户名   | YES |      |
| user_alias     | string | 真实姓名 | YES |      |
| phone_number| string | 电话号码 | YES |      |
| password  | string | 密码 （前端直接用md5加密） | YES |      |
| remark    | string | 描述     | NO |      |

**请求示例:**

```json
{
  "user_name": "lory",
  "user_alias": "lory.Lee",
  "phone_number": "18682371690",
  "password": "e10adc3949ba59abbe56e057f20f883e",
  "remark": "普通用戶"
}
```

- **响应数据**

```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {
        "user_id": 11
    }
}
```

## 3.5 编辑系统账户信息

- **请求方式** POST

- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/edit/user
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段      | 类型   | 描述     | 必填 | 备注 |
| --------- | ------ | -------- | ---- | ---- |
| user_name | string | 账户名   | YES  |      |
| user_alias     | string | 真实姓名 | NO   |      |
| phone_number| string | 电话号码 | NO  |      |
| password  | string | 密码 （前端直接用md5加密） | NO  |      |
| remark    | string | 描述     | NO   |      |


**请求示例:**

```json
{
	"user_name": "lory",
	"user_alias": "lory.Lee",
	"phone_number": "1370000009124",
	"password": "2412423sfdasdfsdf",
	"remark": "普通用戶"
}
```

- **响应数据**

```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.6 删除系统账户

- **请求方式** POST
- **数据格式**   application/json
- **请求路径** http://127.0.0.1:8088/api/v1/platform/delete/user
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段      | 类型   | 描述   | 必填 | 备注 |
| --------- | ------ | ------ | ---- | ---- |
| user_name | string | 账户名 | YES  |      |


**请求示例:**

```json
{
	"user_name": "lory2"
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.7 系统角色列表

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/list/role
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain


| 字段      | 类型 | 描述     | 必填 | 备注              |      |
| --------- | ---- | -------- | ---- | ----------------- | ---- |
| page_no   | int  | 分页页码 | NO   | 分页时从0开始递增 |      |
| page_size | int  | 分页条数 | NO   | 不分页填0即可     |      |

**请求示例:**

```json
{
    "page_no":0,
    "page_size": 50
}
```

- **响应数据**

```json
{
  "header": {
    "code": 0,
    "message": "",
    "total": 1,
    "count": 1
  },
  "data": {
    "roles": [
      {
        "id": 1,
        "role_name": "admin",
        "create_user": "admin",
        "remark": "supper administrator role",
        "created_time": "2024-01-11 15:33:06",
        "role": [
          "UserAccess",
          "UserAdd",
          "UserEdit",
          "UserDelete",
          "UserEnableOrDisable",
          "RoleAccess",
          "RoleAdd",
          "RoleDelete",
          "RoleEdit",
          "RoleAuthority",
          "RoleEnableOrDisable"
        ]
      }
    ]
  }
```
## 3.8 创建系统角色

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/create/role
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段             | 类型   | 描述         | 必填 | 备注 |
| ---------------- | ------ | ------------ | ---- | ---- |
| role_name             | string | 平台角色名称 | YES  |      |
| remark           | string | 备注         | NO   |      |

**请求示例:**

```json
{
    "role_name":"platform-roles-manager",
    "remark":"自定义角色（非平台固有角色）"
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.9 编辑系统角色



- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/edit/role
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段   | 类型   | 描述         | 必填 | 备注 |
| ------ | ------ | ------------ | ---- | ---- |
| id     | int    | 角色ID       | YES  |      |
| role_name   | string | 平台角色名称 | YES  |      |
| remark | string | 备注         | NO   |      |

**请求示例:**

```json
{
    "id":1,
    "role_name":"platform-roles-manager",
    "remark":"自定义角色（非平台固有角色,可删除）",
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.10 删除系统角色

- **请求方式** POST
- **数据格式**   application/json
- **请求路径** http://127.0.0.1:8088/api/v1/platform/delete/role
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段 | 类型   | 描述         | 必填 | 备注 |
| ---- | ------ | ------------ | ---- | ---- |
| role_name | string | 平台角色名称 | YES  |      |

**请求示例:**

```json
{
    "role_name":"platform-roles-manager",
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.11 系统角色授权

- **请求方式** POST
- **数据格式**   application/json
- **请求路径** http://127.0.0.1:8088/api/v1/platform/auth/role
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段     | 类型   | 描述         | 必填 | 备注                        |
| -------- | ------ | ------------ | ---- | --------------------------- |
| role_name     | string | 平台角色名称 | YES  |                             |
| privilege |[]string|权限列表|YES||


**请求示例:**

```json
{
    "role_name":"platform-roles-manager",
    "privilege":["RoleAuthority","RoleEdit"]
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```
## 3.12 系统角色/用户权限查询

- **请求方式** POST
- **数据格式**   application/json
- **请求路径** http://127.0.0.1:8088/api/v1/platform/inquire/auth
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段     | 类型   | 描述         | 必填 | 备注                        |
| -------- | ------ | ------------ | ---- | --------------------------- |
| name     | string | 名称 | YES  |                             |
| name_type | int | 类型（0：用户，1：角色） | YES  |                             |


**请求示例:**

```json
{
    "name":"platform-roles-manager",
    "name_type":0
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {
    "privilege":["PlatformUserAccess","PlatformUserAdd","PlatformUserDisabled"]
    }
}
```


## 3.13 用户修改密码

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/change/password
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段         | 类型   | 描述 | 必填 | 备注 |
| ------------ | ------ | ---- | ---- | ---- |
| old_password | string |      |      |      |
| new_password | string |      |      |      |

**请求示例:**

```json
{
	"old_password": "123456",
	"new_password": "admin"
}
```

- **响应数据**

```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```

## 3.14 管理员重置用户密码



- **请求方式** POST
- **数据格式**   application/json
- **请求路径** http://127.0.0.1:8088/api/v1/platform/reset/password
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

| 字段         | 类型   | 描述   | 必填 | 备注 |
| ------------ | ------ | ------ | ---- | ---- |
| user_name    | string | 账户名 | YES  |      |
| new_password | string | 新密码 | YES  |      |


**请求示例:**

```json
{
    "user_name":"admin",
	"new_password": "admin"
}
```

- **响应数据**

```json
{
    "header":{
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {}
}
```


## 3.15 获取角色用户数

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/list/role-user
- **请求头部**  鉴权信息

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据** 

|  字段    |  类型    |   描述   |  必填    |   备注   |
| ---- | ---- | ---- | ---- | ---- |
| role_name | 角色名 |      | YES |      |


**请求示例:**

```json
{
    "role_name":"platform-admin"
}
```

- **响应数据**

|  字段    |  类型    |   描述   |   备注   |
| ---- | ---- | ---- | ---- |
| data | JSON对象 | 返回数据结构体 |      |
| \|-   role_name | string | 角色名 |      |
| \|-   user_count | int | 拥有该角色的用户数 |      |


```json
{
    "header": {
        "code": 0,
        "message": "CODE_OK",
        "total": 1,
        "count": 1
    },
    "data": {
        "role_name": "platform-admin",
        "user_count": 1
    }
}
```


## 3.16 刷新token

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/refresh/token

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据**

| 字段       | 类型   | 描述   | 必填 | 备注 |
| ---------- | ------ | ------ | ---- | ---- |
|     |    |  |   |      |



**请求示例:**

```json
{
}
```

- **响应数据**

| 字段       | 类型   | 描述         | 备注          |
| ---------- | ------ | ------------ | ------------- |
| auth_token | string | token |               |


```json
{
  "header": {
    "code": 0,
    "message": "CODE_OK",
    "total": 1,
    "count": 1
  },
  "data": {
    "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MzI0MjI1NDgsImNsYWltX2lhdCI6MTYzMjM3OTM0OCwiY2xhaW1faWQiOjAsImNsYWltX2lzX2FkbWluIjpmYWxzZSwiY2xhaW1fdXNlcm5hbWUiOiIifQ.op3xdTUt7sZqXJ-KkWDl5kHNHQm8zye3sNTZIV7_GJw"
  }
}
```

## 3.17 操作日志

- **请求方式** POST
- **数据格式**   application/json
- **请求路径**  http://127.0.0.1:8088/api/v1/platform/list/oper-log

**Request Headers**

**Auth-Token:** eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGFpbV9leHBpcmUiOjE2MTk0MDM2NDcsImNsYWltX2lhdCI6MTYxODc5ODg0NywiY2xhaW1faWQiOjExLCJjbGFpbV9pc19hZG1pbiI6ZmFsc2UsImNsYWltX3VzZXJuYW1lIjoibG9yeSJ9.Y_0irxrsYGpPRvaO_8Od_jgqFbcGPvFM4MRefCHEdzA

**Content-Type:** text/plain

- **请求数据**

| 字段       | 类型   | 描述   | 必填 | 备注 |
| ---------- | ------ | ------ | ---- | ---- |
| page_no | int | 分页页码 | NO | 分页时从0开始递增 |
| page_size | int | 分页条数 | NO | 不分页填0即可 |
| oper_user | string | 操作人名称 | NO | 搜索条件 |
| oper_type | int | 操作类型 | NO | 搜索条件: 0=全部 1=首页 2=系统管理 3=项目管理 4=资源管理 5=告警中心 |
| start_time | string | 开始时间 | NO | 搜索条件 |
| end_time | string | 结束时间 | NO | 搜索条件 |



**请求示例:**

```json
{
    "page_no":0,
    "page_size": 50
}
```

- **响应数据**

| 字段       | 类型   | 描述         | 备注          |
| ---------- | ------ | ------------ | ------------- |
|  |  |  |               |


```json
{
  "header": {
    "code": 0,
    "message": "CODE_OK",
    "total": 1,
    "count": 1
  },
  "data": {
    "list":[
        {
            "oper_user":"libin", //操作人
            "oper_type": 1, //操作类型(1=首页 2=...)
            "oper_time":"2021-10-11 03:22:33", //操作时间
            "oper_content":"添加zzz用户", //操作内容描述
        }
    ]
  }
}
```
