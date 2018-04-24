package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var app App

func TestMain(m *testing.M) {

	app = App{}
	app.Initialize()

	code := m.Run()

	os.Exit(code)
}

func TestNIsAlphabeticString(t *testing.T) {

	request := buldGetRequest("/tribonacci/abc")

	response := executeRequest(request)

	expectedTribResponse := TribonacciResponse{
		Code: StatusBadRequest,
		Desc: StatusText(StatusBadRequest),
		Data: map[string]*big.Int{}}

	checkResponse(t, response, http.StatusOK, &expectedTribResponse)
}

func buldGetRequest(url string) *http.Request {

	request, _ := http.NewRequest("GET", url, nil)

	return request
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()

	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponse(t *testing.T, response *httptest.ResponseRecorder, expectedResponseCode int, expectedTribResponse *TribonacciResponse) {

	checkResponseCode(t, expectedResponseCode, response.Code)

	actualTribResponse := convertResponseBodyToTribResponse(response.Body.Bytes())

	checkJSONResultCode(t, expectedTribResponse.Code, actualTribResponse.Code)

	checkJSONResultDesc(t, expectedTribResponse.Desc, actualTribResponse.Desc)

	checkJSONResultData(t, expectedTribResponse.Data, actualTribResponse.Data)
}

func convertResponseBodyToTribResponse(responseBodyBytes []byte) TribonacciResponse {

	var tribResponce TribonacciResponse

	json.Unmarshal(responseBodyBytes, &tribResponce)

	return tribResponce
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkJSONResultCode(t *testing.T, expectedCode int, actualCode int) {

	if actualCode != expectedCode {
		t.Errorf("Expected the 'code' key of the response to be set to \"%d\". Got \"%d\"", expectedCode, actualCode)
	}
}

func checkJSONResultDesc(t *testing.T, expectedDesc string, actualDesc string) {

	if actualDesc != expectedDesc {
		t.Errorf("Expected the 'desc' key of the response to be set to \"%s\". Got \"%s\"", expectedDesc, actualDesc)
	}
}

func checkJSONResultData(t *testing.T, expectedData map[string]*big.Int, actualData map[string]*big.Int) {

	if !reflect.DeepEqual(actualData, expectedData) {
		t.Errorf("Expected the 'data' key of the response to be set to \"%s\". Got \"%s\"", expectedData, actualData)
	}
}

func TestNIsNegativeNumber(t *testing.T) {

	for i := 0; i < 1000; i++ {

		random := generateNegativeRandomNumber(-1*math.MaxInt32, -1)

		request := buldGetRequest(fmt.Sprintf("/tribonacci/%d", random))

		response := executeRequest(request)

		expectedTribResponse := TribonacciResponse{
			Code: StatusBadRequest,
			Desc: StatusText(StatusBadRequest),
			Data: map[string]*big.Int{}}

		checkResponse(t, response, http.StatusOK, &expectedTribResponse)
	}
}

func generateNegativeRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

var nToTribonacciValidMap = map[string]string{
	"1":    "0",
	"2":    "0",
	"3":    "1",
	"4":    "1",
	"5":    "2",
	"6":    "4",
	"7":    "7",
	"8":    "13",
	"9":    "24",
	"10":   "44",
	"20":   "19513",
	"50":   "1697490356184",
	"100":  "28992087708416717612934417",
	"1000": "443382579490226307661986241584270009256355236429858450381499235934108943134478901646797270328593836893366107162717822510963842586116043942479088674053663996392411782672993524690287662511197858910187264664163782145563472265666010074477859199789932765503984125240893",
}

func TestNIsValidNumber(t *testing.T) {

	for n, expectedTribonacci := range nToTribonacciValidMap {

		request := buldGetRequest(fmt.Sprintf("/tribonacci/%s", n))

		response := executeRequest(request)

		expectedTribResponse := TribonacciResponse{
			Code: StatusOK,
			Desc: StatusText(StatusOK),
			Data: map[string]*big.Int{"n": strToBigInt(n), "tribonacci": strToBigInt(expectedTribonacci)}}

		checkResponse(t, response, http.StatusOK, &expectedTribResponse)
	}
}

func strToBigInt(s string) *big.Int {
	result := new(big.Int)
	result, _ = result.SetString(s, 10)
	return result
}
