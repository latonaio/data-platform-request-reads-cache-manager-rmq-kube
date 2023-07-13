package routers

import (
	"data-platform-request-reads-cache-manager-rmq-kube/cache"
	"data-platform-request-reads-cache-manager-rmq-kube/config"
	controllersBillOfMaterialDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/detail-list"
	controllersBillOfMaterialList "data-platform-request-reads-cache-manager-rmq-kube/controllers/bill-of-material/list"
	controllersBusinessPartnerList "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/list"
	controllersBusinessPartnerDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/business-partner/detail-general"
	controllersOperationsDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/detail-list"
	controllersOperationsList "data-platform-request-reads-cache-manager-rmq-kube/controllers/operations/list"
	controllersPriceMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/price-master/list"
	controllersPriceMasterDetailList "data-platform-request-reads-cache-manager-rmq-kube/controllers/price-master/detail-list"
	controllersProductMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/list"
	controllersProductMasterDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/product-master/detail-general"
	controllersSupplyChainRelationshipList "data-platform-request-reads-cache-manager-rmq-kube/controllers/supply-chain-relationship/list"
	controllersEquipmentMasterList "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/list"
	controllersEquipmentMasterDetailGeneral "data-platform-request-reads-cache-manager-rmq-kube/controllers/equipment-master/detail-general"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func init() {
	l := logger.NewLogger()
	conf := config.NewConf()

	redisCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 0, nil)
	//_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, "token")

	redisTokenCacheKeyPrefix := "tokens"

	_ = cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1, &redisTokenCacheKeyPrefix)
	//redisTokenCache := cache.NewCache(conf.REDIS.Address, conf.REDIS.Port, l, 1)

	billOfMaterialListController := &controllersBillOfMaterialList.BillOfMaterialListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	billOfMaterialDetailListController := &controllersBillOfMaterialDetailList.BillOfMaterialDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	operationsListController := &controllersOperationsList.OperationsListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	operationsDetailListController := &controllersOperationsDetailList.OperationsDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterListController := &controllersProductMasterList.ProductMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	productMasterDetailGeneralController := &controllersProductMasterDetailGeneral.ProductMasterDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	priceMasterListController := &controllersPriceMasterList.PriceMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	priceMasterDetailListController := &controllersPriceMasterDetailList.PriceMasterDetailListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	supplyChainRelationshipListController := &controllersSupplyChainRelationshipList.SupplyChainRelationshipListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	businessPartnerListController := &controllersBusinessPartnerList.BusinessPartnerListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	businessPartnerDetailGeneralController := &controllersBusinessPartnerDetailGeneral.BusinessPartnerDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	equipmentMasterListController := &controllersEquipmentMasterList.EquipmentMasterListController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	equipmentMasterDetailGeneralController := &controllersEquipmentMasterDetailGeneral.EquipmentMasterDetailGeneralController{
		RedisCache:   redisCache,
		CustomLogger: l,
	}

	billOfMaterial := beego.NewNamespace(
		"/bill-of-material",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", billOfMaterialListController),
		beego.NSRouter("/detail/list/:userType", billOfMaterialDetailListController),
	)

	operations := beego.NewNamespace(
		"/operations",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", operationsListController),
		beego.NSRouter("/detail/list/:userType", operationsDetailListController),
	)

	productMaster := beego.NewNamespace(
		"/product-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", productMasterListController),
		beego.NSRouter("/list/:userType", productMasterDetailGeneralController),
	)

	priceMaster := beego.NewNamespace(
		"/price-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", priceMasterListController),
		beego.NSRouter("/detail/list/:userType", priceMasterDetailListController),
	)

	supplyChainRelationship := beego.NewNamespace(
		"/supply-chain-relationship",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", supplyChainRelationshipListController),
	)

	businessPartner := beego.NewNamespace(
		"/business-partner",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", businessPartnerListController),
		beego.NSRouter("/list/:userType", businessPartnerDetailGeneralController),
	)

	equipmentMaster := beego.NewNamespace(
		"/equipment-master",
		beego.NSCond(func(ctx *context.Context) bool { return true }),
		beego.NSRouter("/list/:userType", equipmentMasterListController),
		beego.NSRouter("/list/:userType", equipmentMasterDetailGeneralController),
	)

	beego.AddNamespace(
		businessPartner,
		productMaster,
		billOfMaterial,
		operations,
		supplyChainRelationship,
		priceMaster,
		equipmentMaster,
	)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, func(ctx *context.Context) {
		_ = ctx.Input.Header("Authorization")
		//jwtToken := ctx.Input.Header("Authorization")
		//
		//trimJwtToken := strings.TrimPrefix(jwtToken, "Bearer ")
		//
		//token, err := redisTokenCache.GetRaw(trimJwtToken)
		//
		//fmt.Sprintf("token: %v", token)
		//
		//if err == nil {
		//	return
		//}

		//services.VerifyToken(ctx, l, jwtToken)
	})

//	beego.AddNamespace(billOfMaterial)

	//beego.Router("/:aPIServiceName/:aPIType", &controllers.APIModuleRuntimesController{})
	//beego.Router("/register", &controllers.RegisterController{})
	//beego.Router("/login", &controllers.LoginController{})
}
