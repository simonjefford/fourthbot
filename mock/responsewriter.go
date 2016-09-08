package mock

import "bytes"

type MockResponseWriter struct {
	*bytes.Buffer
	WriteStatusInvoked bool
	WrittenStatus      string
}

func NewMockResponseWriter() *MockResponseWriter {
	return &MockResponseWriter{
		Buffer: new(bytes.Buffer),
	}
}

func (m *MockResponseWriter) WriteStatus(s string) {
	m.WrittenStatus = s
	m.WriteStatusInvoked = true
}
