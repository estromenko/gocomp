package stack

// stackOverflowError ...
type stackOverflowError struct{}

// Error ...
func (s *stackOverflowError) Error() string {
	return "Error: stack overflow!"
}
