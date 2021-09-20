package serializer

import (
	"fmt"
	"gRPCDemo/pb"
	"gRPCDemo/sample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	if err != nil {
		fmt.Errorf("test ReadProtobufToBinaryFile failed: %w", err)
	}
	if proto.Equal(laptop1, laptop2) {
		fmt.Println("equal!")
	}

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
