package service

import (
    "fakeYun/model"
    "fmt"
    "time"
)

var hosts = model.Hosts{
    Hosts: []model.Host{
        {
            Uuid:            "07d93017-291f-4d70-aeb5-6d05ec48f076",
            Hostname:        "compute9",
            MemTotal:        32707,
            MemFree:         32707,
            DiskTotal:       107,
            DiskFree:        107,
            CpusTotal:       24,
            CpusFree:        24,
            Disabled:        false,
            TotalContainers: 0,
            NumaTopology:    "[{\"cpuset\": [0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22], \"mem_free\": 16021, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 16021}, {\"cpuset\": [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23], \"mem_free\": 16122, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 16122}]",
            CreatedAt:       time.Date(2019, 12, 22, 9, 33, 32, 0, time.Local),
            UpdatedAt:       time.Date(2020, 1, 9, 18, 53, 17, 0, time.Local),
        },
        {
            Uuid:            "077822e4-9411-41ff-9bfc-8a10e11fbb8c",
            Hostname:        "compute11",
            MemTotal:        16323,
            MemFree:         16323,
            DiskTotal:       107,
            DiskFree:        107,
            CpusTotal:       24,
            CpusFree:        24,
            Disabled:        false,
            TotalContainers: 0,
            NumaTopology:    "[{\"cpuset\": [0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22], \"mem_free\": 7838, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 7838}, {\"cpuset\": [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23], \"mem_free\": 7935, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 7935}]",
            CreatedAt:       time.Date(2019, 12, 22, 9, 33, 59, 0, time.Local),
            UpdatedAt:       time.Date(2020, 1, 9, 18, 53, 22, 0, time.Local),
        },
        {
            Uuid:            "003cabe6-25b6-422f-9eee-6e1d7f14de95",
            Hostname:        "compute12",
            MemTotal:        16323,
            MemFree:         16323,
            DiskTotal:       107,
            DiskFree:        107,
            CpusTotal:       24,
            CpusFree:        24,
            Disabled:        false,
            TotalContainers: 0,
            NumaTopology:    "[{\"cpuset\": [0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22], \"mem_free\": 7838, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 7838}, {\"cpuset\": [1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23], \"mem_free\": 7935, \"pinned_cpus\": [], \"id\": 0, \"mem_total\": 7935}]",
            CreatedAt:       time.Date(2019, 12, 22, 9, 34, 41, 0, time.Local),
            UpdatedAt:       time.Date(2020, 1, 9, 18, 53, 18, 0, time.Local),
        },
    },
}

func ListHosts() model.Hosts {
    return hosts
}

func GetHost(uuid string) (model.Host, error) {
    for _, host := range hosts.Hosts {
        if host.Uuid == uuid {
            return host, nil
        }
    }
    return model.Host{}, fmt.Errorf("%v", NO_SUCH_HOST)
}

func EnableHost(uuid string) error {
    return enableOrDisableHost(uuid, false)
}

func DisableHost(uuid string) error {
    return enableOrDisableHost(uuid, true)
}

func enableOrDisableHost(uuid string, disable bool) error {
    for index, host := range hosts.Hosts {
        if host.Uuid == uuid {
            hosts.Hosts[index].Disabled = disable
            return nil
        }
    }
    return fmt.Errorf("%v", NO_SUCH_HOST)
}

func deleteHostWithHostname(hostname string) error {
    for i := 0; i < len(hosts.Hosts); i++ {
        if hosts.Hosts[i].Hostname == hostname {
            hosts.Hosts = append(hosts.Hosts[:i], hosts.Hosts[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("%v", NO_SUCH_HOST)
}

func ConsumeResources(container *model.Container) error {
    hostname := container.Host
    for index, host := range hosts.Hosts {
        if host.Hostname == hostname {
            return hosts.Hosts[index].AddOneContainer()
        }
    }
    return fmt.Errorf("%v", NO_SUCH_HOST)
}
