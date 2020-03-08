package controller

import (
    "github.com/gorilla/mux"
)

type Controller interface {
    RegisterFuncs(r *mux.Router)
}
