package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPostcodeRepository(database *mongo.Database) PostcodeRepository {
	return &postcodeRepositoryImpl{
		Collection: database.Collection("postcode"),
	}
}

type postcodeRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *postcodeRepositoryImpl) FindAllColPostcode(request model.GetPostcodeRequest) (response []model.GetPostcodeResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	// if request.Postcode != "" {
	// filter = bson.M{"$text": bson.M{"$search": "'" + request.Postcode + "'"}, "_id": request.Id}
	// } else {
	filter = bson.M{"kodepos": request.Kodepos}
	// }

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetPostcodeResponse{
			Id:              document["_id"].(int32),
			IdKelurahan:     document["_id_kelurahan"].(string),
			IdProvinsi:      document["_id_provinsi"].(string),
			Provinsi:        document["provinsi"].(string),
			IdKotaKabupaten: document["_id_kota_kabupaten"].(string),
			KotaKabupaten:   document["kota_kabupaten"].(string),
			IdKecamatan:     document["_id_kecamatan"].(string),
			Kecamatan:       document["kecamatan"].(string),
			Kelurahan:       document["kelurahan"].(string),
			Kodepos:         document["kodepos"].(string),
		})
	}

	return response
}
