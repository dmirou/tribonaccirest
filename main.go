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
// "data" is result array, which contain "n" and "tribonachi" keys, if request
// successfully processed.
package main

func main() {

	app := App{}

	app.Initialize()

	app.Run(":8080")
}
