package utils

// Form the appropiate url to the index for provided channel
func MakeUrl(channel string) string {
	return `https://www.googleapis.com/storage/v1/b/dart-archive/o?delimiter=%2F&prefix=channels%2F` + channel + `%2Frelease%2F&alt=json`
}
