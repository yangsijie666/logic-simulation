package model

import (
    "encoding/json"
    "time"
)

type Hosts struct {
    Hosts []Host `json:"hosts"`
}

type Host struct {
    Uuid            string    `json:"uuid"`
    Hostname        string    `json:"hostname"`
    MemTotal        int32     `json:"mem_total"`
    MemFree         int32     `json:"mem_free"`
    DiskTotal       int32     `json:"disk_total"`
    DiskFree        int32     `json:"disk_free"`
    CpusTotal       int8      `json:"cpus_total"`
    CpusFree        int8      `json:"cpus_free"`
    Disabled        bool      `json:"disabled,string"`
    TotalContainers int       `json:"total_containers"`
    NumaTopology    string    `json:"numa_topology"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (h *Host) AddOneContainer() error {
    h.TotalContainers += 1
    return nil
}

func (h Host) MarshalJSON() ([]byte, error) {
    type host Host

    tmpHost := struct {
        host
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
    }{
        host:      host(h),
        CreatedAt: h.CreatedAt.Format(TimeFormat),
        UpdatedAt: h.UpdatedAt.Format(TimeFormat),
    }
    return json.Marshal(tmpHost)
}
