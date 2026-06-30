package v1

import (
	"net/http"

	"redbank/edge/connect/gen/transaction/v1/transactionv1connect"
	// @ahum: imports
)

func RegisterVersion(mux *http.ServeMux) {
	edge := NewEdge()
	path, handler := transactionv1connect.NewServiceHandler(edge)
	mux.Handle(path, handler)
	// @ahum: edges
}
