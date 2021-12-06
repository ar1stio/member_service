package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewKelurahanRepository(database *mongo.Database) KelurahanRepository {
	return &kelurahanRepositoryImpl{
		Collection: database.Collection("kelurahan"),
	}
}

type kelurahanRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *kelurahanRepositoryImpl) FindAllColKelurahan(request model.GetKelurahanRequest) (response []model.GetKelurahanResponse) {
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
