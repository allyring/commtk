/*
Commtk is a suite of TUIs for common TCP/IP networking tools.
It displays tools on a scalable grid, where the contents and sizing of the grid's cells are defined by a config file.
It contains TUIs for the following tools:
- ping
- traceroute
- ifconfig

Usage:

	commtk [config file path ...]

Licensed under GNU GPLv3
*/
package main

import (
	"fmt"
	"os"

	// Packages for the project
	"github.com/allyring/commtk/cfg"

	// Bubbletea, stickers, and bubbles - for TUI rendering
	"github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
// ---------------------------------------------------------------------------------------------------------------------
// Lipgloss styling
var (
	defaultStyle = 	lipgloss.NewStyle().
						Border(lipgloss.NormalBorder()).
						BorderForeground(lipgloss.Color("e5e5e5"))
)
// ---------------------------------------------------------------------------------------------------------------------
// Bubbletea model

type model struct {
	err			error				// Contains the most recent error encountered
	flexBox		*stickers.FlexBox	// Pointer to a flexbox in memory
}
// ---------------------------------------------------------------------------------------------------------------------
// Basic model renderer

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.flexBox.SetWidth(msg.Width)
		m.flexBox.SetHeight(msg.Height)

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}
func (m *model) View() string {
	return m.flexBox.Render()
}



// ---------------------------------------------------------------------------------------------------------------------

func main() {
	// The args are a path to a custom config. If the args exist, pass them into the config layout generator.
	configPath := ""
	if len(os.Args[1:]) > 1 {
		fmt.Println("Cannot accept several configs.")
		os.Exit(1)
	}
	if len(os.Args[1:]) == 1 {
		if len(os.Args[1]) > 0 {
			configPath = os.Args[1]
		}
	}


	// Create the layout config from the given file / the default config.
	err, schemaErr := cfg.LoadLayoutFile(configPath)
	if err != nil {
		panic(err)
	}

	if schemaErr != nil {
		fmt.Println("JSON errors")
		for _, schemaErrItem := range schemaErr {
			fmt.Println("- " + schemaErrItem.Description())
		}
		panic(err)
	}

	// The config is known to be valid, so now create a layout from it to pass into the initial model.
	// Create a new variable to hold pointer to the flexbox we generate
	generatedFlexBox := stickers.NewFlexBox(0,0)
	var allRows []*stickers.FlexBoxRow


	// Loop through the rows in the parsed config
	for _, rowContents := range cfg.Config.Layout {
		newRow := generatedFlexBox.NewRow()
		var newRowCells []*stickers.FlexBoxCell

		for _, cellData := range rowContents {
			newCell := stickers.NewFlexBoxCell(cellData.RatioX,cellData.RatioY).SetStyle(defaultStyle)
			newRowCells = append(newRowCells, newCell)
		}
		newRow.AddCells(newRowCells)
		allRows = append(allRows, newRow)
	}

	generatedFlexBox.AddRows(allRows)



	initialModel := model{
		flexBox: generatedFlexBox,
		err:     nil,
	}

	p := tea.NewProgram(&initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error when starting commtk: %v", err)
		os.Exit(1)
	}

}