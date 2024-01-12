package middleware

import (
	"admin-system/pkg/types"
	"github.com/civet148/log"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	strToken, err := GenerateToken(&types.Session{
		UserId:   1,
		UserName: "admin",
		IsAdmin:  true,
	}, 6000*time.Hour)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("token [%s]", strToken)
	var session types.Session
	err = GetAuthSession(strToken, &session)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("session [%+v]", session)
}
