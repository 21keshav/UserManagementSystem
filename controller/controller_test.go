package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/labstack/echo"

	"github.com/21keshav/UserManagementSystem/controller"
	. "github.com/21keshav/UserManagementSystem/controller"
	"github.com/21keshav/UserManagementSystem/controller/resources/fakes"
	"github.com/21keshav/UserManagementSystem/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	xGlobalTransactionID = "x-Global-TransactionID"
)

var _ = Describe("Controller", func() {
	var (
		request                    *http.Request
		ctx                        echo.Context
		e                          *echo.Echo
		controller                 controller.Controller
		userManager                *fakes.FakeUserManager
		recorder                   *httptest.ResponseRecorder
		receivedUser, expecteduser util.User
	)
	BeforeEach(func() {
		request = new(http.Request)
		request.Header = make(http.Header)
		request.Header.Add(xGlobalTransactionID, "foo")
		e = echo.New()
		ctx = e.NewContext(request, nil)
		ctx.Request().Header.Add("Authorization", "token")
		userManager = &fakes.FakeUserManager{}
		controller = NewController(userManager)
		recorder = httptest.NewRecorder()
	})
	Describe("GetUser", func() {
		BeforeEach(func() {
			ctx.SetParamNames("id")
			ctx.SetParamValues("dxxv")
			requestURL, _ := url.Parse("doesnotmatter?id=qqq")
			ctx.Response().Writer = recorder
			ctx.Request().URL = requestURL
			expecteduser = util.User{ID: "WOO", Firstname: "DAAM", Lastname: "POO"}
			userManager.GetUserReturns(expecteduser, nil)
		})
		It("it runs sucessfull", func() {

			err := controller.GetUser(ctx)
			Expect(err).ToNot(HaveOccurred())

			Expect(recorder.Code).To(Equal(http.StatusOK))
			readBytes := make([]byte, 1000)
			_, err = recorder.Body.Read(readBytes)
			readBytes = bytes.Trim(readBytes, "\x00")
			Expect(err).ToNot(HaveOccurred())
			_ = json.Unmarshal(readBytes, &receivedUser)
			Expect(receivedUser).To(Equal(expecteduser))
		})
	})

	Describe("CreateUser", func() {
		var (
			recorder2      *httptest.ResponseRecorder
			expectedString string
		)
		BeforeEach(func() {
			recorder2 = httptest.NewRecorder()
			requestURL, _ := url.Parse("doesnotmatter")
			ctx.Response().Writer = recorder2
			ctx.Request().URL = requestURL
			expecteduser = util.User{ID: "WOO", Firstname: "DAAM", Lastname: "POO"}
			bytes1, _ := json.Marshal(expecteduser)
			r := ioutil.NopCloser(bytes.NewReader(bytes1))
			ctx.Request().Body = r
			userManager.CreateUserReturns("WODO", nil)
		})
		It("it runs sucessfull", func() {

			err := controller.CreateUser(ctx)
			Expect(err).ToNot(HaveOccurred())
			Expect(recorder2.Code).To(Equal(http.StatusCreated))
			readBytes := make([]byte, 1000)
			_, err = recorder2.Body.Read(readBytes)
			readBytes = bytes.Trim(readBytes, "\x00")
			_ = json.Unmarshal(readBytes, &expectedString)
			Expect(string(expectedString)).To(Equal("WODO"))

		})
	})
})
