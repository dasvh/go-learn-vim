package testutils

import tea "github.com/charmbracelet/bubbletea"

// MockTeaModel is a mock implementation of tea.Model for testing
type MockTeaModel struct {
	InitCalled   bool
	UpdateCalled bool
	ViewCalled   bool
}

func (m *MockTeaModel) Init() tea.Cmd {
	m.InitCalled = true
	return nil
}

func (m *MockTeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.UpdateCalled = true
	return m, nil
}

func (m *MockTeaModel) View() string {
	m.ViewCalled = true
	return "mock screen view"
}
