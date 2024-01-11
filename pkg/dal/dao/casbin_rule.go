package dao

import (
    "fmt"
	"github.com/civet148/sqlca/v2"
	"admin-system/pkg/dal/models"
)


type CasbinRuleDAO struct {
	db *sqlca.Engine
}


func NewCasbinRuleDAO(db *sqlca.Engine) *CasbinRuleDAO {
	return &CasbinRuleDAO{
		db: db,
	}
}


//insert into table by data model
func (dao *CasbinRuleDAO) Insert(do *models.CasbinRuleDO) (lastInsertId int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameCasbinRule).Insert()
}


//insert if not exist or update columns on duplicate key...
func (dao *CasbinRuleDAO) Upsert(do *models.CasbinRuleDO, columns...string) (lastInsertId int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameCasbinRule).Select(columns...).Upsert()
}


//update table set columns where id=xxx
func (dao *CasbinRuleDAO) Update(do *models.CasbinRuleDO, columns...string) (rows int64, err error) {
	return dao.db.Model(&do).Table(models.TableNameCasbinRule).Select(columns...).Update()
}


//query records by id
func (dao *CasbinRuleDAO) QueryById(id interface{}, columns...string) (do *models.CasbinRuleDO, err error) {
	if _, err = dao.db.Model(&do).Table(models.TableNameCasbinRule).Id(id).Select(columns...).Query(); err != nil {
		return nil, err
	}
	return
}


//query records by conditions
func (dao *CasbinRuleDAO) QueryByCondition(conditions map[string]interface{}, columns...string) (dos []*models.CasbinRuleDO, err error) {
    if len(conditions) == 0 {
        return nil, fmt.Errorf("condition must not be empty")
    }
    e := dao.db.Model(&dos).Table(models.TableNameCasbinRule).Select(columns...)
    for k, v := range conditions {
        e.Eq(k, v)
    }
	if _, err = e.Query(); err != nil {
		return nil, err
	}
	return
}

