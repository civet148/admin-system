package storage

import (
	"admin-system/pkg/config"
	"admin-system/pkg/types"
	"encoding/json"
	"fmt"
	"github.com/civet148/log"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"strconv"
	"strings"
)

const (
	LEVEL_DB_KEY_PREFFIX_CONTEXT    = "/context/" //the last '/' cannot be removed
	LEVEL_DB_KEY_CONFIG_OSS_MANAGER = "/config/admin-system"
)

type LocalStorage struct {
	ldb *leveldb.DB
}

var Store *LocalStorage

func init() {
	Store = NewLocalStorage()
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) open() (err error) {
	s.ldb, err = leveldb.OpenFile(types.DefaultLevelDBHome, nil)
	if err != nil {
		return fmt.Errorf("open level db path [%s] error [%s]", types.DefaultLevelDBHome, err.Error())
	}
	return
}

func (s *LocalStorage) close() (err error) {
	return s.ldb.Close()
}

func (s *LocalStorage) delete(strKey string) (err error) {
	return s.ldb.Delete([]byte(strKey), nil)
}

func (s *LocalStorage) put(strKey string, data []byte) (err error) {
	if err = s.ldb.Put([]byte(strKey), data, nil); err != nil {
		log.Errorf("level db put [%s] error [%s]", strKey, err)
		return
	}
	return
}

func (s *LocalStorage) clean() error {
	return os.RemoveAll(types.DefaultLevelDBHome)
}

func (s *LocalStorage) get(strKey string) []byte {
	data, err := s.ldb.Get([]byte(strKey), nil)
	if err != nil {
		log.Errorf("level db get [%s] error [%s]", strKey, err)
		return nil
	}
	return data
}

// clean storage directory
func (s *LocalStorage) Clean() error {
	return s.clean()
}

func (s *LocalStorage) DeleteContext(userId int32) {
	var strKey = s.MakeContextKey(userId)
	if err := s.open(); err != nil {
		log.Errorf(err.Error())
		return
	}
	defer s.close()
	if err := s.delete(strKey); err != nil {
		log.Errorf(err.Error())
	}
}

func (s *LocalStorage) PutContext(ctx *types.Context) (err error) {
	var data []byte
	var userId = ctx.UserId()
	data, err = json.Marshal(ctx)
	if err != nil {
		log.Errorf("json marshal error [%s]", err)
		return
	}
	if err = s.open(); err != nil {
		log.Errorf(err.Error())
		return
	}
	defer s.close()
	var strKey = s.MakeContextKey(userId)
	if err = s.put(strKey, data); err != nil {
		log.Errorf("level db put error [%s]", err)
		return
	}
	return
}

func (s *LocalStorage) LoadContexts() (contexts map[int32]*types.Context, err error) {
	contexts = make(map[int32]*types.Context, 0)
	if err = s.open(); err != nil {
		log.Errorf(err.Error())
		return
	}
	defer s.close()
	iter := s.ldb.NewIterator(nil, nil)
	for iter.Next() {
		strKey := string(iter.Key())
		value := iter.Value()

		if strings.HasPrefix(strKey, LEVEL_DB_KEY_PREFFIX_CONTEXT) {
			strKey = strings.TrimPrefix(strKey, LEVEL_DB_KEY_PREFFIX_CONTEXT)
			userId, _ := strconv.Atoi(strKey)
			if userId <= 0 {
				log.Errorf("user id %d <= 0", userId)
				continue
			}
			var ctx = &types.Context{}
			if err = json.Unmarshal(value, ctx); err != nil {
				log.Errorf("key [%s] value unmarshal error [%s]", strKey, err)
				continue
			}
			contexts[int32(userId)] = ctx
			log.Debugf("user id [%v] context [%+v] from level db ok", userId, ctx.Session)
		}
	}
	iter.Release()
	return
}

func (s *LocalStorage) MakeContextKey(userId int32) string {
	strKey := fmt.Sprintf("%s%v", LEVEL_DB_KEY_PREFFIX_CONTEXT, userId)
	return strKey
}

func (s *LocalStorage) PutConfig(cfg *config.Config) {
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if err := s.open(); err != nil {
		log.Errorf(err.Error())
		return
	}
	defer s.close()
	_ = s.put(LEVEL_DB_KEY_CONFIG_OSS_MANAGER, data)
}

func (s *LocalStorage) GetConfig() (cfg *config.Config) {
	cfg = &config.Config{}
	if err := s.open(); err != nil {
		log.Errorf(err.Error())
		return
	}
	defer s.close()
	data := s.get(LEVEL_DB_KEY_CONFIG_OSS_MANAGER)
	if len(data) == 0 {
		log.Error("get config data failed")
		return nil
	}
	if err := json.Unmarshal(data, cfg); err != nil {
		log.Error("unmarshal to config struct failed, data [%s]", data)
		return nil
	}
	return cfg
}
