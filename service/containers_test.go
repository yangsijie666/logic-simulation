package service

import (
    "encoding/json"
    "fakeYun/model"
    "fmt"
    uuid "github.com/satori/go.uuid"
    "log"
    "testing"
)

func TestGetContainer(t *testing.T) {
    _, err := GetContainer(uuid.NewV4().String())
    if err.Error() != NO_SUCH_CONTAINER {
        t.FailNow()
    }
}

func TestCreateContainer(t *testing.T) {
    createContainer := model.Container{
        Name:  "test1",
        Image: "busybox",
        Networks: []model.Network{
            {
                NetworkId: "a4f79459-6659-4a7e-a05a-a42a03befd9e",
                FixedIps: []model.FixedIp{
                    {
                        SubnetId:  "5bac07f5-2762-4391-8ed8-029d4ed0eb80",
                        IpAddress: "1.0.0.25",
                    },
                },
            },
            {
                NetworkId: "d96dfe0d-5d3a-4215-8937-d589f447fc8f",
            },
        },
        Privileged: true,
    }
    returnContainer, err := CreateContainer(&createContainer)
    if err != nil {
        log.Fatalln(err)
    }
    js, err := json.Marshal(returnContainer)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("创建请求返回的容器对象：" + string(js))

    js, err = json.Marshal(createContainer)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("正在运行的容器对象: " + string(js))

    numberOfContainersInHost, err := getHostContainerNumberWithHostname(createContainer.Host)
    if err != nil {
        log.Fatalln(err)
    }

    if numberOfContainersInHost != 1 {
        t.Fatalf("The container number of %s should be 1, but here is %d.", createContainer.Host, numberOfContainersInHost)
    }
}

func getHostContainerNumberWithHostname(hostname string) (int, error) {
    for _, host := range hosts.Hosts {
        if host.Hostname == hostname {
            return host.TotalContainers, nil
        }
    }
    return -1, fmt.Errorf("%v", NO_SUCH_HOST)
}
