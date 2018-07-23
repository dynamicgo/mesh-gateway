package gateway

import (
	"fmt"
	"sync"

	"github.com/dynamicgo/slf4go"
)

type serviceMuxRegister struct {
	sync.RWMutex
	slf4go.Logger
	mux map[string]ServiceMux
}

var register = &serviceMuxRegister{
	mux:    make(map[string]ServiceMux),
	Logger: slf4go.Get("mesh-gateway-register"),
}

// RegisterServiceMux .
func RegisterServiceMux(name string, mux ServiceMux) {
	register.Lock()
	defer register.Unlock()

	if _, ok := register.mux[name]; ok {
		panic(fmt.Sprintf("duplicate register service mux %s", name))
	}

	register.mux[name] = mux
}

func mux(name string) (ServiceMux, bool) {
	register.RLock()
	defer register.RUnlock()

	m, ok := register.mux[name]

	return m, ok
}

func allmux() (result map[string]ServiceMux) {
	register.RLock()
	defer register.RUnlock()

	result = make(map[string]ServiceMux)
	for k, v := range register.mux {
		result[k] = v
	}

	return
}
