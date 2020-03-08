package model

import (
    "encoding/json"
    "log"
    "testing"
    "time"
)

func TestMarshal(t *testing.T)  {
    container := Container{
        Name:      "ysj",
        CreatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
    }
    js, err := json.Marshal(&container)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(js))
}

func TestUnmarshal(t *testing.T)  {
    container := Container{
        Name:      "ysj",
        CreatedAt: time.Date(1995, 2, 1, 6, 6, 18, 0, time.Local),
    }
    js, err := json.Marshal(container)
    if err != nil {
        t.Fatal(err)
    }
    newContainer := new(Container)
    err = json.Unmarshal(js, newContainer)
    if err != nil {
        t.Fatal(err)
    }
    log.Printf("%v", newContainer)
}
