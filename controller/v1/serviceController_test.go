package v1

import (
    "io/ioutil"
    "log"
    "net/http"
    "testing"
)

func TestServiceController_ListServices(t *testing.T) {
    resp, err := http.Get("http://localhost:51227/v1/services")
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Fatal(err)
    }
    log.Println(string(body))
}

func TestServiceController_DeleteService(t *testing.T) {
    req, err := http.NewRequest("DELETE", "http://localhost:51227/v1/services?host=compute9&binary=yun-compute", nil)
    if err != nil {
        t.Fatal(err)
    }
    resp, _ := http.DefaultClient.Do(req)
    defer resp.Body.Close()
    log.Println(resp.StatusCode)
    if resp.StatusCode != http.StatusNoContent {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatal(err)
        }
        log.Println(string(body))
    }
}
