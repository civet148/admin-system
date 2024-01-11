package dao

import (
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"admin-system/pkg/dal/models"
)

type LoginDAO struct {
	db *sqlca.Engine
}

func NewLoginDAO(db *sqlca.Engine) *LoginDAO {

	return &LoginDAO{
		db: db,
	}
}

func (dao *LoginDAO) Insert(dos ...*models.LoginDO) (lastId int64, err error) {
	if lastId, err = dao.db.Model(&dos).Table(models.TableNameLogin).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}
