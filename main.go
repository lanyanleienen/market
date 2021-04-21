package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lanyanleienen/market/common"
	"github.com/lanyanleienen/market/domain/repository"
	service2 "github.com/lanyanleienen/market/domain/service"
	"github.com/lanyanleienen/market/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	market "github.com/lanyanleienen/market/proto/market"
)

var QPS = 100

func main() {
	//配置中心
	consulConfig,err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil{
		log.Error(err)
	}
	//注册中心
	reg := consul2.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//jaeger链路追踪
	t,io,err := common.NewTracer("go.micro.service.market", "127.0.0.1:6831")
	if err != nil{
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	//数据库
	//db,err := gorm.Open("mysql", "root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	db,err:= gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//repository.NewMarketRepository(db).InitTable()

	marketService := service2.NewMarketService(repository.NewMarketRepository(db))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.market"),
		micro.Version("latest"),
		micro.Address(":8082"),
		micro.Registry(reg),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	market.RegisterMarketHandler(service.Server(), &handler.Market{marketService})


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
