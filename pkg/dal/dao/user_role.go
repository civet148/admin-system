package dao

import (
	"admin-system/pkg/dal/models"
	"admin-system/pkg/proto"
	"fmt"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type UserRoleDAO struct {
	db *sqlca.Engine
}

func NewUserRoleDAO(db *sqlca.Engine) *UserRoleDAO {

	return &UserRoleDAO{
		db: db,
	}
}

func (dao *UserRoleDAO) Insert(do *models.UserRoleDO) (err error) {
	if _, err = dao.db.Model(&do).Table(models.TableNameUserRole).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) Upsert(do *models.UserRoleDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).Table(models.TableNameUserRole).Select(columns...).Upsert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) SelectUserByRole(role string) (dos []*models.UserRoleDO, err error) {
	if _, err = dao.db.Model(&dos).
		Table(models.TableNameUserRole).
		Eq(models.USER_ROLE_COLUMN_ROLE_NAME, role).
		Eq(models.USER_ROLE_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) SelectUserByName(Name string) (dos []*models.UserRoleDO, err error) {
	if _, err = dao.db.Model(&dos).
		Table(models.TableNameUserRole).
		Eq(models.USER_ROLE_COLUMN_USER_NAME, Name).
		Eq(models.USER_ROLE_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) UpdateRoleNameByUser(do *models.UserRoleDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).
		Table(models.TableNameUserRole).
		Select(columns...).
		Eq(models.USER_ROLE_COLUMN_USER_NAME, do.UserName).
		Eq(models.USER_ROLE_COLUMN_DELETED, 0).
		Update(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) UpdateUserById(do *models.UserRoleDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).
		Table(models.TableNameUserRole).
		Select(columns...).
		Id(do.Id).
		Update(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) Delete(do *models.UserRoleDO) (err error) {

	if _, err = dao.db.Model(do).
		Table(models.TableNameUserRole).
		Select(
			models.USER_ROLE_COLUMN_USER_NAME,
			models.USER_ROLE_COLUMN_DELETED,
			models.USER_ROLE_COLUMN_EDIT_USER,
		).
		Eq(models.USER_ROLE_COLUMN_USER_NAME, do.UserName).
		Eq(models.USER_ROLE_COLUMN_DELETED, 0).
		Delete(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) SelectUsers(req *proto.PlatformListUserReq) (users []*proto.PlatformUser, total int64, err error) {

	users = make([]*proto.PlatformUser, 0)
	/*
			-- 查询平台用户
			SELECT
			    a.user_name, a.role_name, a.create_user,
			    b.id AS user_id, b.user_alias, b.phone_number, b.is_admin, b.password,
		        b.email, b.address, b.remark as user_remark, b.state, b.login_ip, b.login_time,  b.created_time, b.updated_time
			FROM user_role a,  USER b
			WHERE a.user_name=b.user_name AND b.deleted=0 AND b.user_name='john'
			ORDER BY b.created_time DESC
	*/
	strSelect := ` a.user_name, a.role_name, a.create_user,
	    b.id AS user_id, b.user_alias, b.phone_number, b.is_admin, b.password, 
        b.email, b.address, b.remark as user_remark, b.state, b.login_ip, b.login_time,  b.created_time, b.updated_time`

	var strWhere string
	if req.UserName != "" {
		strWhere = fmt.Sprintf(`a.user_name=b.user_name AND b.deleted=0 AND b.user_name='%s'`, req.UserName)
	} else {
		strWhere = fmt.Sprintf(`a.user_name=b.user_name AND b.deleted=0`)
	}
	strFrom := fmt.Sprintf("%s a, %s b", models.TableNameUserRole, models.TableNameUser)

	if _, total, err = dao.db.Model(&users).
		Table(strFrom).
		Select(strSelect).
		Where(strWhere).
		Page(req.PageNo, req.PageSize).
		Desc("b.created_time").
		QueryEx(); err != nil {

		log.Errorf("database query error [%s]", err.Error())
		return
	}
	return
}

func (dao *UserRoleDAO) SelectRoleUsers(strRoleName string, pageNo, pageSize int) (users []*proto.PlatformUser, total int64, err error) {

	users = make([]*proto.PlatformUser, 0)
	/*
		-- 查询平台用户
		SELECT
		    a.user_name, a.role_name,
		    b.id AS user_id, b.user_alias, b.phone_number, b.is_admin, b.email, b.address, b.remark as user_remark, b.state, b.login_ip, b.login_time,  b.created_time, b.updated_time,
		    c.role_name AS role_name, c.alias AS role_alias, c.privileges
		FROM user_role a,  USER b, ROLE c
		WHERE a.pool_id=0 AND a.cluster_id=0 AND a.user_name=b.user_name AND b.deleted=0 AND a.role_name=c.role_name AND a.role_name='platform-admin'
		ORDER BY b.created_time DESC
	*/
	strSelect := ` a.user_name, a.role_name, 
	    b.id AS user_id, b.user_alias, b.phone_number, b.is_admin, b.email, b.address, b.remark as user_remark, b.state, b.login_ip, b.login_time,  b.created_time, b.updated_time,
	    c.role_name, c.alias AS role_alias, c.privileges `

	strWhere := fmt.Sprintf(`a.user_name=b.user_name AND b.deleted=0 AND a.role_name=c.role_name AND a.role_name='%s'`, strRoleName)

	strFrom := fmt.Sprintf("%s a, %s b, %s c", models.TableNameUserRole, models.TableNameUser, models.TableNameRole)

	if _, total, err = dao.db.Model(&users).
		Table(strFrom).
		Select(strSelect).
		Where(strWhere).
		Page(pageNo, pageSize).
		Desc("b.created_time").
		QueryEx(); err != nil {

		log.Errorf("database query error [%s]", err.Error())
		return
	}

	return
}
