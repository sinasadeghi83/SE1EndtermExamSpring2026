package transaction

import (
	"net/http"

	v1 "redbank/edge/connect/transaction/v1"
	// @ahum: imports
)

func RegisterService(mux *http.ServeMux) {
	v1.RegisterVersion(mux)
	// @ahum: versions
}
