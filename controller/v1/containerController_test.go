package v1

import (
    "encoding/json"
    "fakeYun/model"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "testing"
)

func TestContainerController_CreateAndGet(t *testing.T) {
    createContainerString := `{"name":"test1","image":"busybox","networks":[{"network_id":"a4f79459-6659-4a7e-a05a-a42a03befd9e","fixed_ips":[{"subnet_id":"5bac07f5-2762-4391-8ed8-029d4ed0eb80","ip_address":"1.0.0.25"}]},{"network_id":"d96dfe0d-5d3a-4215-8937-d589f447fc8f"}],"privileged":true}`
    createContainerReader := strings.NewReader(createContainerString)
    req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:51227/v1/containers", createContainerReader)
    if err != nil {
        log.Fatal(err)
    }
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
    c := new(model.Container)
    err = json.Unmarshal(data, c)
    if err != nil {
        log.Fatal(err)
    }
    resp.Body.Close()

    resp, err = http.Get("http://127.0.0.1:51227/v1/containers/" + c.Uuid)
    if err != nil {
        log.Fatal(err)
    }
    data, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(string(data))
    resp.Body.Close()
}
