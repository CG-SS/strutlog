package strutlog

// SingleLogSamplerFn is a log sampler that samples a Message based on the given rate.
// It will ignore the MessageLevel, treating every Message the same.
func SingleLogSamplerFn(rate uint) HookFn {
	// Rate cannot be 0, as that will cause a div by 0.
	if rate == 0 {
		rate = 1
	}

	count := uint(0)

	return func(message Message) Message {
		count++

		if count%rate == 0 {
			return message
		}

		return nil
	}
}
