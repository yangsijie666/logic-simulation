package service

import (
    "fakeYun/model"
    "fmt"
    uuid "github.com/satori/go.uuid"
    "time"
)

const DEFAULT_IP_ADDRESS = "2.0.0.12"

var containers = model.Containers{
    Containers: []model.Container{},
}

func CreateContainer(container *model.Container) (model.Container, error) {
    initContainer(container)
    returnContainer := *container

    // 包装Networks中的值进入Addresses
    for _, network := range container.Networks {
        tmpAddress := model.Address{
            NetworkId: network.NetworkId,
            Details: []model.AddressDetail{},
        }
        if network.FixedIps != nil {
            for _, fixedIp := range network.FixedIps {
                tmpAddress.Details = append(tmpAddress.Details, model.AddressDetail{
                    SubnetId:  fixedIp.SubnetId,
                    IpAddress: fixedIp.IpAddress,
                    PortId:    uuid.NewV4().String(),
                })
            }
        } else {
            tmpAddress.Details = append(tmpAddress.Details, model.AddressDetail{
                SubnetId:  uuid.NewV4().String(),
                IpAddress: DEFAULT_IP_ADDRESS,
                PortId:    uuid.NewV4().String(),
            })
        }
        container.Addresses = append(container.Addresses, tmpAddress)
    }
    container.Networks = nil
    // 更新容器状态
    container.Status = "Running"
    // 更新所在主机
    container.Host = "compute9"
    // 更新容器id
    container.ContainerId = "1d1788b4560d04530584d9cb6787843301b301ee65e714f80defbc73055bf9f0"
    // 更新容器命令
    container.Command = "sh"
    // 更新NUMA拓扑
    container.NumaTopology = "[]"
    // 更新容器时间
    container.UpdatedAt = time.Date(2020, 1, 9, 18, 54, 55, 0, time.Local)

    containers.Containers = append(containers.Containers, *container)
    // 消耗对应计算节点的资源
    err := ConsumeResources(container)
    if err != nil {
        return returnContainer, err
    }

    return returnContainer, nil
}

func GetContainer(uuid string) (model.Container, error) {
    for _, container := range containers.Containers {
        if container.Uuid == uuid {
            return container, nil
        }
    }
    return model.Container{}, fmt.Errorf("%v", NO_SUCH_CONTAINER)
}

func ListContainers() model.Containers {
    return containers
}

func initContainer(c *model.Container)  {
    // 设置默认值以及检查
    c.Uuid = uuid.NewV4().String()

    c.Status = "Creating"

    _, err := uuid.FromString(c.Image)
    if err != nil {
        c.Image = uuid.NewV4().String()
    }

    if c.Vcpus != 0 || c.Ram != 0 || c.Disk != 0 {
        if c.Vcpus == 0 {
            c.Vcpus = 1
        }
        if c.Ram == 0 {
            c.Ram = 256
        }
        if c.Disk == 0 {
            c.Disk = 10
        }
    }

    createTime := time.Date(2020, 1, 9, 18, 54, 52, 0, time.Local)
    c.CreatedAt = createTime
    c.UpdatedAt = createTime
}
