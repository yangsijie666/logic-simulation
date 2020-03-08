package main

import (
    "fakeYun/controller"
    "fakeYun/controller/defaultController"
    "fakeYun/controller/v1"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

var controllers = []controller.Controller{
    defaultController.DefaultController{},
    v1.V1Controller{},
}

func registerFuncs(r *mux.Router)  {
    router := r.PathPrefix("/").Subrouter()
    for _, c := range controllers {
        c.RegisterFuncs(router)
    }
}

func main() {
    muxRouter := mux.NewRouter()
    registerFuncs(muxRouter)
    log.Fatal(http.ListenAndServe(":51227", muxRouter))
}
