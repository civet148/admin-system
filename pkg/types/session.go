package types

import (
	"os"
	"path/filepath"
)

var (
	DefaultStaticHome  = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/static"))
	DefaultDocsHome    = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/docs"))
	DefaultImagesHome  = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/images"))
	DefaultConfigHome  = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/config"))
	DefaultLevelDBHome = os.ExpandEnv(filepath.Join("$HOME", ".admin-system/db"))
)

type Context struct {
	*Session
}

type Session struct {
	UserId      int32  `json:"user_id" db:"user_id"`
	UserName    string `json:"user_name" db:"user_name"`
	Alias       string `json:"alias" db:"alias"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	IsAdmin     bool   `json:"is_admin" db:"is_admin"`
	LoginIP     string `json:"login_ip" db:"login_ip"`
	AuthToken   string `json:"auth_token" db:"auth_token"`
}

func (ctx *Context) UserId() int32 {
	return ctx.Session.UserId
}

func (ctx *Context) AuthToken() string {
	return ctx.Session.AuthToken
}

func (ctx *Context) UserName() string {
	return ctx.Session.UserName
}

func (ctx *Context) Alias() string {
	return ctx.Session.Alias
}

func (ctx *Context) PhoneNumber() string {
	return ctx.Session.PhoneNumber
}

func (ctx *Context) LoginIP() string {
	return ctx.Session.LoginIP
}

func (ctx *Context) IsAdmin() bool {
	return ctx.Session.IsAdmin
}

func (ctx *Context) GetUserId() int32 {
	return ctx.Session.UserId
}
