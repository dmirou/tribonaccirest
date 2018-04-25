package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/dmirou/tribonacci"
	"github.com/gorilla/mux"
)

// App is application struct
type App struct {
	Router *mux.Router
}

// Initialize function initializes application.
func (app *App) Initialize() {

	app.Router = mux.NewRouter().StrictSlash(true)

	app.Router.HandleFunc("/tribonacci/{n}", handleTribonacciRequest)

	app.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

// Run function start application and http request handling.
func (app *App) Run(addr string) {

	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	outputInvalidResult(w, StatusNotFound)
}

// TribonacciResponse describes response structure.
type TribonacciResponse struct {
	Code int                 `json:"code"`
	Desc string              `json:"desc"`
	Data map[string]*big.Int `json:"data"`
}

// ToJSON converts TribonacciResponse to json string.
func (tribResponse *TribonacciResponse) ToJSON() string {

	jsonResult, _ := json.Marshal(tribResponse)

	return string(jsonResult)
}

var errMaxExecTimeExceeded = errors.New("max execution time exceeded")

const maxTribonacciCalcTimeInMs = 4000

func handleTribonacciRequest(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	nString := vars["n"]

	if nString == "" {
		outputInvalidResult(w, StatusBadRequest)
		return
	}

	nInt, err := strconv.Atoi(nString)

	if err != nil {
		outputInvalidResult(w, StatusBadRequest)
		return
	}

	if nInt <= 0 {
		outputInvalidResult(w, StatusBadRequest)
		return
	}

	tribonacci, err := calculateTribValue(nInt, maxTribonacciCalcTimeInMs)

	if err != nil {

		switch err {
		case errMaxExecTimeExceeded:
			outputInvalidResult(w, StatusMaxExecutionTimeExceeded)
			return
		default:
			outputInvalidResult(w, StatusInternalServerError)
			return
		}
	}

	outputSuccessResult(w, nInt, tribonacci)
}

func outputInvalidResult(w http.ResponseWriter, code int) {

	tribResponse := TribonacciResponse{Code: code, Desc: StatusText(code), Data: map[string]*big.Int{}}

	jsonString := tribResponse.ToJSON()

	fmt.Fprintf(w, jsonString)
}

func calculateTribValue(n int, maxExecTimeInMillisec int) (*big.Int, error) {

	quit := make(chan bool)
	done := make(chan bool)

	var tribonacciValue *big.Int
	var calculationError error

	go func() {
		tribonacciValue, calculationError = tribonacci.MatrixManaged(n, quit)
		close(done)
	}()

	select {
	case <-done:
		return tribonacciValue, calculationError
	case <-time.After(time.Duration(maxExecTimeInMillisec) * time.Millisecond):
		close(quit)
		return tribonacciValue, errMaxExecTimeExceeded
	}
}

func outputSuccessResult(w http.ResponseWriter, n int, tribonacci *big.Int) {

	successTribResponse := TribonacciResponse{
		Code: StatusOK,
		Desc: StatusText(StatusOK),
		Data: map[string]*big.Int{"n": big.NewInt(int64(n)), "tribonacci": tribonacci}}

	fmt.Fprintf(w, successTribResponse.ToJSON())
}
