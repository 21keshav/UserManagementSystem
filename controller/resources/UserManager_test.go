package resources_test

import (
	"context"
	"errors"

	"github.com/21keshav/UserManagementSystem/util/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/21keshav/UserManagementSystem/controller/resources"
	"github.com/21keshav/UserManagementSystem/util"
)

var _ = Describe("UserManager", func() {
	var (
		um              UserManager
		fakeMongoClient *fakes.FakeMongoClient
		dbName          string
		collectionName  string
	)
	BeforeEach(func() {
		ctx := context.TODO()
		fakeMongoClient = &fakes.FakeMongoClient{}
		um = NewUserManager(dbName, collectionName, fakeMongoClient, ctx)
	})

	Describe("GetUser", func() {
		var (
			expecteduser util.User
		)
		BeforeEach(func() {
			ctx := context.TODO()
			expecteduser = util.User{ID: "WOO", Firstname: "DAAM", Lastname: "POO"}
			fakeMongoClient.FindUserReturns(expecteduser, nil)
			/*
				  fakeMongoQuery.OneSTub = func(result interface{})  error {
					  r :=  result.(*util.User)
					  *r := expecteduser
				  }
			*/
			um = NewUserManager(dbName, collectionName, fakeMongoClient, ctx)
		})
		It("it runs sucessfull", func() {
			user, err := um.GetUser("36363")
			Expect(err).ToNot(HaveOccurred())
			Expect(user).To(Equal(expecteduser))
		})

		Context("Errors", func() {
			BeforeEach(func() {
				fakeMongoClient.FindUserReturns(util.User{}, errors.New("woo-ho"))
			})

			It("it errors", func() {
				_, err := um.GetUser("36363")
				Expect(err).To(HaveOccurred())
			})
		})

	})

	Describe("CreateUser", func() {
		var (
			user              util.User
			mongoInsertResult *mongo.InsertOneResult
			primitiveIDByte   primitive.ObjectID
		)
		BeforeEach(func() {
			ctx := context.TODO()
			str := "abc"
			for k, v := range []byte(str) {
				primitiveIDByte[k] = byte(v)
			}
			mongoInsertResult = &mongo.InsertOneResult{
				InsertedID: primitiveIDByte,
			}
			user = util.User{ID: "WOO", Firstname: "DAAM", Lastname: "POO"}
			fakeMongoClient.InsertDataReturns(mongoInsertResult, nil)
			/*
				  fakeMongoQuery.OneSTub = func(result interface{})  error {
					  r :=  result.(*util.User)
					  *r := expecteduser
				  }
			*/
			um = NewUserManager(dbName, collectionName, fakeMongoClient, ctx)
		})
		It("it runs sucessfull", func() {
			mongoResult, err := um.CreateUser(user)
			Expect(err).ToNot(HaveOccurred())
			Expect(mongoResult).To(Equal("ObjectID(\"616263000000000000000000\")"))
		})

		Context("Errors", func() {
			BeforeEach(func() {
				fakeMongoClient.InsertDataReturns(&mongo.InsertOneResult{}, errors.New("wodo"))
			})

			It("it errors", func() {
				_, err := um.CreateUser(user)
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
