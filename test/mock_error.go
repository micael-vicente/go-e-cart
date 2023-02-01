package services

type MockError struct {
	errorMessage string
}

func (m *MockError) Error() string {
	return m.errorMessage
}
