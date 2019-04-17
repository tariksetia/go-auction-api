package auc_test_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User API", func() {
	//SignUp
	It("should create a new user", func() {
		postData := []byte(`{"username":"tariksetia","password":"1234"}`)
		req, _ := http.NewRequest("POST", "/v1/signup", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		Expect(response.Code).To(Equal(http.StatusCreated))
	})

	//Login
	It("should login with newly created user", func() {
		postData := []byte(`{"username":"tariksetia","password":"1234"}`)
		req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		var body map[string]interface{}
		r.ServeHTTP(response, req)
		json.NewDecoder(response.Body).Decode(&body)
		Expect(response.Code).To(Equal(200))
		Expect(body["token"]).NotTo(Equal(""))
	})

	//Login with non existing user
	It("should try to login with non-existing user", func() {
		postData := []byte(`{"username":"tarikhg","password":"1234"}`)
		req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		Expect(response.Code).To(Equal(http.StatusForbidden))

	})

	//Login with wrong password
	It("should login with wrong password", func() {
		postData := []byte(`{"username":"tariksetia","password":"12354"}`)
		req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		Expect(response.Code).To(Equal(http.StatusForbidden))
	})

	//try signup with existing user name
	It("should try signup with existing user name", func() {
		postData := []byte(`{"username":"tariksetia","password":"12354"}`)
		req, _ := http.NewRequest("POST", "/v1/signup", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		Expect(response.Code).To(Equal(http.StatusBadRequest))
	})

})
