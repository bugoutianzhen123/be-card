package grpc

import (
	"context"
	v1 "github.com/asynccnu/be-api/gen/proto/card/v1"
	"github.com/asynccnu/be-card/domain"
	ser "github.com/asynccnu/be-card/service"
	"google.golang.org/grpc"
)

type CardService struct {
	v1.UnimplementedCardServer
	ser ser.Service
}

func NewCardGrpcService(ser ser.Service) *CardService {
	return &CardService{ser: ser}
}

func (s *CardService) Register(server grpc.ServiceRegistrar) {
	v1.RegisterCardServer(server, s)
}

func (s *CardService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.OperationResponse, error) {
	err := s.ser.CreateUser(ctx, domain.ServiceMsg{
		StudentId: req.StudentId,
		Key:       req.Key})
	if err != nil {
		return nil, err
	}
	return &v1.OperationResponse{}, nil
}

func (s *CardService) UpdateUserKey(ctx context.Context, req *v1.UpdateUserKeyRequest) (*v1.OperationResponse, error) {
	err := s.ser.UpdateUserKey(ctx, domain.ServiceMsg{
		StudentId: req.StudentId,
		Key:       req.Key,
	})
	if err != nil {
		return nil, err
	}
	return &v1.OperationResponse{}, nil
}

func (s *CardService) GetRecordOfConsumption(ctx context.Context, req *v1.GetRecordOfConsumptionRequest) (*v1.GetRecordOfConsumptionResponse, error) {
	records, err := s.ser.GetRecordOfConsumption(ctx, domain.ServiceMsg{
		StudentId: req.StudentId,
		Key:       req.Key,
		StartTime: req.StartTime,
		Type:      req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetRecordOfConsumptionResponse{
		Records: records,
	}, nil
}
