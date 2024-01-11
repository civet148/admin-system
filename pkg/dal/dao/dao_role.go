package dao

import (
	"admin-system/pkg/dal/models"
	"admin-system/pkg/proto"
	"admin-system/pkg/utils"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type RoleDAO struct {
	db *sqlca.Engine
}

func NewRoleDAO(db *sqlca.Engine) *RoleDAO {

	return &RoleDAO{
		db: db,
	}
}

func (dao *RoleDAO) SelectRoleById(id int32) (do *models.RoleDO, err error) {
	if _, err = dao.db.Model(&do).Table(models.TableNameRole).Where("%s='%d'", models.ROLE_COLUMN_ID, id).Query(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RoleDAO) Insert(do *models.RoleDO) (id int64, err error) {
	if id, err = dao.db.Model(&do).Table(models.TableNameRole).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RoleDAO) Upsert(do *models.RoleDO, columns ...string) (lastId int64, err error) {
	if lastId, err = dao.db.Model(do).Table(models.TableNameRole).Select(columns...).Upsert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RoleDAO) Update(do *models.RoleDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).
		Table(models.TableNameRole).
		Select(columns...).
		Exclude(models.ROLE_COLUMN_ROLE_NAME).
		Update(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RoleDAO) Delete(do *models.RoleDO) (err error) {
	var strName = do.RoleName
	do.RoleName = utils.MakeTimestampSuffix(do.RoleName)
	_, err = dao.db.Model(do).
		Table(models.TableNameRole).
		Select(
			models.ROLE_COLUMN_ROLE_NAME,
			models.ROLE_COLUMN_DELETED,
			models.ROLE_COLUMN_EDIT_USER,
		).
		Eq(models.ROLE_COLUMN_ROLE_NAME, strName).
		Eq(models.ROLE_COLUMN_IS_INHERENT, 0).
		Update()
	return
}

func (dao *RoleDAO) SelectPlatformRoles(pageNo, pageSize int, strRoleName string) (roles []*proto.PlatformRole, total int64, err error) {
	e := dao.db.Model(&roles).
		Table(models.TableNameRole).
		Eq(models.ROLE_COLUMN_POOL_ID, 0).
		Eq(models.ROLE_COLUMN_CLUSTER_ID, 0).
		Eq(models.ROLE_COLUMN_DELETED, 0).
		Page(pageNo, pageSize).
		Desc(models.ROLE_COLUMN_CREATED_TIME)

	if strRoleName != "" {
		e.And("%s='%s'", models.ROLE_COLUMN_ROLE_NAME, strRoleName)
	}
	if _, total, err = e.QueryEx(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *RoleDAO) SelectUserRole(strUserName string) (role *models.RoleDO, err error) {

	if _, err = dao.db.Model(&role).
		Table("user_role a", "`role` b").
		Select("b.id, b.role_name, b.role_alias, b.remark").
		Where("a.user_name='%v'", strUserName).
		And("a.deleted=0").
		And("a.role_name=b.role_name").
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}

	return
}

func (dao *RoleDAO) SelectRoleByName(strRoleName string) (role *models.RoleDO, err error) {
	_, err = dao.db.Model(&role).
		Table(models.TableNameRole).
		Eq(models.ROLE_COLUMN_ROLE_NAME, strRoleName).
		Eq(models.ROLE_COLUMN_DELETED, 0).
		Query()
	return
}

func (dao *RoleDAO) CheckRoleExistByName(strRoleName string) (ok bool, err error) {
	var count int64
	var do *models.RoleDO
	if count, err = dao.db.Model(&do).
		Table(models.TableNameRole).
		Eq(models.ROLE_COLUMN_ROLE_NAME, strRoleName).
		Eq(models.ROLE_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf("error [%s]", err.Error())
		return
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}
