package d

import "fmt"

// UI represents the user interface
type UI struct {
	m *Model // The UI has a concrete dependency with the model
}

// DisplayData displays data to the screen
func (ui UI) DisplayData() {
	rawData := ui.m.GetData()
	fmt.Println(rawData)
}

// Model contains representation of the business model
type Model struct {
	data string
}

// GetData retrieves business data from the model
func (m *Model) GetData() string {
	return m.data
}
