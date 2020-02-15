package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtobufToJSON convert protocal buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	return marshaler.MarshalToString(message)
}

// JSONToProtobufMessage convert JSON string to protocal buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return jsonpb.UnmarshalString(data, message)
}
