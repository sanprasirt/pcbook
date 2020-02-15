package serializer_test

import (
	"testing"

	"github.com/sanprasirt/pcbook/pb"
	"github.com/sanprasirt/pcbook/sample"
	"github.com/sanprasirt/pcbook/serializer"
	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"
	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

func TestReadProtobufFromBinaryFile(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	laptop2 := &pb.Laptop{}
	err := serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

}
