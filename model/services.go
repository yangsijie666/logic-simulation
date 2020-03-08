package model

import (
    "encoding/json"
    "time"
)

type Services struct {
    Services    []Service   `json:"services"`
}

type Service struct {
    // TODO: 注意字段输出顺序（由于Marshal的嵌套结构体问题，目前无良好解决方案）
    Host        string      `json:"host"`
    Binary      string      `json:"binary"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
}

func (s Service) MarshalJSON() ([]byte, error) {
    type service Service
    tmpService := struct {
        service
        CreatedAt   string  `json:"created_at"`
        UpdatedAt   string  `json:"updated_at"`
    }{
        service: service(s),
        CreatedAt: s.CreatedAt.Format(TimeFormat),
        UpdatedAt: s.UpdatedAt.Format(TimeFormat),
    }
    return json.Marshal(tmpService)
}
