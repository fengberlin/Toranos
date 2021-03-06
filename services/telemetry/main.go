package main

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/loehnertz/Toranos/common"
	"github.com/loehnertz/Toranos/services/telemetry/proto"
	"github.com/micro/go-config"
	"github.com/micro/go-micro"
	"time"
)

var conf config.Config
var redisClient *redis.Client
var service micro.Service

type Telemetry struct{}

func (tm *Telemetry) AllVehicles(ctx context.Context, req *telemetry.Empty, res *telemetry.AllVehiclesResponse) error {
	res.Vehicles = retrieveAllVehicles(redisClient)

	return nil
}

func main() {
	// Initialize the configuration
	conf = common.InitConfig()

	// Initialize a Redis client
	redisClient = common.InitRedisClient(
		common.GetConfigStringByPath(conf, "caching", "redis", "host"),
		common.GetConfigStringByPath(conf, "caching", "redis", "password"),
		common.GetConfigIntByPath(conf, "caching", "redis", "databaseId"),
	)

	// Create the service
	service = micro.NewService(
		micro.Name(common.GetConfigStringByPath(conf, "service-names", "telemetry")),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()

	// Register the handler
	telemetry.RegisterTelemetryHandler(service.Server(), new(Telemetry))

	// Run the server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
