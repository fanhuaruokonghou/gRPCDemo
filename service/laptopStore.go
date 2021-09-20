package service

import (
	"errors"
	"fmt"
	"gRPCDemo/pb"
	"github.com/jinzhu/copier"
	"sync"
)

var ErrAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	data  map[string]*pb.Laptop
}

func (i *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	laptop := i.data[id]

	if laptop == nil {
		return nil, nil
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %v", err)
	}
	return other, nil
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (i *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, ok := i.data[laptop.Id]; ok {
		return ErrAlreadyExists
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}
	i.data[laptop.Id] = laptop
	return nil
}
