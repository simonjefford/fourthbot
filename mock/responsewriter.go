package mock

import "bytes"

type MockResponseWriter struct {
	*bytes.Buffer
	WriteStatusInvoked bool
	WrittenStatus      int
}

func NewMockResponseWriter() *MockResponseWriter {
	return &MockResponseWriter{
		Buffer: new(bytes.Buffer),
	}
}

func (m *MockResponseWriter) WriteStatus(s int) {
	m.WrittenStatus = s
	m.WriteStatusInvoked = true
}
