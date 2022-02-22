package serializer_test

import (
	"github.com/eswzy/grpc-learn/pb"
	"github.com/golang/protobuf/proto"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/eswzy/grpc-learn/sample"
	"github.com/eswzy/grpc-learn/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel() // Make all test running in parallel to detect racing condition

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
