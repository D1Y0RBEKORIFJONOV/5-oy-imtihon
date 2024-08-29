package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/entity"
)

type MongoDB struct {
	mongoClient *mongo.Client
	db          *mongo.Database
	collection  *mongo.Collection
}

func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	uri := "mongodb://" + cfg.DB.Host + cfg.DB.Port
	log.Printf("%s", uri)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		mongoClient: client,
		db:          client.Database(cfg.DB.Name),
		collection:  client.Database(cfg.DB.Name).Collection(cfg.DB.CollectionName),
	}, nil
}

func (m *MongoDB) SaveUser(ctx context.Context, req *entity.User) error {
	_, err := m.GetUser(ctx, &entity.GetUserReq{
		Field: "email",
		Value: req.Email,
	})
	objectID := primitive.NewObjectID()
	req.ID = objectID.Hex()
	if !errors.Is(err, entity.ErrorNotFound) {
		return entity.ErrorConflict
	}
	_, err = m.collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) GetUser(ctx context.Context, req *entity.GetUserReq) (*entity.User, error) {
	var value interface{} = req.Value
	if req.Field == "id" {

		value = req.Value
		req.Field = "_id"
	}
	var user entity.User
	err := m.collection.FindOne(ctx, bson.M{req.Field: value}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrorNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (m *MongoDB) GetAllUser(ctx context.Context, req *entity.GetAllUserReq) ([]*entity.User, error) {
	var value interface{}
	if req.Field != "" {
		value = req.Value
		if req.Field == "id" {
			value = req.Value
			req.Field = "_id"
		}
	}
	filter := bson.M{}
	if req.Field != "" {
		filter[req.Field] = value
	}

	findOptions := options.Find()
	if req.Page != 0 {
		findOptions.SetSkip(req.Page)
	}
	if req.Limit != 0 {
		findOptions.SetLimit(req.Limit)
	}

	cursor, err := m.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*entity.User
	for cursor.Next(ctx) {
		var user entity.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MongoDB) UpdateUser(ctx context.Context, req *entity.UpdateUserReq) error {
	filter := bson.M{"_id": req.UserID}
	update := bson.M{}
	if req.FirstName != "" {
		update["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		update["last_name"] = req.LastName
	}
	if len(update) == 0 {
		return nil
	}
	update["profile.updated_at"] = time.Now()
	updateBson := bson.M{"$set": update}
	_, err := m.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateUserPassword(ctx context.Context, newPassword, userID string) error {
	update := bson.M{}
	update["password"] = newPassword
	update["profile.updated_at"] = time.Now()
	updateBson := bson.M{"$set": update}
	filter := bson.M{"_id": userID}
	_, err := m.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateUserEmail(ctx context.Context, req *entity.UpdateEmailReq) error {
	update := bson.M{}
	update["email"] = req.NewEmail
	update["profile.updated_at"] = time.Now()
	updateBson := bson.M{"$set": update}
	filter := bson.M{"_id": req.UserID}

	_, err := m.collection.UpdateOne(ctx, filter, updateBson)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) DeleteUser(ctx context.Context, req *entity.DeleteUserReq) error {
	objId := req.UserID
	filter := bson.M{"_id": objId}
	if req.IsHardDelete {
		_, err := m.collection.DeleteOne(ctx, filter)
		if err != nil {
			return err
		}
		return nil
	}
	update := bson.M{"$set": bson.M{"profile.deleted_at": time.Now()}}
	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) IsDeleted(ctx context.Context, req entity.GetUserReq) bool {
	user, err := m.GetUser(ctx, &req)
	if err != nil {
		return false
	}
	return !user.Profile.DeletedAt.IsZero()
}
