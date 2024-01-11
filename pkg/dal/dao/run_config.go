package dao

import (
	"admin-system/pkg/dal/models"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type RunConfigDAO struct {
	db *sqlca.Engine
}

func NewRunConfigDAO(db *sqlca.Engine) *RunConfigDAO {

	return &RunConfigDAO{
		db: db,
	}
}

func (dao *RunConfigDAO) Insert(dos ...*models.RunConfigDO) (lastId int64, err error) {
	if lastId, err = dao.db.Model(&dos).
		Table(models.TableNameRunConfig).
		Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RunConfigDAO) Update(do *models.RunConfigDO, columns ...string) (lastId int64, err error) {
	if lastId, err = dao.db.Model(&do).
		Select(columns...).
		Table(models.TableNameRunConfig).
		Update(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RunConfigDAO) UpdateByConfigKey(do *models.RunConfigDO, columns ...string) (lastId int64, err error) {
	if lastId, err = dao.db.Model(&do).
		Select(columns...).
		Table(models.TableNameRunConfig).
		Eq(models.RUN_CONFIG_COLUMN_CONFIG_NAME, do.ConfigName).
		Eq(models.RUN_CONFIG_COLUMN_CONFIG_KEY, do.ConfigKey).
		Update(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RunConfigDAO) SelectIntValueByConfigKey(strConfigName, strConfigKey string) (value int, err error) {
	if _, err = dao.db.Model(&value).
		Select(models.RUN_CONFIG_COLUMN_CONFIG_VALUE).
		Table(models.TableNameRunConfig).
		Eq(models.RUN_CONFIG_COLUMN_CONFIG_NAME, strConfigName).
		Eq(models.RUN_CONFIG_COLUMN_CONFIG_KEY, strConfigKey).
		Query(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}
