package gateway

import (
	"context"
	"net/http"
)

// ServiceHandler .
type ServiceHandler func(context.Context, http.ResponseWriter, *http.Request)

// ServiceMux .
type ServiceMux map[string]ServiceHandler

// Gateway gateway interface
type Gateway interface {
	http.Handler // mixin http handler interface
}

type gatewayImpl struct {
}
