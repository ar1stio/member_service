package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewAddressRepository(database *mongo.Database) AddressRepository {
	return &addressRepositoryImpl{
		Collection: database.Collection("postcode"),
	}
}

type addressRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *addressRepositoryImpl) FindAllProvinsi(request model.GetProvinsiRequest) (response []model.GetProvinsiResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	if request.Provinsi != "" {
		filter = bson.M{"$text": bson.M{"$search": "'" + request.Provinsi + "'"}, "_id_provinsi": request.Id}
	} else {
		filter = bson.M{"_id_provinsi": request.Id}
	}

	options := options.Find()
	// options.SetSort(bson.D{{"expired_at", -1}})
	options.SetLimit(1)

	cursor, err := repository.Collection.Find(ctx, filter, options)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetProvinsiResponse{
			Id:       document["_id_provinsi"].(string),
			Provinsi: document["provinsi"].(string),
		})
	}

	return response
}

func (repository *addressRepositoryImpl) FindAllKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	if request.KotaKabupaten != "" {
		filter = bson.M{"$text": bson.M{"$search": "'" + request.KotaKabupaten + "'"}, "_id_provinsi": request.Id}
	} else {
		filter = bson.M{"_id_provinsi": request.Id}
	}

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetKotaKabupatenResponse{
			Id:            document["_id"].(int32),
			KotaKabupaten: document["kota_kabupaten"].(string),
			IdProvinsi:    document["_id_provinsi"].(string),
			Provinsi:      document["provinsi"].(string),
		})
	}

	return response
}

func (repository *addressRepositoryImpl) FindAllKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	if request.Kecamatan != "" {
		filter = bson.M{"$text": bson.M{"$search": "'" + request.Kecamatan + "'"}, "_id_kota_kabupaten": request.Id}
	} else {
		filter = bson.M{"_id_kota_kabupaten": request.Id}
	}

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetKecamatanResponse{
			Id:            document["_id"].(int32),
			IdKotaKab:     document["_id_kota_kabupaten"].(string),
			KotaKabupaten: document["kota_kabupaten"].(string),
			IdProvinsi:    document["_id_provinsi"].(string),
			Provinsi:      document["provinsi"].(string),
			Kecamatan:     document["kecamatan"].(string),
		})
	}

	return response
}

func (repository *addressRepositoryImpl) FindAllKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	if request.Kelurahan != "" {
		filter = bson.M{"$text": bson.M{"$search": "'" + request.Kelurahan + "'"}, "_id_kecamatan": request.Id}
	} else {
		filter = bson.M{"_id_kecamatan": request.Id}
	}

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetKelurahanResponse{
			Id:              document["_id"].(int32),
			IdKecamatan:     document["_id_kecamatan"].(string),
			IdKotaKabupaten: document["_id_kota_kabupaten"].(string),
			KotaKabupaten:   document["kota_kabupaten"].(string),
			IdProvinsi:      document["_id_provinsi"].(string),
			Provinsi:        document["provinsi"].(string),
			Kecamatan:       document["kecamatan"].(string),
			Kelurahan:       document["kelurahan"].(string),
		})
	}

	return response
}

func (repository *addressRepositoryImpl) FindAllAddress(request model.GetAddressRequest) (response []model.GetAddressResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	options := options.Find()
	// options.SetSort(bson.D{{"expired_at", -1}})
	options.SetLimit(10)

	filter := bson.M{}
	filter = bson.M{"$text": bson.M{"$search": "'" + request.Filter + "'"}}

	cursor, err := repository.Collection.Find(ctx, filter, options)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetAddressResponse{
			Id:              document["_id"].(int32),
			IdKelurahan:     document["_id_kelurahan"].(string),
			IdKecamatan:     document["_id_kecamatan"].(string),
			IdKotaKabupaten: document["_id_kota_kabupaten"].(string),
			KotaKabupaten:   document["kota_kabupaten"].(string),
			IdProvinsi:      document["_id_provinsi"].(string),
			Provinsi:        document["provinsi"].(string),
			Kecamatan:       document["kecamatan"].(string),
			Kelurahan:       document["kelurahan"].(string),
			KodePos:         document["kodepos"].(string),
			Address:         document["address"].(string),
		})
	}

	return response
}

func (repository *addressRepositoryImpl) FindAllPostCode(request model.GetAddressByIdKecamatanRequest) (response []model.GetAddressResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	options := options.Find()
	// options.SetSort(bson.D{{"expired_at", -1}})
	options.SetLimit(10)

	filter := bson.M{}
	filter = bson.M{"_id_kecamatan": request.Id}

	cursor, err := repository.Collection.Find(ctx, filter, options)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetAddressResponse{
			Id:              document["_id"].(int32),
			IdKelurahan:     document["_id_kelurahan"].(string),
			IdKecamatan:     document["_id_kecamatan"].(string),
			IdKotaKabupaten: document["_id_kota_kabupaten"].(string),
			KotaKabupaten:   document["kota_kabupaten"].(string),
			IdProvinsi:      document["_id_provinsi"].(string),
			Provinsi:        document["provinsi"].(string),
			Kecamatan:       document["kecamatan"].(string),
			Kelurahan:       document["kelurahan"].(string),
			KodePos:         document["kodepos"].(string),
			Address:         document["address"].(string),
		})
	}

	return response
}
