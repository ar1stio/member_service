package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewKotakabupatenRepository(database *mongo.Database) KotakabupatenRepository {
	return &kotakabupatenRepositoryImpl{
		Collection: database.Collection("kota_kabupaten"),
	}
}

type kotakabupatenRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *kotakabupatenRepositoryImpl) FindAllColKotaKabupaten(request model.GetKotaKabupatenRequest) (response []model.GetKotaKabupatenResponse) {
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
