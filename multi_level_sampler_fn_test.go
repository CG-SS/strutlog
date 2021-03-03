package strutlog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testSampleGivenLevel(t *testing.T, level MessageLevel, rate uint, fn HookFn) {
	t.Helper()

	for i := uint(0); i < rate; i++ {
		if i != rate-1 {
			msg := fn(defaultMessage{level: level})

			assert.Nil(t, msg)
		}
	}
	msg := fn(defaultMessage{level: level})

	assert.NotNil(t, msg)
}

func TestMultiLevelLogSamplerFn(t *testing.T) {
	type testCase struct {
		name string
		cfg  LogSampleConfig
	}

	testCases := []testCase{
		{
			name: "Samples multi given config",
			cfg: LogSampleConfig{
				TraceRate: 1,
				DebugRate: 2,
				InfoRate:  3,
				WarnRate:  4,
				ErrorRate: 5,
				FatalRate: 6,
				PanicRate: 7,
			},
		},
		{
			name: "Samples another multi given config",
			cfg: LogSampleConfig{
				TraceRate: 10,
				DebugRate: 20,
				InfoRate:  30,
				WarnRate:  40,
				ErrorRate: 50,
				FatalRate: 60,
				PanicRate: 70,
			},
		},
		{
			name: "Handle zero rate",
			cfg: LogSampleConfig{
				TraceRate: 0,
				DebugRate: 0,
				InfoRate:  0,
				WarnRate:  0,
				ErrorRate: 0,
				FatalRate: 0,
				PanicRate: 0,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			cfg := test.cfg

			fn := MultiLevelLogSamplerFn(cfg)

			testSampleGivenLevel(t, TraceLevel, cfg.TraceRate, fn)
			testSampleGivenLevel(t, DebugLevel, cfg.DebugRate, fn)
			testSampleGivenLevel(t, InfoLevel, cfg.InfoRate, fn)
			testSampleGivenLevel(t, WarnLevel, cfg.WarnRate, fn)
			testSampleGivenLevel(t, ErrorLevel, cfg.ErrorRate, fn)
			testSampleGivenLevel(t, FatalLevel, cfg.FatalRate, fn)
			testSampleGivenLevel(t, PanicLevel, cfg.PanicRate, fn)
		})
	}
}
