package main

import (
	"github.com/civet148/log"
	"admin-system/pkg/middleware"
	"admin-system/pkg/types"
	"time"
)

func main() {
	strToken, claims, err := middleware.GenerateToken(&types.Session{
		UserId:   1,
		UserName: "admin",
		IsAdmin:  true,
	}, 6000*time.Hour)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("token [%s]", strToken)
	log.Infof("claims [%+v]", claims)
}
