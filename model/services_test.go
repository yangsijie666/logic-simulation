package model

import (
    "encoding/json"
    "log"
    "testing"
    "time"
)

func TestService_MarshalJSON(t *testing.T) {
    service := Service{
        Host:      "compute1",
        Binary:    "yun-api",
        CreatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
        UpdatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
    }
    js, err := json.Marshal(service)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
}
