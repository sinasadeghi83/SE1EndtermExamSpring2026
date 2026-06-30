package hello

import (
	"net/http"

	v1 "redbank/edge/connect/hello/v1"
	// @ahum: imports
)

func RegisterService(mux *http.ServeMux) {
	v1.RegisterVersion(mux)
	// @ahum: versions
}
