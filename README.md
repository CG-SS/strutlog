# strutlog

Structured logging library for Golang.

It's similar to `zerolog`, however, it isn't only a JSON logger, you can implement any form of structured logging by 
implementing your own marshaller. It also has better interface principles as everything is an interface, allowing the 
user to easily implement their own messages and loggers.

Currently, it can log JSON and Protobuf.

## Example

```
package main

import (
	"context"
	"errors"
	"io"
	"os"
	"strutlog"
	"time"
)

type fieldTest struct {
}

func (f fieldTest) Fields() []strutlog.Field {
	return []strutlog.Field{
		{
			Key:   "testField",
			Value: strutlog.UInt(1),
		},
		{
			Key:   "testField2",
			Value: strutlog.Bool(false),
		},
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	f, err := os.Create("./log.txt")
	if err != nil {
		panic(err)
	}

    // This will return a logger with a buffer size of 4096, that will write messages to the stdout and to the file
    // 'log.txt', where the timestamp will be in the RFC3339 format and with a field 'hookFn' with the value 'message
    // from a HookFn'.    
	logger := strutlog.NewLogger(ctx, strutlog.DefaultLoggerConfig{
		BufferSize: 4096,
		Writer: strutlog.SingleMessageWriter{
			Writer: io.MultiWriter(f, os.Stdout),
			Marshal: strutlog.JsonMarshaller(-1, strutlog.NoExponentFloatFormat),
		},
		Timestamp: strutlog.TimestampConfig{
			Type:      strutlog.RFC3339NanoTimestamp,
			FieldName: "timestamp",
		},
		HookFns: []strutlog.HookFn{
			func(message strutlog.Message) strutlog.Message {
				return message.String("hookFn", "message from a HookFn")
			},
		},
	})

	msg := logger.
		Info().
		String("message", "hello there").
		Bool("test", true).
		Float32("x", 1.23456789).
		Float64("y", 3.234567891234).
		Array("arr", []strutlog.Value{
			strutlog.UInt64(2),
			strutlog.String("test"),
			strutlog.Complex64(1),
			strutlog.Complex128(111),
			strutlog.Int(-11),
		}).
		Map("object", map[strutlog.Value]strutlog.Value{
			strutlog.String("field"): strutlog.Byte('0'),
			strutlog.String("test"):  strutlog.Byte(100),
		}).
		Strut("strutTest", fieldTest{}).
		Error("error", errors.New("failed to send"))

	msg.Int("messageNum", 2).Send()
	msg.Send()

	msg.Int("anotherOne", 100).String("thing", "heelo there this is another message").Send()

	debugMsg := logger.Info().String("small", "yes")

	for i := 0; i < 10; i++ {
		debugMsg.Send()
	}

	logger.Error().String("oops", "oops").Strut("error", msg).Send()

	<-time.After(5 * time.Second)
}
```

## Todo

- More tests
- More docs
- More marshallers
- Add examples

