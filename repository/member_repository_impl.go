package repository

import (
	"member-service/config"
	"member-service/entity"
	"member-service/exception"
	"member-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewMemberRepository(database *mongo.Database) MemberRepository {
	return &memberRepositoryImpl{
		Collection: database.Collection("member"),
	}
}

type memberRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *memberRepositoryImpl) FindAllMember(request model.GetMemberFindRequest) (response []model.GetMemberResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var username = bson.M{"_id": request.Filter}
	var gender = bson.M{"full_name": request.Filter}
	var phone = bson.M{"gender": request.Filter}

	filteror := bson.A{}
	filteror = append(filteror, username)
	filteror = append(filteror, gender)
	filteror = append(filteror, phone)
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
		response = append(response, model.GetMemberResponse{
			Id:             document["_id"].(string),
			FullName:       document["full_name"].(string),
			BirthDate:      document["birth_date"].(string),
			Gender:         document["gender"].(string),
			ProfilePicture: document["profile_picture"].(string),
		})
	}

	return response
}

func (repository *memberRepositoryImpl) Create(auths entity.Member) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":             auths.Id,
		"full_name":       auths.FullName,
		"birth_date":      auths.BirthDate,
		"profile_picture": auths.ProfilePicture,
		"object_picture":  auths.ObjectPicture,
		"gender":          auths.Gender,
		"created":         auths.Created,
		"modified":        auths.Modified,
		// "created_at":           auths.CreateAt,
		// "created_date_local":   auths.CreateDateLocal,
		// "created_date_server":  auths.CreateDateServer,
		// "modified_at":          auths.ModifiedAt,
		// "modified_date_local":  auths.ModifiedDateLocal,
		// "modified_date_server": auths.ModifiedDateServer,
	})
	exception.PanicIfNeeded(err)
}

func (repository *memberRepositoryImpl) Update(auths entity.Member) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": auths.Id}, bson.M{"$set": bson.M{
		"_id":             auths.Id,
		"full_name":       auths.FullName,
		"birth_date":      auths.BirthDate,
		"profile_picture": auths.ProfilePicture,
		"object_picture":  auths.ObjectPicture,
		"gender":          auths.Gender,
		"modified":        auths.Modified,
	}})
	exception.PanicIfNeeded(err)
}

func (repository *memberRepositoryImpl) DeleteFieldPicture(auths entity.Member) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": auths.Id}, bson.M{"$set": bson.M{
		"profile_picture": "",
		"object_picture":  "",
	}})
	exception.PanicIfNeeded(err)
}

func (repository *memberRepositoryImpl) UpdateProfile(auths entity.Member) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, err := repository.Collection.UpdateOne(ctx, bson.M{"_id": auths.Id}, bson.M{
		"profile_picture": auths.ProfilePicture,
		"modified":        auths.Modified,
	})
	exception.PanicIfNeeded(err)
}

func (repository *memberRepositoryImpl) FindAll() (auths []entity.Member) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		auths = append(auths, entity.Member{
			Id: document["_id"].(string),
		})
	}

	return auths
}

func (repository *memberRepositoryImpl) CheckUser(request string, check string) (member model.GetMemberResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	// var username = bson.M{"username": request}
	// var email = bson.M{"email": request}
	// var phone = bson.M{"phone": request}

	// filteror := bson.A{}
	// filteror = append(filteror, username)
	// filteror = append(filteror, email)
	// filteror = append(filteror, phone)
	filter := bson.M{"$or": "filteror"}

	err := repository.Collection.FindOne(ctx, filter).Decode(&member)
	// exception.PanicIfNeeded(err)
	if err == nil {
		exception.PanicIfNeeded(exception.AuthorizedError{Status: "UNAUTHORIZED", Message: check + " sudah ada"})
	}

	return member
}

func (repository *memberRepositoryImpl) FindMember(request model.GetFindUserRequest) (member model.GetFindUserResponse) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var username = bson.M{"_id": request.UserId}

	filter := username

	var test bson.M

	errone := repository.Collection.FindOne(ctx, filter).Decode(&test)
	if errone == nil {

		member.Id = test["_id"].(string)
		member.FullName = test["full_name"].(string)
		member.BirthDate = test["birth_date"].(string)
		member.Gender = test["gender"].(string)
		member.ProfilePicture = test["profile_picture"].(string)
		member.ObjectPicture = test["object_picture"].(string)
	}

	return member
}
