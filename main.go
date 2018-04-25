// Package main contains RESTful service to calculate tribonacci number.
// To get tribonacci number with position "n" send request
// to url "<protocol>://<your-domain>:8080/tribonacci/{n}", where {n} is positive integer.
//
// For example, if you domain is "localhost" and requested url
// is http://localhost:8080/tribonacci/9 you will receive following
// result in json format {"code":200,"desc":"OK","data":{"n":9,"tribonacci":24}}
// where
// "code" is one of status codes,
// "desc" is text of status with specified code,
// "data" is result array, which contain "n" and "tribonacci" keys, if request
// successfully processed.
//
// Max calculation time of one tribonacci number is 4 seconds (4000 ms). If the calculation
// is more than 4 seconds you receive the result with code 460 and empty data.
// You can change max calculation time by changing the constant
// "maxTribonacciCalcTimeInMs" in app.go
package main

func main() {

	app := App{}

	app.Initialize()

	app.Run(":8080")
}
