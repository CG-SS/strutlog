package strutlog

// defaultMessage represents a 'normal' message from the defaultLogger.
// In short, it will append a Field on fields, and then send it to the internal messageCh, in which it will be written
// by the defaultLogger.
// All the receivers are value receivers, and each field append function return a copy of the Message with the new field
// appended.
type defaultMessage struct {
	messageCh chan<- Message
	fields    []Field
	level     MessageLevel
}

// Fields will return the appended fields so far.
func (d defaultMessage) Fields() []Field {
	return d.fields
}

// Bool appends a bool value to this Message with the given key.
func (d defaultMessage) Bool(key string, value bool) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Bool(value),
		}),
		level: d.level,
	}
}

// String appends a string value to this Message with the given key.
func (d defaultMessage) String(key string, value string) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: String(value),
		}),
		level: d.level,
	}
}

// Int appends an int value to this Message with the given key.
func (d defaultMessage) Int(key string, value int) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Int(value),
		}),
		level: d.level,
	}
}

// Int8 appends an int8 value to this Message with the given key.
func (d defaultMessage) Int8(key string, value int8) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Int8(value),
		}),
		level: d.level,
	}
}

// Int16 appends an int16 value to this Message with the given key.
func (d defaultMessage) Int16(key string, value int16) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Int16(value),
		}),
		level: d.level,
	}
}

// Int32 appends an int32 value to this Message with the given key.
func (d defaultMessage) Int32(key string, value int32) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Int32(value),
		}),
		level: d.level,
	}
}

// Int64 appends an int64 value to this Message with the given key.
func (d defaultMessage) Int64(key string, value int64) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Int64(value),
		}),
		level: d.level,
	}
}

// UInt appends an uint value to this Message with the given key.
func (d defaultMessage) UInt(key string, value uint) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: UInt(value),
		}),
		level: d.level,
	}
}

// UInt8 appends an uint8 value to this Message with the given key.
func (d defaultMessage) UInt8(key string, value uint8) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: UInt8(value),
		}),
		level: d.level,
	}
}

// UInt16 appends an uint16 value to this Message with the given key.
func (d defaultMessage) UInt16(key string, value uint16) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: UInt16(value),
		}),
		level: d.level,
	}
}

// UInt32 appends an uint32 value to this Message with the given key.
func (d defaultMessage) UInt32(key string, value uint32) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: UInt32(value),
		}),
		level: d.level,
	}
}

// UInt64 appends an uint64 value to this Message with the given key.
func (d defaultMessage) UInt64(key string, value uint64) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: UInt64(value),
		}),
		level: d.level,
	}
}

// Byte appends a byte value to this Message with the given key.
func (d defaultMessage) Byte(key string, value byte) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Byte(value),
		}),
		level: d.level,
	}
}

// Rune appends a rune value to this Message with the given key.
func (d defaultMessage) Rune(key string, value rune) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Rune(value),
		}),
		level: d.level,
	}
}

// Float32 appends a float32 value to this Message with the given key.
func (d defaultMessage) Float32(key string, value float32) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Float32(value),
		}),
		level: d.level,
	}
}

// Float64 appends a float64 value to this Message with the given key.
func (d defaultMessage) Float64(key string, value float64) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Float64(value),
		}),
		level: d.level,
	}
}

// Complex64 appends a complex64 value to this Message with the given key.
func (d defaultMessage) Complex64(key string, value complex64) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Complex64(value),
		}),
		level: d.level,
	}
}

// Complex128 appends a complex128 value to this Message with the given key.
func (d defaultMessage) Complex128(key string, value complex128) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Complex128(value),
		}),
		level: d.level,
	}
}

// Array appends an array of Value to this Message with the given key.
func (d defaultMessage) Array(key string, value []Value) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Array(value),
		}),
		level: d.level,
	}
}

// Map appends a map[Value]Value to this Message with the given key.
func (d defaultMessage) Map(key string, value map[Value]Value) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Map(value),
		}),
		level: d.level,
	}
}

// Strut appends a Fieldable to this Message with the given key.
func (d defaultMessage) Strut(key string, value Fieldable) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: Strut(value),
		}),
		level: d.level,
	}
}

// Error appends an error to this Message with the given key.
func (d defaultMessage) Error(key string, value error) Message {
	return defaultMessage{
		messageCh: d.messageCh,
		fields: append(d.fields, Field{
			Key:   key,
			Value: String(value.Error()),
		}),
		level: d.level,
	}
}

// Level returns the MessageLevel of this Message.
func (d defaultMessage) Level() MessageLevel {
	return d.level
}

// Send the Message to be logged.
func (d defaultMessage) Send() Message {
	d.messageCh <- d

	return d
}
