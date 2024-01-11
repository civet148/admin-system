package dao

import (
	"admin-system/pkg/dal/models"
	"admin-system/pkg/proto"
	"admin-system/pkg/utils"

	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

const (
	UserState_Enabled  = 1
	UserState_Disabled = 2
)

type UserDAO struct {
	db *sqlca.Engine
}

func NewUserDAO(db *sqlca.Engine) *UserDAO {

	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(do *models.UserDO) (id int64, err error) {
	if id, err = dao.db.Model(&do).Table(models.TableNameUser).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) Upsert(do *models.UserDO, columns ...string) (lastId int64, err error) {
	if lastId, err = dao.db.Model(do).Table(models.TableNameUser).Select(columns...).Upsert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) UpdateByName(do *models.UserDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).
		Table(models.TableNameUser).
		Select(columns...).
		Exclude(models.USER_COLUMN_USER_NAME, models.USER_COLUMN_IS_ADMIN).
		Eq(models.USER_COLUMN_USER_NAME, do.UserName).
		Update(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) DeleteUser(do *models.UserDO) (err error) {
	var strUserName = do.UserName
	do.UserName = utils.MakeTimestampSuffix(do.UserName)
	do.Email = utils.MakeTimestampSuffix(do.Email)
	_, err = dao.db.Model(do).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_USER_NAME,
			models.USER_COLUMN_EMAIL,
			models.USER_COLUMN_DELETED,
			models.USER_COLUMN_EDIT_USER,
		).
		Eq(models.USER_COLUMN_USER_NAME, strUserName).
		Update()
	return
}

func (dao *UserDAO) CheckActiveUserByUserName(strUserName string) (ok bool, err error) {
	var count int64
	var do *models.UserDO
	if count, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_ID).
		Eq(models.USER_COLUMN_USER_NAME, strUserName).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

func (dao *UserDAO) CheckActiveUserByEmail(strEmail string) (ok bool, err error) {
	var count int64
	var do *models.UserDO
	if count, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_ID).
		Eq(models.USER_COLUMN_EMAIL, strEmail).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

func (dao *UserDAO) CheckActiveUserByPhone(strPhone string) (ok bool, err error) {
	var count int64
	var do *models.UserDO
	if count, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_ID).
		Eq(models.USER_COLUMN_PHONE_NUMBER, strPhone).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	if count != 0 {
		return true, nil
	}
	return false, nil
}

func (dao *UserDAO) SelectUsers(pageNo, pageSize int) (dos []*models.UserDO, total int64, err error) {
	dos = make([]*models.UserDO, 0)
	if _, total, err = dao.db.Model(&dos).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_ID,
			models.USER_COLUMN_USER_NAME,
			models.USER_COLUMN_USER_ALIAS,
			models.USER_COLUMN_PHONE_NUMBER,
			models.USER_COLUMN_IS_ADMIN,
			models.USER_COLUMN_EMAIL,
			models.USER_COLUMN_ADDRESS,
			models.USER_COLUMN_REMARK,
			models.USER_COLUMN_STATE,
			models.USER_COLUMN_LOGIN_IP,
			models.USER_COLUMN_LOGIN_TIME,
			models.USER_COLUMN_CREATED_TIME,
			models.USER_COLUMN_UPDATED_TIME,
		).
		Page(pageNo, pageSize).
		Eq(models.USER_COLUMN_DELETED, 0).
		QueryEx(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

// user name
func (dao *UserDAO) SelectUserByName(strName string) (do *models.UserDO, err error) {
	if _, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_ID,
			models.USER_COLUMN_USER_NAME,
			models.USER_COLUMN_PASSWORD,
			models.USER_COLUMN_SALT,
			models.USER_COLUMN_USER_ALIAS,
			models.USER_COLUMN_PHONE_NUMBER,
			models.USER_COLUMN_IS_ADMIN,
			models.USER_COLUMN_EMAIL,
			models.USER_COLUMN_ADDRESS,
			models.USER_COLUMN_REMARK,
			models.USER_COLUMN_STATE,
			models.USER_COLUMN_LOGIN_IP,
			models.USER_COLUMN_LOGIN_TIME,
			models.USER_COLUMN_CREATED_TIME,
			models.USER_COLUMN_UPDATED_TIME,
		).
		Eq(models.USER_COLUMN_USER_NAME, strName).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

// user email
func (dao *UserDAO) SelectUserByEmail(strEmail string) (do *models.UserDO, err error) {
	if _, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_ID,
			models.USER_COLUMN_USER_NAME,
			models.USER_COLUMN_PASSWORD,
			models.USER_COLUMN_SALT,
			models.USER_COLUMN_USER_ALIAS,
			models.USER_COLUMN_PHONE_NUMBER,
			models.USER_COLUMN_IS_ADMIN,
			models.USER_COLUMN_EMAIL,
			models.USER_COLUMN_ADDRESS,
			models.USER_COLUMN_REMARK,
			models.USER_COLUMN_STATE,
			models.USER_COLUMN_LOGIN_IP,
			models.USER_COLUMN_LOGIN_TIME,
			models.USER_COLUMN_CREATED_TIME,
			models.USER_COLUMN_UPDATED_TIME,
		).
		Eq(models.USER_COLUMN_EMAIL, strEmail).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) SelectUserByPhone(strPhone string) (do *models.UserDO, err error) {
	if _, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_ID,
			models.USER_COLUMN_USER_NAME,
			models.USER_COLUMN_PASSWORD,
			models.USER_COLUMN_SALT,
			models.USER_COLUMN_USER_ALIAS,
			models.USER_COLUMN_PHONE_NUMBER,
			models.USER_COLUMN_IS_ADMIN,
			models.USER_COLUMN_EMAIL,
			models.USER_COLUMN_ADDRESS,
			models.USER_COLUMN_REMARK,
			models.USER_COLUMN_STATE,
			models.USER_COLUMN_LOGIN_IP,
			models.USER_COLUMN_LOGIN_TIME,
			models.USER_COLUMN_CREATED_TIME,
			models.USER_COLUMN_UPDATED_TIME,
		).
		Eq(models.USER_COLUMN_PHONE_NUMBER, strPhone).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) SelectActiveUserPasswordAndSalt(strUserName string) (do *models.UserDO, err error) {
	if _, err = dao.db.Model(&do).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_PASSWORD,
			models.USER_COLUMN_SALT,
		).
		Eq(models.USER_COLUMN_USER_NAME, strUserName).
		Eq(models.USER_COLUMN_DELETED, 0).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *UserDAO) SelectUserName(req *proto.PlatformUserQueryReq) (dos []*models.UserDO, total int64, err error) {
	dos = make([]*models.UserDO, 0)
	c := dao.db.Model(&dos).
		Table(models.TableNameUser).
		Select(
			models.USER_COLUMN_ID,
			models.USER_COLUMN_USER_NAME,
		).
		Where("%s=0", models.USER_COLUMN_DELETED)
	if req.Name != "" {
		c.And("%s like '%%%s%%'", models.USER_COLUMN_USER_NAME, req.Name)
	}
	if _, total, err = c.
		QueryEx(); err != nil {

		log.Errorf(err.Error())
		return
	}
	return
}

// user name
func (dao *UserDAO) SelectUserByDid(strDid string) (do *models.UserDO, err error) {
	//SELECT b.* FROM user_acl a, `user` b WHERE a.user_did='0xbeFDC8e41103D1F720B6F4D9aB046cE693521C4a' AND a.user_id=b.id AND b.deleted=0
	if _, err = dao.db.Model(&do).
		Select("b.*").
		Table("user_acl a, `user` b").
		Where("a.user_did='%s'", strDid).
		And("a.user_id=b.id").
		And("b.deleted=0").
		Query(); err != nil {

		log.Errorf(err.Error())
		return nil, err
	}
	return
}

func (dao *UserDAO) UpdateUserState(strUserName string, state int) error {
	if _, err := dao.db.Model(&state).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_STATE).
		Where("%s='%s'", models.USER_COLUMN_USER_NAME, strUserName).
		Update(); err != nil {
		log.Errorf(err.Error())
		return err
	}
	return nil
}

func (dao *UserDAO) UpdateUserStateById(id int32, state int) error {
	if _, err := dao.db.Model(&state).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_STATE).
		Where("%s=%d", models.USER_COLUMN_ID, id).
		Update(); err != nil {
		log.Errorf(err.Error())
		return err
	}
	return nil
}

func (dao *UserDAO) IsUserBanned(userId int32) (bool, error) {
	var state int32
	if _, err := dao.db.Model(&state).
		Table(models.TableNameUser).
		Select(models.USER_COLUMN_STATE).
		Eq(models.USER_COLUMN_ID, userId).
		Query(); err != nil {
		log.Errorf(err.Error())
		return false, err
	}
	if state == UserState_Disabled {
		return true, nil
	}
	return false, nil
}
