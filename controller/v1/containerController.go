package v1

import (
    "encoding/json"
    "fakeYun/model"
    "fakeYun/service"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
)

const containersPathPrefix = "/containers"

type ContainerController struct{}

func (ContainerController) RegisterFuncs(r *mux.Router) {
    containerRouter := r.PathPrefix(containersPathPrefix).Subrouter()
    for _, f := range containerFunctions {
        containerRouter.HandleFunc(f.path, f.function).Methods(f.method)
    }
}

var containerFunctions = []function{
    {
        path:     "",
        method:   http.MethodGet,
        function: httpListContainers,
    },
    {
        path:     "/{uuid}",
        method:   http.MethodGet,
        function: httpGetContainer,
    },
    {
        path:     "",
        method:   http.MethodPost,
        function: httpCreateContainer,
    },
}

func httpListContainers(w http.ResponseWriter, r *http.Request) {
    // GET /v1/containers
    result, err := listContainers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, result)
}

func httpGetContainer(w http.ResponseWriter, r *http.Request) {
    // GET /v1/containers/{uuid}
    vars := mux.Vars(r)
    uuid, ok := vars["uuid"]
    if !ok {
        http.Error(w, service.MISSING_PARAMETERS+": uuid", http.StatusBadRequest)
        return
    }
    result, err := getContainer(uuid)
    if err != nil {
        if err.Error() == service.NO_SUCH_CONTAINER {
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

func httpCreateContainer(w http.ResponseWriter, r *http.Request) {
    // POST /v1/containers
    defer r.Body.Close()
    containerRequest, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    result, err := createContainer(containerRequest)
    if err != nil {
        if err.Error() != service.WRONG_PARAMETERS && err.Error() != service.MISSING_PARAMETERS {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        } else {
            http.Error(w, err.Error(), http.StatusBadRequest)
        }
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, result)
}

func listContainers() (string, error) {
    containers := service.ListContainers()
    js, err := json.Marshal(containers)
    if err != nil {
        return "", err
    }
    return string(js), nil
}

func getContainer(uuid string) (string, error) {
    container, err := service.GetContainer(uuid)
    if err != nil {
        return "", err
    }
    js, err := json.Marshal(container)
    if err != nil {
        return "", err
    }
    return string(js), nil
}

func createContainer(containerRequest []byte) (string, error) {
    container := new(model.Container)
    err := json.Unmarshal(containerRequest, container)
    if err != nil {
        return "", err
    }
    returnContainer, err := service.CreateContainer(container)
    if err != nil {
        return "", err
    }
    js, err := json.Marshal(returnContainer)
    if err != nil {
        return "", err
    }
    return string(js), err
}
