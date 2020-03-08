package service

import (
    "encoding/json"
    "log"
    "testing"
)

func TestList(t *testing.T) {
    listServices()
}

func listServices()  {
    services := ListServices()
    js, err := json.Marshal(services)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
}

func hasHostWithHostname(hostname string) bool {
    for _, host := range hosts.Hosts{
        if host.Hostname == hostname {
            return true
        }
    }
    return false
}

func TestDelete(t *testing.T) {
    host := "compute9"
    binary := "yun-compute"
    // 删除成功的情况
    if err := DeleteService(host, binary); err == nil {
        listServices()
        // 查看对应的host是否删除成功
        if hasHostWithHostname(host) {
            t.Fatalf("对应host没被删除。")
        }
    } else {
        log.Fatalln(err)
    }
}

func TestDelete2(t *testing.T) {
    // 删除失败的情况
    if err := DeleteService("compute4", "yun-compute"); err == nil {
        listServices()
    } else {
        if err.Error() != NO_SUCH_SERVICE {
            t.FailNow()
        }
    }
}
