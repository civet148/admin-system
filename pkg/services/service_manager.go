package services

import (
	cainit "admin-system/pkg/cache"
	"admin-system/pkg/config"
	"admin-system/pkg/controllers"
	"admin-system/pkg/middleware"
	"admin-system/pkg/routers"
	"admin-system/pkg/storage"
	"github.com/civet148/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Manager struct {
	*controllers.Controller
	cfg       *config.Config
	router    *gin.Engine
	routerRPC *gin.Engine
}

func NewManager(cfg *config.Config) *Manager {
	m := &Manager{
		cfg:        cfg,
		router:     gin.New(),
		routerRPC:  gin.New(),
		Controller: controllers.NewController(cfg),
	}
	return m
}

func (m *Manager) Run() (err error) {
	_ = m.runManager(func() error {
		//save config to local storage
		storage.Store.PutConfig(m.cfg)
		cainit.InitBigCache()
		//start up web service, if success this routine will be blocked
		if err = m.startWebService(); err != nil {
			m.Close()
			log.Errorf("start web service error [%s]", err)
			return err
		}
		return err
	})
	return
}

func (m *Manager) Close() {

}

func (m *Manager) initRouterMgr() (r *gin.Engine) {

	m.router.Use(gin.Logger())
	m.router.Use(gin.Recovery())
	m.router.Use(middleware.Cors())
	m.router.Static("/", m.cfg.Static)
	routers.InitRouterGroupPlatform(m.router, m) //平台系统管理
	return m.router
}

func (m *Manager) runManager(run func() error) (err error) {
	storage.Store.PutConfig(m.cfg)
	return run()
}

func (m *Manager) startWebService() (err error) {
	routerMgr := m.initRouterMgr()
	strHttpAddr := m.cfg.HttpAddr
	log.Infof("starting http server on %s \n", strHttpAddr)
	//Web manager service
	if err = http.ListenAndServe(strHttpAddr, routerMgr); err != nil { //if everything is fine, it will block this routine
		log.Panic("listen http server [%s] error [%s]\n", strHttpAddr, err)
	}
	return
}
