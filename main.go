package main

import (
	"member-service/config"
	"member-service/consumer"
	"member-service/controller"
	"member-service/exception"
	"member-service/producer"
	"member-service/repository"
	"member-service/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)
	// nsq := config.NewProducer(configuration)
	nsq := config.NewNsq(configuration)
	producerconfig := config.NewProducer(nsq, configuration)
	rabbitmq := config.NewRabbitMq(configuration)
	messageQueue := config.NewRabbitConsumeMq(configuration)
	oss := config.NewOss(configuration)
	bucketossProfile := config.NewBucket(configuration.Get("OSS_BUCKET"), oss, configuration)
	// producer := config.NewProducer(configuration)

	// Setup Repository
	memberRepository := repository.NewMemberRepository(database)
	addressRepository := repository.NewAddressRepository(database)
	provinsiRepository := repository.NewProvinsiRepository(database)
	kotakabupatenRepository := repository.NewKotakabupatenRepository(database)
	kecamatanRepository := repository.NewKecamatanRepository(database)
	kelurahanRepository := repository.NewKelurahanRepository(database)
	postcodeRepository := repository.NewPostcodeRepository(database)
	memberaddressRepository := repository.NewMemberAddressRepository(database)

	// Setup Producer
	registerProducer := producer.NewRegisterProducer(producerconfig, rabbitmq)

	// Setup Service
	memberService := service.NewMemberService(&memberRepository, &registerProducer, configuration, bucketossProfile)
	addressService := service.NewAddressService(&addressRepository, &provinsiRepository, &kotakabupatenRepository, &kecamatanRepository, &kelurahanRepository, &postcodeRepository, &registerProducer, configuration, bucketossProfile)
	memberaddressService := service.NewMemberAddressService(&memberaddressRepository, &addressRepository, &registerProducer, configuration, bucketossProfile)

	// Setup Consumer
	createMemberConsumer := consumer.NewMemberConsumer(messageQueue, memberService)

	// Setup Controller
	memberController := controller.NewMemberController(&memberService, &addressService)
	addressController := controller.NewAddressController(&memberService, &addressService)
	memberaddressController := controller.NewMemberAddressController(&memberaddressService, &memberService, &addressService)

	go createMemberConsumer.AddNewMember()

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	memberController.Route(app)
	addressController.RouteAddress(app)
	memberaddressController.RouteMemberAddress(app)

	// Start App
	err := app.Listen(":" + os.Getenv("PORT"))
	exception.PanicIfNeeded(err)
}
