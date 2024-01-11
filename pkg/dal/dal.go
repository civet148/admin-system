package dal

import (
	"admin-system/pkg/config"
	"admin-system/pkg/dal/core"
	"fmt"
	"github.com/civet148/sqlca/v2"
)

// data access layer
type Dal struct {
	PlatformCore *core.PlatformCore
}

var dal *Dal

func GetDal(cfg *config.Config) *Dal {
	if dal == nil {
		dal = newDal(cfg)
	}
	return dal
}

func newDal(cfg *config.Config) *Dal {

	db, err := sqlca.NewEngine(cfg.DSN)
	if err != nil {
		panic(fmt.Sprintf("database error, please check your DSN [%s] error [%s]", cfg.DSN, err.Error()))
		return nil
	}
	if cfg.Debug {
		db.SlowQuery(true, 500) //print database query elapse time
	}

	return &Dal{
		PlatformCore: core.NewPlatformCore(cfg, db),
	}
}
