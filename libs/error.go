package libs

import "net/http"

// InternalServerError is useful error handler
func InternalServerError(w http.ResponseWriter) {
	cause := recover()

	if cause == nil {
		return
	}

	if err, ok := cause.(error); ok {
		http.Error(w, err.Error(), 500)
	} else {
		http.Error(w, "Internal Server Error", 500)
	}
}
