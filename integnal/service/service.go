package service

import "github.com/test/pkg/api"

type Service struct {
	api.OrderServiceServer
}

func New() *Service {
	return &Service{}
}
