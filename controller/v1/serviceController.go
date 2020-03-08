package v1

import (
    "encoding/json"
    "fakeYun/service"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

const servicesPathPrefix = "/services"

type ServiceController struct{}

func (ServiceController) RegisterFuncs(r *mux.Router) {
    servicesRouter := r.PathPrefix(servicesPathPrefix).Subrouter()
    for _, f := range servicesFunctions {
        servicesRouter.HandleFunc(f.path, f.function).Methods(f.method)
    }
}

var servicesFunctions = []function{
    {
        path:     "",
        method:   http.MethodGet,
        function: httpListServices,
    },
    {
        path:     "",
        method:   http.MethodDelete,
        function: httpDeleteServices,
    },
}

func httpListServices(w http.ResponseWriter, r *http.Request) {
    // GET /v1/services
    result, err := listServices()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, result)
}

func httpDeleteServices(w http.ResponseWriter, r *http.Request) {
    // DELETE /v1/services?host={hostname}&binary={binary}
    hostname, binary, err := parseDeleteParams(r)
    if err != nil {
        if err.Error() == service.MISSING_PARAMETERS || err.Error() == service.WRONG_PARAMETERS {
            http.Error(w, err.Error(), http.StatusBadRequest)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    err = deleteService(hostname, binary)
    if err != nil {
        if err.Error() == service.NO_SUCH_SERVICE {
            http.Error(w, err.Error(), http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

func listServices() (string, error) {
    services := service.ListServices()
    js, err := json.Marshal(services)
    if err != nil {
        return "", err
    }
    return string(js), nil
}

func deleteService(host, binary string) error {
    return service.DeleteService(host, binary)
}

func parseDeleteParams(r *http.Request) (hostname, binary string, err error) {
    if err = r.ParseForm(); err != nil {
        return
    }
    hostnameFlag := false
    binaryFlag := false
    for k, v := range r.Form {
        if k == "host" {
            hostname = v[0]
            hostnameFlag = true
        } else if k == "binary" {
            binary = v[0]
            binaryFlag = true
        } else {
            err = fmt.Errorf("%v", service.WRONG_PARAMETERS)
            return
        }
    }
    if !hostnameFlag || !binaryFlag {
        err = fmt.Errorf("%v", service.MISSING_PARAMETERS)
        return
    }
    return
}
