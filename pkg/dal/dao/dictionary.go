package dao

import (
	"admin-system/pkg/dal/models"
	"admin-system/pkg/proto"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

const (
	Dictionary_Name_Email_Server    = "邮箱服务器"
	Dictionary_Name_Email_Port      = "端口"
	Dictionary_Name_Email_Name      = "邮箱名"
	Dictionary_Name_Email_Auth_Code = "授权码"
	Dictionary_Name_Email_Send_Name = "发件人名称"
)

const (
	Dictionary_Key_Email_Server    = "email_server"
	Dictionary_Key_Email_Port      = "port"
	Dictionary_Key_Email_Name      = "email_name"
	Dictionary_Key_Email_Auth_Code = "auth_code"
	Dictionary_Key_Email_Send_Name = "send_name"
)

const (
	Dictionary_Remark_Email_Server    = "邮箱SMTP服务器"
	Dictionary_Remark_Email_Port      = "SMTP服务器端口号"
	Dictionary_Remark_Email_Name      = "SMTP邮箱服务器用户自己的邮箱名"
	Dictionary_Remark_Email_Auth_Code = "SMTP服务器密码，这里是设置账户中的授权码"
	Dictionary_Remark_Email_Send_Name = "邮件发送人名称"
)

type DictionaryDAO struct {
	db *sqlca.Engine
}

func NewDictionaryDAO(db *sqlca.Engine) *DictionaryDAO {

	return &DictionaryDAO{
		db: db,
	}
}

func (dao *DictionaryDAO) SelectKey(key string) (do *models.DictionaryDO, ok bool, err error) {
	var count int64
	if count, err = dao.db.Model(&do).
		Table(models.TableNameDictionary).
		Where("`%s`='%s'", models.DICTIONARY_COLUMN_CONFIG_KEY, key).
		And("%s=0", models.DICTIONARY_COLUMN_DELETED).
		Query(); err != nil {

		log.Errorf(err.Error())
		return
	}
	if count != 0 {
		return do, true, nil
	}
	return do, false, nil
}

func (dao *DictionaryDAO) Insert(do *models.DictionaryDO) (id int64, err error) {
	if id, err = dao.db.Model(&do).Table(models.TableNameDictionary).Insert(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *DictionaryDAO) UpdateByKey(do *models.DictionaryDO, columns ...string) (err error) {
	if _, err = dao.db.Model(do).
		Table(models.TableNameDictionary).
		Select(columns...).
		Exclude(models.DICTIONARY_COLUMN_CONFIG_KEY).
		Where("`%s`='%s'", models.DICTIONARY_COLUMN_CONFIG_KEY, do.ConfigKey).
		Update(); err != nil {
		log.Errorf(err.Error())
		return
	}
	return
}

func (dao *DictionaryDAO) Upsert(do *models.DictionaryDO) (err error) {
	var ok bool
	_, ok, err = dao.SelectKey(do.ConfigKey)
	if err != nil {
		return err
	}
	if ok {
		err = dao.UpdateByKey(do)
		if err != nil {
			return err
		}
	} else {
		_, err = dao.Insert(do)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dao *DictionaryDAO) SelectEmailConfig() (emailConfig proto.PlatformEmailConfigListResp) {
	emailServer, _, _ := dao.SelectKey(Dictionary_Key_Email_Server)
	port, _, _ := dao.SelectKey(Dictionary_Key_Email_Port)
	emailName, _, _ := dao.SelectKey(Dictionary_Key_Email_Name)
	authCode, _, _ := dao.SelectKey(Dictionary_Key_Email_Auth_Code)
	sendName, _, _ := dao.SelectKey(Dictionary_Key_Email_Send_Name)
	emailConfig = proto.PlatformEmailConfigListResp{
		EmailServer: emailServer.Value,
		Port:        port.Value,
		EmailName:   emailName.Value,
		AuthCode:    authCode.Value,
		SendName:    sendName.Value,
	}
	return emailConfig
}
