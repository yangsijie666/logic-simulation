package v1

import (
    "encoding/json"
    "fakeYun/service"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

const hostsPathPrefix = "/hosts"

type HostController struct{}

func (HostController) RegisterFuncs(r *mux.Router) {
    hostsRouter := r.PathPrefix(hostsPathPrefix).Subrouter()
    for _, f := range hostsFunctions {
        hostsRouter.HandleFunc(f.path, f.function).Methods(f.method)
    }
}

var hostsFunctions = []function{
    {
        path:     "",
        method:   http.MethodGet,
        function: httpListHosts,
    },
    {
        path:     "/{uuid}",
        method:   http.MethodGet,
        function: httpGetHosts,
    },
    {
        path:     "/{uuid}/enable",
        method:   http.MethodPut,
        function: httpEnableHosts,
    }, {
        path:     "/{uuid}/disable",
        method:   http.MethodPut,
        function: httpDisableHosts,
    },
}

func httpListHosts(w http.ResponseWriter, r *http.Request) {
    // GET /v1/hosts
    result, err := listHosts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, result)
}

func httpGetHosts(w http.ResponseWriter, r *http.Request) {
    // GET /v1/hosts/{uuid}
    vars := mux.Vars(r)
    uuid, ok := vars["uuid"]
    if !ok {
        http.Error(w, service.MISSING_PARAMETERS+": uuid", http.StatusBadRequest)
        return
    }
    result, err := getHost(uuid)
    if err != nil {
        if err.Error() == service.NO_SUCH_HOST {
            http.Error(w, err.Error(), http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, result)
}

func httpEnableHosts(w http.ResponseWriter, r *http.Request) {
    // PUT /v1/hosts/{uuid}/enable
    vars := mux.Vars(r)
    uuid, ok := vars["uuid"]
    if !ok {
        http.Error(w, service.MISSING_PARAMETERS+": uuid", http.StatusBadRequest)
        return
    }
    err := enableHost(uuid)
    if err != nil {
        if err.Error() == service.NO_SUCH_HOST {
            http.Error(w, err.Error(), http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusOK)
}

func httpDisableHosts(w http.ResponseWriter, r *http.Request) {
    // PUT /v1/hosts/{uuid}/disable
    vars := mux.Vars(r)
    uuid, ok := vars["uuid"]
    if !ok {
        http.Error(w, service.MISSING_PARAMETERS+": uuid", http.StatusBadRequest)
        return
    }
    err := disableHost(uuid)
    if err != nil {
        if err.Error() == service.NO_SUCH_HOST {
            http.Error(w, err.Error(), http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusOK)
}

func listHosts() (string, error) {
    hosts := service.ListHosts()
    js, err := json.Marshal(hosts)
    if err != nil {
        return "", err
    }
    return string(js), nil
}

func getHost(uuid string) (string, error) {
    host, err := service.GetHost(uuid)
    if err != nil {
        return "", err
    }
    js, err := json.Marshal(host)
    if err != nil {
        return "", err
    }
    return string(js), nil
}

func enableHost(uuid string) error {
    return service.EnableHost(uuid)
}

func disableHost(uuid string) error {
    return service.DisableHost(uuid)
}
