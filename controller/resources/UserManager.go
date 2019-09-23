package resources

import (
	"context"

	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/21keshav/UserManagementSystem/util"
)

//go:generate counterfeiter -o fakes/fake_user_manager.go --fake-name FakeUserManager . UserManager
type UserManager interface {
	CreateUser(user util.User) (string, error)
	GetUser(id string) (util.User, error)
}

type UserManagerImpl struct {
	MongoClient util.MongoClient
	DbDetails   util.DBDetails
	ctx         context.Context
}

func NewUserManager(dbName, collectionName string, mongoClient util.MongoClient, ctx context.Context) UserManager {
	dbdetails := util.DBDetails{
		DbName:         dbName,
		CollectionName: collectionName,
	}
	return &UserManagerImpl{
		mongoClient,
		dbdetails,
		ctx,
	}
}

func (um *UserManagerImpl) CreateUser(user util.User) (string, error) {
	glog.Info("um-create-user")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	result, err := um.MongoClient.InsertData(um.DbDetails.DbName,
		um.DbDetails.CollectionName, user)
	if err != nil {
		glog.Error("mongo error inserting object", err)
		return "", err
	}
	id := (result.InsertedID.(primitive.ObjectID)).String()

	return id, nil
}

func (um *UserManagerImpl) GetUser(id string) (util.User, error) {
	glog.Info("um-get-user")
	glog.InfoDepth(1, "started")
	defer glog.InfoDepth(1, "completed")
	user, err := um.MongoClient.FindUser(um.DbDetails.DbName,
		um.DbDetails.CollectionName, util.User{ID: id})
	if err != nil {
		glog.Error("mongo error finding object", err)
		return user, err
	}
	return user, nil
}
