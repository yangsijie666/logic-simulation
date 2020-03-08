package service

import (
    "encoding/json"
    "log"
    "testing"
)

func TestListHosts(t *testing.T) {
    hosts := ListHosts()
    js, err := json.Marshal(hosts)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
}

func TestGetHost(t *testing.T) {
    // 获取存在的host
    host, _ := GetHost("077822e4-9411-41ff-9bfc-8a10e11fbb8c")
    js, err := json.Marshal(host)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
    // 获取不存在的host
    _, err = GetHost("123456")
    if err.Error() != NO_SUCH_HOST {
        t.FailNow()
    }
}

func TestEnableHost(t *testing.T) {
    uuid := "077822e4-9411-41ff-9bfc-8a10e11fbb8c"
    _ = EnableHost(uuid)
    host, _ := GetHost(uuid)
    if host.Disabled != false {
        t.Fail()
    }
}

func TestDisableHost(t *testing.T) {
    uuid := "077822e4-9411-41ff-9bfc-8a10e11fbb8c"
    _ = DisableHost(uuid)
    host, _ := GetHost(uuid)
    if host.Disabled != true {
        t.Fail()
    }
}

func T()  {

}
