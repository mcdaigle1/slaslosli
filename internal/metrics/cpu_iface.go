package metrics

// Define an interface
type CPU interface {
	Load() string
	ResponseCount(string, string) string
}
