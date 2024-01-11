package sessions

import (
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"admin-system/pkg/middleware"
	"admin-system/pkg/storage"
	"admin-system/pkg/types"
	"sync"
)

var locker sync.RWMutex
var contexts = make(map[int32]*types.Context, 0)

func init() {
	preloadContexts()
}

func NewContext(s *types.Session) *types.Context {
	ctx := &types.Context{
		Session: &types.Session{
			UserId:      s.UserId,
			UserName:    s.UserName,
			Alias:       s.Alias,
			PhoneNumber: s.PhoneNumber,
			IsAdmin:     s.IsAdmin,
			LoginIP:     s.LoginIP,
			AuthToken:   s.AuthToken,
		},
	}
	locker.Lock()
	defer locker.Unlock()
	contexts[s.UserId] = ctx
	putLevelDB(ctx)
	return ctx
}

func RemoveContext(c *gin.Context) *types.Context {
	strToken := middleware.GetAuthToken(c)
	locker.Lock()
	defer locker.Unlock()
	for k, v := range contexts {
		if v.AuthToken() == strToken {
			delete(contexts, k)
			deleteLevelDB(k)
			return v
		}
	}
	return nil
}

func GetContext(c *gin.Context) *types.Context {
	strToken := middleware.GetAuthToken(c)
	locker.RLock()
	defer locker.RUnlock()
	for _, v := range contexts {
		if v.AuthToken() == strToken {
			return v
		}
	}
	return nil
}

func putLevelDB(ctx *types.Context) {
	if err := storage.Store.PutContext(ctx); err != nil {
		log.Errorf("level db put error [%s]", err)
		return
	}
}

func deleteLevelDB(userId int32) {
	storage.Store.DeleteContext(userId)
}

func preloadContexts() {
	var err error
	locker.Lock()
	defer locker.Unlock()
	if contexts, err = storage.Store.LoadContexts(); err != nil {
		log.Errorf(err.Error())
	}
}
