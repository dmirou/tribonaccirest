// Package main contains RESTful service to calculate tribonacci number.
// Valid url path is "/tribonacci/{n}" where {n} is positive integer.
// For example, if you send GET request with url path /tribonacci/9
// you receive result {"code":200,"desc":"OK","data":{"n":9,"tribonacci":24}}
package main

func main() {

	app := App{}

	app.Initialize()

	app.Run(":8080")
}
