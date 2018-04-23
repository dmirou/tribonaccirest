package main

// Status codes
const (
	StatusOK                       = 200
	StatusBadRequest               = 400
	StatusNotFound                 = 404
	StatusMaxExecutionTimeExceeded = 460
	StatusInternalServerError      = 500
)

var statusText = map[int]string{
	StatusOK:                       "OK",
	StatusBadRequest:               "Bad Request",
	StatusNotFound:                 "Not Found",
	StatusMaxExecutionTimeExceeded: "Max Execution Time Exceeded",
	StatusInternalServerError:      "Internal Server Error",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
