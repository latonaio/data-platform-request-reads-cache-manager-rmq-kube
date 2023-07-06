package routers

import (
	cache "data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	"data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()
	redisCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l)

	billOfMaterialController := &controllers.BillOfMaterialListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	billOfMaterial := beego.NewNamespace("/bill-of-material", beego.NSCond(
		func(ctx *context.Context) bool { return true },
	),
		beego.NSRouter("/list/:userType", billOfMaterialController),
	)

	beego.AddNamespace(billOfMaterial)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	//beego.Router("/:aPIServiceName/:aPIType", &controllers.APIModuleRuntimesController{})
	//beego.Router("/register", &controllers.RegisterController{})
	//beego.Router("/login", &controllers.LoginController{})
}
