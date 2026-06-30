package v1

import (
	"net/http"

	"redbank/edge/connect/gen/account/v1/accountv1connect"
	// @ahum: imports
)

func RegisterVersion(mux *http.ServeMux) {
	edge := NewEdge()
	path, handler := accountv1connect.NewServiceHandler(edge)
	mux.Handle(path, handler)
	// @ahum: edges
}
