package account

import (
	"net/http"

	v1 "redbank/edge/connect/account/v1"
	// @ahum: imports
)

func RegisterService(mux *http.ServeMux) {
	v1.RegisterVersion(mux)
	// @ahum: versions
}
