package dao

import (
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"admin-system/pkg/dal/models"
)

type OperLogDAO struct {
	db *sqlca.Engine
}

func NewOperLogDAO(db *sqlca.Engine) *OperLogDAO {

	return &OperLogDAO{
		db: db,
	}
}

func (dao *OperLogDAO) Insert(dos ...*models.OperLogDO) (lastId int64, err error) {
	if lastId, err = dao.db.Model(&dos).Table(models.TableNameOperLog).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}
