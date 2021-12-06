package repository

import (
	"member-service/config"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProvinsiRepository(database *mongo.Database) ProvinsiRepository {
	return &provinsiRepositoryImpl{
		Collection: database.Collection("provinsi"),
	}
}

type provinsiRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *provinsiRepositoryImpl) FindAllColProvinsi(request model.GetColProvinsiRequest) (response []model.GetColProvinsiResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{}
	if request.Provinsi != "" {
		filter = bson.M{"$text": bson.M{"$search": "'" + request.Provinsi + "'"}, "_id": request.Id}
	} else {
		filter = bson.M{"_id": request.Id}
	}

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetColProvinsiResponse{
			Id:       document["_id"].(int32),
			Provinsi: document["provinsi"].(string),
		})
	}

	return response
}
