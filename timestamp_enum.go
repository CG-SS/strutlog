package strutlog

// TimestampType represents the type of timestamp that will get used on the message.
type TimestampType int

const (
	// RFC3339Timestamp corresponds to the time.RFC3339 format.
	RFC3339Timestamp TimestampType = iota
	// RFC3339NanoTimestamp corresponds to the time.RFC3339Nano format.
	RFC3339NanoTimestamp
	// LayoutTimestamp corresponds to the time.Layout format.
	LayoutTimestamp
	// ANSICTimestamp corresponds to the time.ANSIC format.
	ANSICTimestamp
	// UnixDateTimestamp corresponds to the time.UnixDate format.
	UnixDateTimestamp
	// RubyDateTimestamp corresponds to the time.RubyDate format.
	RubyDateTimestamp
	// RFC822Timestamp corresponds to the time.RFC822 format.
	RFC822Timestamp
	// RFC822ZTimestamp corresponds to the time.RFC822Z format.
	RFC822ZTimestamp
	// RFC850Timestamp corresponds to the time.RFC850 format.
	RFC850Timestamp
	// RFC1123Timestamp corresponds to the time.RFC1123 format.
	RFC1123Timestamp
	// RFC1123ZTimestamp corresponds to the time.RFC1123Z format.
	RFC1123ZTimestamp
	// KitchenTimestamp corresponds to the time.Kitchen format.
	KitchenTimestamp
	// StampTimestamp corresponds to the time.Stamp format.
	StampTimestamp
	// StampMilliTimestamp corresponds to the time.StampMilli format.
	StampMilliTimestamp
	// StampMicroTimestamp corresponds to the time.StampMicro format.
	StampMicroTimestamp
	// StampNanoTimestamp corresponds to the time.StampNano format.
	StampNanoTimestamp
	// DateTimeTimestamp corresponds to the time.DateTime format.
	DateTimeTimestamp
	// DateOnlyTimestamp corresponds to the time.DateOnly format.
	DateOnlyTimestamp
	// TimeOnlyTimestamp corresponds to the time.TimeOnly format.
	TimeOnlyTimestamp
	// UnixTimestamp corresponds to the time.Unix format.
	UnixTimestamp
	// UnixNanoTimestamp corresponds to the time.UnixNano format.
	UnixNanoTimestamp
	// UnixMicroTimestamp corresponds to the time.UnixMicro format.
	UnixMicroTimestamp
	// UnixMilliTimestamp corresponds to the time.UnixMilli format.
	UnixMilliTimestamp
	// CustomTimestamp represents a custom timestamp that follows time.Format.
	CustomTimestamp
	// DisabledTimestamp disables timestamping.
	DisabledTimestamp
)
