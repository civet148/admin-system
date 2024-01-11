package privilege

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/civet148/log"
	_ "github.com/go-sql-driver/mysql"
)

var CasRule *casbin.Enforcer

func InitCasbin(strDSN string) {
	// 要使用自己定义的数据库oss-manager,最后的true很重要.默认为false,使用缺省的数据库名casbin,不存在则创建，表不需要自己创建默认为casbin_rule
	a, err := xormadapter.NewAdapter("mysql", strDSN, true)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && r.act == p.act || r.sub == "root"
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	CasRule, err = casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	//从DB加载策略
	CasRule.LoadPolicy()
}
