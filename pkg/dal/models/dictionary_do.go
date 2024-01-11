// Code generated by db2go. DO NOT EDIT.
// https://github.com/civet148/sqlca

package models

const TableNameDictionary = "dictionary" // 

const (
DICTIONARY_COLUMN_ID = "id"
DICTIONARY_COLUMN_NAME = "name"
DICTIONARY_COLUMN_CONFIG_KEY = "config_key"
DICTIONARY_COLUMN_VALUE = "value"
DICTIONARY_COLUMN_REMARK = "remark"
DICTIONARY_COLUMN_DELETED = "deleted"
DICTIONARY_COLUMN_CREATED_TIME = "created_time"
DICTIONARY_COLUMN_UPDATED_TIME = "updated_time"
)

type DictionaryDO struct { 
	Id int32 `json:"id" db:"id" bson:"_id"` //自增ID 
	Name string `json:"name" db:"name" bson:"name"` //名称 
	ConfigKey string `json:"config_key" db:"config_key" bson:"config_key"` //KEY 
	Value string `json:"value" db:"value" bson:"value"` //VALUE 
	Remark string `json:"remark" db:"remark" bson:"remark"` //备注 
	Deleted bool `json:"deleted" db:"deleted" bson:"deleted"` //是否已删除(0=未删除 1=已删除) 
	CreatedTime string `json:"created_time" db:"created_time" sqlca:"readonly" bson:"created_time"` //创建时间 
	UpdatedTime string `json:"updated_time" db:"updated_time" sqlca:"readonly" bson:"updated_time"` //更新时间 
}

func (do *DictionaryDO) GetId() int32 { return do.Id } 
func (do *DictionaryDO) SetId(v int32) { do.Id = v } 
func (do *DictionaryDO) GetName() string { return do.Name } 
func (do *DictionaryDO) SetName(v string) { do.Name = v } 
func (do *DictionaryDO) GetConfigKey() string { return do.ConfigKey } 
func (do *DictionaryDO) SetConfigKey(v string) { do.ConfigKey = v } 
func (do *DictionaryDO) GetValue() string { return do.Value } 
func (do *DictionaryDO) SetValue(v string) { do.Value = v } 
func (do *DictionaryDO) GetRemark() string { return do.Remark } 
func (do *DictionaryDO) SetRemark(v string) { do.Remark = v } 
func (do *DictionaryDO) GetDeleted() bool { return do.Deleted } 
func (do *DictionaryDO) SetDeleted(v bool) { do.Deleted = v } 
func (do *DictionaryDO) GetCreatedTime() string { return do.CreatedTime } 
func (do *DictionaryDO) SetCreatedTime(v string) { do.CreatedTime = v } 
func (do *DictionaryDO) GetUpdatedTime() string { return do.UpdatedTime } 
func (do *DictionaryDO) SetUpdatedTime(v string) { do.UpdatedTime = v } 
/*
CREATE TABLE `dictionary` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `config_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'KEY',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'VALUE',
  `remark` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已删除(0=未删除 1=已删除)',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`config_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;
*/