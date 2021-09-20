package service

import (
	"context"
	"errors"
	"gRPCDemo/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer(Store LaptopStore) *LaptopServer {
	return &LaptopServer{Store}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, request *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "laptop generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	time.Sleep(6 * time.Second)
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	err := s.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}
	log.Printf("saved laptop with id: %s", laptop.Id)
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}
func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
