package strutlog

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"strutlog/generated"
	"testing"
)

func TestProtoMarshallerMarshalValidProtobuf(t *testing.T) {
	protoMarshaller := ProtoMarshaller()

	type testCase struct {
		name  string
		input Message
	}

	testCases := []testCase{
		{
			name: "Marshals message to proto",
			input: defaultMessage{
				fields: []Field{
					{
						Key:   "test",
						Value: stringValue{"hello there testing out this"},
					},
				},
				level: InfoLevel,
			},
		},
		{
			name: "Marshals another message to proto",
			input: defaultMessage{
				fields: []Field{
					{
						Key: "map",
						Value: mapValue{map[Value]Value{
							stringValue{"test"}:  stringValue{"hello"},
							stringValue{"found"}: boolValue{true},
						}},
					},
					{
						Key:   "array",
						Value: arrayValue{[]Value{intValue{1}, runeValue{7}}},
					},
				},
				level: TraceLevel,
			},
		},
		{
			name:  "Deals with nil",
			input: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bytes := protoMarshaller(test.input)

			assert.NotNil(t, bytes)

			protoMessage := &generated.Message{}
			err := proto.Unmarshal(bytes, protoMessage)

			assert.NoError(t, err)
		})
	}
}
