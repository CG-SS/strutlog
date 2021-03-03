package strutlog

// LogSampleConfig is the configuration used by MultiLevelLogSamplerFn. Each field represents how often it will log a
// given message based on 1 in X, where X is the given rate. For instance, if TraceRate is 1, it means it will log 1 in
// 1, that is, it will always log it. If it was 10, it would log 1 in 10, that is, 10% of the time.
type LogSampleConfig struct {
	TraceRate uint
	DebugRate uint
	InfoRate  uint
	WarnRate  uint
	ErrorRate uint
	FatalRate uint
	PanicRate uint
}

// MultiLevelLogSamplerFn is a sampling function that will use LogSampleConfig to sample a Message based on their
// MessageLevel.
func MultiLevelLogSamplerFn(cfg LogSampleConfig) HookFn {
	sampleTraceFn := SingleLogSamplerFn(cfg.TraceRate)
	sampleDebugFn := SingleLogSamplerFn(cfg.DebugRate)
	sampleInfoFn := SingleLogSamplerFn(cfg.InfoRate)
	sampleWarnFn := SingleLogSamplerFn(cfg.WarnRate)
	sampleErrorFn := SingleLogSamplerFn(cfg.ErrorRate)
	samplePanicFn := SingleLogSamplerFn(cfg.PanicRate)
	sampleFatalFn := SingleLogSamplerFn(cfg.FatalRate)

	return func(message Message) Message {
		switch message.Level() {
		case TraceLevel:
			return sampleTraceFn(message)
		case DebugLevel:
			return sampleDebugFn(message)
		case InfoLevel:
			return sampleInfoFn(message)
		case WarnLevel:
			return sampleWarnFn(message)
		case ErrorLevel:
			return sampleErrorFn(message)
		case PanicLevel:
			return samplePanicFn(message)
		case FatalLevel:
			return sampleFatalFn(message)
		default:
			return nil
		}
	}
}
