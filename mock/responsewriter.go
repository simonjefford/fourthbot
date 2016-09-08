package mock

import "bytes"

// ResponseWriter is a mock implementation of fourthbot.ResponseWriter
type ResponseWriter struct {
	*bytes.Buffer
	WriteStatusInvoked bool
	WrittenStatus      int
}

// NewResponseWriter creates a ResponseWriter
func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		Buffer: new(bytes.Buffer),
	}
}

// WriteStatus is a mock implementation of
// fourthbot.ResponseWriter.WriteStatus
func (m *ResponseWriter) WriteStatus(s int) {
	m.WrittenStatus = s
	m.WriteStatusInvoked = true
}
