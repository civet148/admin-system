package dao

import (
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"admin-system/pkg/dal/models"
)

type PrivilegeDAO struct {
	db *sqlca.Engine
}

func NewPrivilegeDAO(db *sqlca.Engine) *PrivilegeDAO {

	return &PrivilegeDAO{
		db: db,
	}
}

func (dao *PrivilegeDAO) Insert(do *models.PrivilegeDO) (id int64, err error) {
	if id, err = dao.db.Model(&do).Table(models.TableNamePrivilege).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *PrivilegeDAO) Upsert(do *models.PrivilegeDO, columns ...string) (lastId int64, err error) {
	if lastId, err = dao.db.Model(do).Table(models.TableNamePrivilege).Select(columns...).Upsert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}
