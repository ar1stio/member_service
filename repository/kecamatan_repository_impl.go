package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewKecamatanRepository(database *mongo.Database) KecamatanRepository {
	return &kecamatanRepositoryImpl{
		Collection: database.Collection("kecamatan"),
	}
}

type kecamatanRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *kecamatanRepositoryImpl) FindAllColKecamatan(request model.GetKecamatanRequest) (response []model.GetKecamatanResponse) {
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
