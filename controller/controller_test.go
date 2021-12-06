package controller

import (
	"member-service/config"
	"member-service/consumer"
	"member-service/producer"
	"member-service/repository"
	"member-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	memberController.Route(app)
	addressController.RouteAddress(app)
	memberaddressController.RouteMemberAddress(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMongoDatabase(configuration)

var nsq = config.NewNsq(configuration)
var producerconfig = config.NewProducer(nsq, configuration)
var rabbitmq = config.NewRabbitMq(configuration)
var messageQueue = config.NewRabbitConsumeMq(configuration)
var oss = config.NewOss(configuration)
var bucketossProfile = config.NewBucket(configuration.Get("OSS_BUCKET"), oss, configuration)

// producer = config.NewProducer(configuration)

// Setup Repository
var memberRepository = repository.NewMemberRepository(database)
var addressRepository = repository.NewAddressRepository(database)
var provinsiRepository = repository.NewProvinsiRepository(database)
var kotakabupatenRepository = repository.NewKotakabupatenRepository(database)
var kecamatanRepository = repository.NewKecamatanRepository(database)
var kelurahanRepository = repository.NewKelurahanRepository(database)
var postcodeRepository = repository.NewPostcodeRepository(database)
var memberaddressRepository = repository.NewMemberAddressRepository(database)

// Setup Producer
var registerProducer = producer.NewRegisterProducer(producerconfig, rabbitmq)

// Setup Service
var memberService = service.NewMemberService(&memberRepository, &registerProducer, configuration, bucketossProfile)
var addressService = service.NewAddressService(&addressRepository, &provinsiRepository, &kotakabupatenRepository, &kecamatanRepository, &kelurahanRepository, &postcodeRepository, &registerProducer, configuration, bucketossProfile)
var memberaddressService = service.NewMemberAddressService(&memberaddressRepository, &addressRepository, &registerProducer, configuration, bucketossProfile)

// Setup Consumer
var createMemberConsumer = consumer.NewMemberConsumer(messageQueue, memberService)

// Setup Controller
var memberController = NewMemberController(&memberService, &addressService)
var addressController = NewAddressController(&memberService, &addressService)
var memberaddressController = NewMemberAddressController(&memberaddressService, &memberService, &addressService)

// go createMemberConsumer.AddNewMember()

var app = createTestApp()
