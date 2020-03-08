package service

import (
    "fakeYun/model"
    "fmt"
    "time"
)

var services = model.Services{
    Services: []model.Service{
        {
            Host:      "controller",
            Binary:    "yun-api",
            CreatedAt: time.Date(2019, 12, 22, 9, 32, 8, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 12, 0, time.Local),
        },
        {
            Host:      "controller",
            Binary:    "yun-conductor",
            CreatedAt: time.Date(2019, 12, 22, 9, 32, 13, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 16, 0, time.Local),
        },
        {
            Host:      "controller",
            Binary:    "yun-scheduler",
            CreatedAt: time.Date(2019, 12, 22, 9, 32, 15, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 18, 0, time.Local),
        },
        {
            Host:      "compute9",
            Binary:    "yun-compute",
            CreatedAt: time.Date(2019, 12, 22, 9, 33, 32, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 17, 0, time.Local),
        },
        {
            Host:      "compute11",
            Binary:    "yun-compute",
            CreatedAt: time.Date(2019, 12, 22, 9, 33, 59, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 22, 0, time.Local),
        },
        {
            Host:      "compute12",
            Binary:    "yun-compute",
            CreatedAt: time.Date(2019, 12, 22, 9, 34, 41, 0, time.Local),
            UpdatedAt: time.Date(2020, 1, 9, 18, 53, 18, 0, time.Local),
        },
    },
}

func ListServices() model.Services {
    return services
}

func DeleteService(host, binary string) error {
    done := false
    for i := 0; i < len(services.Services); i++ {
        if services.Services[i].Host == host && services.Services[i].Binary == binary {
            services.Services = append(services.Services[:i], services.Services[i+1:]...)
            done = true
            break
        }
    }
    if !done {
        return fmt.Errorf("%v", NO_SUCH_SERVICE)
    }
    _ = deleteHostWithHostname(host)
    return nil
}
