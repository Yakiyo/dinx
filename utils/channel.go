package utils

// List of channels
var Channels = [3]string{"stable", "beta", "dev"}

// Check if channel is valid
func IsValidChannel(channel string) bool {
	for _, c := range Channels {
		if c == channel {
			return true
		}
	}
	return false
}
