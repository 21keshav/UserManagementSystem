package util

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate counterfeiter -o fakes/fake_mongo_client.go --fake-name FakeMongoClient . MongoClient
type MongoClient interface {
	GetCollection(dbName, collectionName string) *mongo.Collection
	InsertData(dbName, collectionName string, data interface{}) (*mongo.InsertOneResult, error)
	FindObject(dbName,
		collectionName string, filter, result interface{}) error
	GetDatabase(dbName string) *mongo.Database
	FindUser(dbName,
		collectionName string, filter User) (User, error)
}

type MongoClientImpl struct {
	MongoClient *mongo.Client
	ctx         context.Context
}

func NewMongoClient(ctx context.Context, url string) MongoClient {
	client, _ := CreateClient(ctx, url)
	return &MongoClientImpl{
		client,
		ctx,
	}
}

func (mg *MongoClientImpl) GetCollection(dbName, collectionName string) *mongo.Collection {
	glog.Info("get-collection")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	collection := mg.GetDatabase(dbName).Collection(collectionName)
	return collection
}

func (mg *MongoClientImpl) GetDatabase(dbName string) *mongo.Database {
	glog.Info("get-database")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	database := mg.MongoClient.Database(dbName)
	return database
}

func (mg *MongoClientImpl) InsertData(dbName, collectionName string, data interface{}) (*mongo.InsertOneResult, error) {
	glog.Info("insert-data")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	collection := mg.GetCollection(dbName, collectionName)
	result, err := collection.InsertOne(mg.ctx, data)
	return result, err
}

func (mg *MongoClientImpl) FindObject(dbName,
	collectionName string, filter, result interface{}) error {
	glog.Info("find-object")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	collection := mg.GetCollection(dbName, collectionName)
	err := collection.FindOne(mg.ctx, filter).Decode(&result)
	return err
}

func (mg *MongoClientImpl) FindUser(dbName,
	collectionName string, filter User) (User, error) {
	glog.Info("find-user")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	var result User
	collection := mg.GetCollection(dbName, collectionName)
	resultval := collection.FindOne(mg.ctx, filter).Decode(&result)
	fmt.Printf("%v", resultval)
	return result, nil
}

func CreateClient(ctx context.Context, url string) (*mongo.Client, error) {
	glog.Info("create-mongo-client")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		glog.Error("check-mongo-connection", err)
		return nil, err
	}
	return client, nil
}
