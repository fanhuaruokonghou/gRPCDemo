package service_test

import (
	"context"
	"fmt"
	"gRPCDemo/pb"
	"gRPCDemo/sample"
	"gRPCDemo/serializer"
	"gRPCDemo/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)
	//serverAddress = "localhost:9999"
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	fmt.Println(err)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	requireSameLaptop(t, laptop, other)

}

func requireSameLaptop(t *testing.T, laptop *pb.Laptop, other *pb.Laptop) {
	jsonLaptop, err := serializer.ProtobufToJSON(laptop)
	require.NoError(t, err)

	jsonOther, err := serializer.ProtobufToJSON(other)
	require.NoError(t, err)

	require.Equal(t, jsonOther, jsonLaptop)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", "localhost:9999")
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	fmt.Println(listener.Addr().String())
	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}
