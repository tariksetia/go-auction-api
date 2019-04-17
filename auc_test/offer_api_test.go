package auc_test_test

import (
	e "auction/pkg/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Offer-API-TEST", func() {
	var token string

	BeforeEach(func() {
		//Create A user
		postData := []byte(`{"username":"tariksetia","password":"1234"}`)
		req, _ := http.NewRequest("POST", "/v1/signup", bytes.NewBuffer(postData))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)

		//login with created user user
		postData = []byte(`{"username":"tariksetia","password":"1234"}`)
		req, _ = http.NewRequest("POST", "/v1/login", bytes.NewBuffer(postData))
		response = httptest.NewRecorder()
		r.ServeHTTP(response, req)

		//Parse Response body and get the token
		var body map[string]interface{}
		json.NewDecoder(response.Body).Decode(&body)
		token = body["token"].(string)
	})

	It("should create an offer", func() {
		postData := []byte(`{
			"title":"THis is a 1",
			"bid_price":1.54
		}`)
		req, _ := http.NewRequest("POST", "/v1/signup", bytes.NewBuffer(postData))
		req.Header.Add("Authorization", "Bearer "+token)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, req)
		var ofr e.Offer
		json.NewDecoder(response.Body).Decode(&ofr)
		fmt.Println(ofr)
	})

})
