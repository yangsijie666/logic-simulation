package model

import (
    "encoding/json"
    "log"
    "testing"
    "time"
)

func TestHost_MarshalJSON(t *testing.T) {
    host := Host{
        Uuid:      "52d5eb68-00a0-4689-87da-cd168e9dc61d",
        Hostname:  "compute1",
        MemTotal:  7888260,
        DiskTotal: 2032,
        CpusTotal: 32,
        Disabled:  false,
        CreatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
        UpdatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
    }
    js, err := json.Marshal(host)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
}
