package v1

import (
    "io/ioutil"
    "log"
    "net/http"
    "testing"
)

func TestHostController_ListHosts(t *testing.T) {
    resp, err := http.Get("http://localhost:51227/v1/hosts")
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

func TestHostController_GetHost(t *testing.T) {
    testGetHost("07d93017-291f-4d70-aeb5-6d05ec48f076", t)
}

func testGetHost(uuid string, t *testing.T) {
    resp, err := http.Get("http://localhost:51227/v1/hosts/" + uuid)
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

func TestHostController_EnableAndDisableHost(t *testing.T) {
    req, err := http.NewRequest(http.MethodPut, "http://localhost:51227/v1/hosts/07d93017-291f-4d70-aeb5-6d05ec48f076/disable", nil)
    if err != nil {
        t.Fatal(err)
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    log.Println("已经disable")
    testGetHost("07d93017-291f-4d70-aeb5-6d05ec48f076", t)
    req, err = http.NewRequest(http.MethodPut, "http://localhost:51227/v1/hosts/07d93017-291f-4d70-aeb5-6d05ec48f076/enable", nil)
    if err != nil {
        t.Fatal(err)
    }
    _, err = http.DefaultClient.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    log.Println("已经enable")
    testGetHost("07d93017-291f-4d70-aeb5-6d05ec48f076", t)
}
