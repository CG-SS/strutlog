package strutlog

import (
	"testing"
)

func TestSingleLogSamplerFnLogs(t *testing.T) {
	type testCase struct {
		name string
		rate uint
	}

	testCases := []testCase{
		{
			name: "Logs 1 in 1",
			rate: 1,
		},
		{
			name: "Logs 1 in 10",
			rate: 10,
		},
		{
			name: "Logs 1 in 100",
			rate: 100,
		},
		{
			name: "Handles 0 rate",
			rate: 0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fn := SingleLogSamplerFn(test.rate)

			// Shouldn't distinguish between message levels.
			testSampleGivenLevel(t, InfoLevel, test.rate, fn)
			testSampleGivenLevel(t, DebugLevel, test.rate, fn)
		})
	}
}
