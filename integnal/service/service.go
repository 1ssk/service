package service

import "github.com/1ssk/service/pkg/api"

type Service struct {
	api.OrderServiceServer
}

func New() *Service {
	return &Service{}
}
