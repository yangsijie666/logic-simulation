package model

import (
    "encoding/json"
    "time"
)

const TimeFormat = "2006-01-02 15:04:05"

type Containers struct {
    Containers []Container `json:"containers"`
}

type Container struct {
    Name        string    `json:"name"`
    Image       string    `json:"image"`
    Addresses   []Address `json:"addresses,omitempty"`
    Command     string    `json:"command,omitempty"`
    Status      string    `json:"status"`
    Uuid        string    `json:"uuid"`
    ContainerId string    `json:"container_id,omitempty"`
    Host        string    `json:"host,omitempty"`
    CreatedAt   time.Time `json:"created_at,omitempty"`
    UpdatedAt   time.Time `json:"updated_at,omitempty"`

    Networks     []Network `json:"networks,omitempty"`
    Privileged   bool      `json:"privileged"`
    Vcpus        int       `json:"vcpus"`
    Ram          int       `json:"ram"`
    Disk         int       `json:"disk"`
    Reason       string    `json:"reason,omitempty"`
    NumaTopology string    `json:"numa_topology,omitempty"`
}

func (c Container) MarshalJSON() ([]byte, error) {
    type container Container
    tmpContainer := struct {
        container
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
    }{
        (container)(c),
        c.CreatedAt.Format(TimeFormat),
        c.UpdatedAt.Format(TimeFormat),
    }
    return json.Marshal(tmpContainer)
}

func (c *Container) UnmarshalJSON(in []byte) error {
    // 此时的*c是零值，即每个字段都是零值，还未被填充
    // 转换时间格式
    type container Container
    tmpContainer := &struct {
        *container
        CreatedAt string `json:"created_at"`
        UpdatedAt string `json:"updated_at"`
    }{
        container: (*container)(c),
        CreatedAt: c.CreatedAt.Format(TimeFormat),
        UpdatedAt: c.UpdatedAt.Format(TimeFormat),
    }
    err := json.Unmarshal(in, tmpContainer)
    if err != nil {
        return err
    }
    c.CreatedAt, err = time.Parse(TimeFormat, tmpContainer.CreatedAt)
    if err != nil {
        return err
    }
    c.UpdatedAt, err = time.Parse(TimeFormat, tmpContainer.UpdatedAt)
    if err != nil {
        return err
    }

    return nil
}

type Network struct {
    NetworkId string    `json:"network_id"`
    FixedIps   []FixedIp `json:"fixed_ips,omitempty"`
}

type FixedIp struct {
    SubnetId  string `json:"subnet_id"`
    IpAddress string `json:"ip_address"`
}

type Address struct {
    NetworkId string          `json:"network_id"`
    Details   []AddressDetail `json:"details,omitempty"`
}

type AddressDetail struct {
    SubnetId  string `json:"subnet_id,omitempty"`
    IpAddress string `json:"ip_address,omitempty"`
    PortId    string `json:"port_id,omitempty"`
}
