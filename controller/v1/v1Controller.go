package v1

import (
    "fakeYun/controller"
    "github.com/gorilla/mux"
    "net/http"
)

const v1PathPrefix = "/v1"

var v1Controllers = []controller.Controller{
    ServiceController{},
    HostController{},
    ContainerController{},
}

type V1Controller struct{}

func (V1Controller) RegisterFuncs(r *mux.Router) {
    subRouter := r.PathPrefix(v1PathPrefix).Subrouter()
    for _, c := range v1Controllers {
        c.RegisterFuncs(subRouter)
    }
}

type function struct {
    path     string
    method   string
    function func(w http.ResponseWriter, r *http.Request)
}
