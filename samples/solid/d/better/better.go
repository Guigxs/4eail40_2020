package d

import "fmt"

// UI presents the user interface
type UI struct {
	m *Controller // Now we depend on an abstraction
}

// DisplayData displays data to the screen
func (ui UI) DisplayData() {
	rawData := ui.m.GetData()
	fmt.Println(rawData)
}

type Controller interface {
	GetData() string
}

// Model contains representation of the business model
type Model struct {
	data string
}

// GetData retrieves business data from the model
func (m *Model) GetData() string {
	return m.data
}
