package defaultController

import (
    "fakeYun/controller"
    v1 "fakeYun/controller/v1"
    "github.com/gorilla/mux"
)

var defaultControllers = []controller.Controller{
    v1.ServiceController{},
    v1.HostController{},
    v1.ContainerController{},
}

type DefaultController struct {}

func (DefaultController) RegisterFuncs(r *mux.Router)  {
    for _, c := range defaultControllers {
        c.RegisterFuncs(r)
    }
}
