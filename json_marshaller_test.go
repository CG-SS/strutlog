package strutlog

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMarshallerMarshalsValidJson(t *testing.T) {
	jsonMarshallerFn := JsonMarshaller(-1, NoExponentFloatFormat)

	type testCase struct {
		name  string
		input Message
	}

	testCases := []testCase{
		{
			name: "Marshal message to valid JSON",
			input: defaultMessage{
				fields: []Field{
					{
						Key:   "test",
						Value: boolValue{true},
					},
					{
						Key:   "message",
						Value: stringValue{"testing message"},
					},
					{
						Key:   "count",
						Value: intValue{12},
					},
					{
						Key:   "result32",
						Value: float32Value{float32(156.77)},
					},
					{
						Key:   "result64",
						Value: float32Value{288.77},
					},
					{
						Key:   "complex",
						Value: complex128Value{complex128(86.4)},
					},
					{
						Key: "array",
						Value: arrayValue{[]Value{
							boolValue{true},
							stringValue{"arrayString"},
							int8Value{int8(6)},
						}},
					},
					{
						Key:   "rune",
						Value: runeValue{rune(2)},
					},
					{
						Key:   "byte",
						Value: byteValue{byte(2)},
					},
				},
			},
		},
		{
			name: "Marshals another message to valid JSON",
			input: defaultMessage{
				fields: []Field{
					{
						Key:   "count",
						Value: int64Value{int64(123456789101112)},
					},
					{
						Key:   "float",
						Value: float64Value{77.8},
					},
					{
						Key: "map",
						Value: mapValue{map[Value]Value{
							stringValue{"key"}: stringValue{"value"},
						}},
					},
				},
			},
		},
		{
			name: "Marshals empty message",
			input: defaultMessage{
				fields: []Field{},
			},
		},
		{
			name:  "Deals with nil",
			input: nil,
		},
		{
			name: "Deals with nil field value",
			input: defaultMessage{
				fields: []Field{
					{
						Key:   "nilKey",
						Value: nil,
					},
					{
						Key:   "test",
						Value: boolValue{true},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bytes := jsonMarshallerFn(test.input)

			if bytes != nil {
				var dat map[string]interface{}
				err := json.Unmarshal(bytes, &dat)

				assert.NoError(t, err)
			}
		})
	}
}

func TestMarshalJsonField(t *testing.T) {
	type testCase struct {
		name  string
		input Field
	}

	testCases := []testCase{
		{
			name: "Marshals given field",
			input: Field{
				Key:   "test",
				Value: stringValue{"testingValue"},
			},
		},
		{
			name: "Marshals another field",
			input: Field{
				Key:   "test",
				Value: boolValue{false},
			},
		},
		{
			name: "Deals with nil value",
			input: Field{
				Key:   "test",
				Value: nil,
			},
		},
		{
			name: "Deals with empty key",
			input: Field{
				Key:   "",
				Value: boolValue{false},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			bytes := marshalJsonField(test.input, -1, NoExponentFloatFormat)

			assert.NotNil(t, bytes)
		})
	}
}
