package repository

import (
	"member-service/config"
	"member-service/entity"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMemberAddressRepository(database *mongo.Database) MemberAddressRepository {
	return &memberAddressRepositoryImpl{
		Collection: database.Collection("member_address"),
	}
}

type memberAddressRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *memberAddressRepositoryImpl) CreateAddress(address entity.Address) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":                address.Id,
		"user_id":            address.UserId,
		"name":               address.Name,
		"is_default":         address.IsDefault,
		"receiver_name":      address.ReceiverName,
		"phone":              address.Phone,
		"_id_provinsi":       address.IdProvinsi,
		"provinsi":           address.Provinsi,
		"_id_kota_kabupaten": address.IdKotaKabupaten,
		"kota_kabupaten":     address.KotaKabupaten,
		"_id_kecamatan":      address.IdKecamatan,
		"kecamatan":          address.Kecamatan,
		"_id_kelurahan":      address.IdKelurahan,
		"kelurahan":          address.Kelurahan,
		"_id_kode_pos":       address.IdKodepos,
		"kode_pos":           address.Kodepos,
		"latitude":           address.Latitude,
		"longitude":          address.Longitude,
		"address":            address.Address,
		"created":            address.Created,
		"modified":           address.Modified,
	})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) UpdateAddress(address entity.Address) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": address.Id}, bson.M{"$set": bson.M{
		"user_id":            address.UserId,
		"name":               address.Name,
		"is_default":         address.IsDefault,
		"receiver_name":      address.ReceiverName,
		"phone":              address.Phone,
		"_id_provinsi":       address.IdProvinsi,
		"provinsi":           address.Provinsi,
		"_id_kota_kabupaten": address.IdKotaKabupaten,
		"kota_kabupaten":     address.KotaKabupaten,
		"_id_kecamatan":      address.IdKecamatan,
		"kecamatan":          address.Kecamatan,
		"_id_kelurahan":      address.IdKelurahan,
		"kelurahan":          address.Kelurahan,
		"_id_kode_pos":       address.IdKodepos,
		"kode_pos":           address.Kodepos,
		"latitude":           address.Latitude,
		"longitude":          address.Longitude,
		"address":            address.Address,
		"created":            address.Created,
		"modified":           address.Modified,
	}})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) DeleteAddress(address model.DeleteAddressRequest) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.DeleteMany(ctx, bson.M{
		"_id": address.Id,
	})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) FindAddress(request model.GetAddressRequest) (response []model.GetAddressFindResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var id = bson.M{"_id": request.Filter}
	var user_id = bson.M{"user_id": request.Filter}
	var name = bson.M{"name": request.Filter}

	filteror := bson.A{}
	filteror = append(filteror, id)
	filteror = append(filteror, user_id)
	filteror = append(filteror, name)
	filter := bson.M{"$or": filteror}

	// filter := bson.M{"$or": filteror}
	// filter := bson.M{}
	// filter = bson.M{"_id_kota_kabupaten": request.Id}

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		response = append(response, model.GetAddressFindResponse{
			Id: document["_id"].(string),
			// UserId:          document["user_id"].(string),
			Name:            document["name"].(string),
			IsDefault:       document["is_default"].(bool),
			ReceiverName:    document["receiver_name"].(string),
			Phone:           document["phone"].(string),
			IdProvinsi:      document["_id_provinsi"].(string),
			Provinsi:        document["provinsi"].(string),
			IdKotaKabupaten: document["_id_kota_kabupaten"].(string),
			KotaKabupaten:   document["kota_kabupaten"].(string),
			IdKecamatan:     document["_id_kecamatan"].(string),
			Kecamatan:       document["kecamatan"].(string),
			IdKelurahan:     document["_id_kelurahan"].(string),
			Kelurahan:       document["kelurahan"].(string),
			IdKodepos:       document["_id_kode_pos"].(string),
			Kodepos:         document["kode_pos"].(string),
			Latitude:        document["latitude"].(float64),
			Longitude:       document["longitude"].(float64),
			Address:         document["address"].(string),
		})
	}

	return response
}

func (repository *memberAddressRepositoryImpl) UpdateNonDefaultAddress(address model.UpdateIsDefaultRequest) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateMany(ctx, bson.M{"user_id": address.UserId}, bson.M{"$set": bson.M{
		"is_default": false,
	}})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) UpdateIsDefaultAddress(address model.UpdateIsDefaultRequest) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": address.Id}, bson.M{"$set": bson.M{
		"is_default": true,
	}})
	exception.PanicIfNeeded(err)
}

func (repository *memberAddressRepositoryImpl) FindAddressOne(request model.GetAddressOneRequest) (memberaddress model.GetCekAddressFindResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var username = bson.M{}
	if request.Type == "id" {
		username = bson.M{"_id": request.Filter}
	} else if request.Type == "nama" {
		username = bson.M{"name": request.Filter, "user_id": request.UserId}
	}
	filter := username

	var test bson.M

	errone := repository.Collection.FindOne(ctx, filter).Decode(&test)

	if errone == nil && test != nil {
		memberaddress.Id = test["_id"].(string)
		memberaddress.UserId = test["user_id"].(string)
		memberaddress.Name = test["name"].(string)
		memberaddress.IsDefault = test["is_default"].(bool)
		memberaddress.ReceiverName = test["receiver_name"].(string)
		memberaddress.Phone = test["phone"].(string)
		memberaddress.IdProvinsi = test["_id_provinsi"].(string)
		memberaddress.Provinsi = test["provinsi"].(string)
		memberaddress.IdKotaKabupaten = test["_id_kota_kabupaten"].(string)
		memberaddress.KotaKabupaten = test["kota_kabupaten"].(string)
		memberaddress.IdKecamatan = test["_id_kecamatan"].(string)
		memberaddress.Kecamatan = test["kecamatan"].(string)
		memberaddress.IdKelurahan = test["_id_kelurahan"].(string)
		memberaddress.Kelurahan = test["kelurahan"].(string)
		memberaddress.IdKodepos = test["_id_kode_pos"].(string)
		memberaddress.Kodepos = test["kode_pos"].(string)
		memberaddress.Latitude = test["latitude"].(float64)
		memberaddress.Longitude = test["longitude"].(float64)
		memberaddress.Address = test["address"].(string)
	}

	return memberaddress
}
