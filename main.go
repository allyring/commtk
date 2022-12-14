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

	// Bubbletea for the TUI
	tea "github.com/charmbracelet/bubbletea"

	// Bubbletea bubbles
	"github.com/charmbracelet/bubbles/spinner"
	boxer "github.com/treilik/bubbleboxer"
)

const (
	upperAddr  = "upper"
	leftAddr   = "left"
	middleAddr = "middle"
	rightAddr  = "right"
	lowerAddr  = "lower"
)


// SECTIONS:

// (Tool bubbles get stored in separate modules that get imported.)

// Model containers & relevant functions

// Main model (bubbleboxer model)

// Model container View, Update, and Init functions

// bubbleboxer tree generator from config parsing

// Model edit function
//

// Main function



func main() {
	// leaf content creation (models)
	upper := spinnerHolder{spinner.New()}
	left := stringer(leftAddr)
	middle := stringer(middleAddr)
	right := stringer(rightAddr)

	lower := stringer(fmt.Sprintf("%s: use ctrl+c to quit", lowerAddr))

	// layout-tree defintion
	m := model{tui: boxer.Boxer{}}
	m.tui.LayoutTree = boxer.Node{
		// orientation
		VerticalStacked: true,
		// spacing
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			return []int{
				// since this node is vertical stacked return the height partioning since the width stays for all children fixed
				widthOrHeight - 2,
				1,
				// make also sure that the amount of the returned ints match the amount of children:
				// in this case two, but in more complex cases read the amount of the chilren from the len(boxer.Node.Children)
			}
		},
		Children: []boxer.Node{
			{
				SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
					return []int{
						widthOrHeight/3,
						(widthOrHeight/3)*2,

					}
				},
				Children: []boxer.Node{
					m.tui.CreateLeaf(upperAddr, upper),
					{
						VerticalStacked: true,
						Children: []boxer.Node{
							{
								Children: []boxer.Node{
									m.tui.CreateLeaf(leftAddr, left),
									m.tui.CreateLeaf(rightAddr, right),
								},
							},
							// make sure to encapsulate the models into a leaf with CreateLeaf:
							m.tui.CreateLeaf(middleAddr, middle),
						},
					},
				},
			},
			m.tui.CreateLeaf(lowerAddr, lower),

		},
	}
	p := tea.NewProgram(m,tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}

type model struct {
	tui boxer.Boxer
}

func (m model) Init() tea.Cmd {
	return spinner.Tick
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		err := m.tui.UpdateSize(msg)
		if err != nil {
			return nil, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		err := m.editModel(upperAddr, func(v tea.Model) (tea.Model, error) {
			v, cmd = v.Update(msg)
			return v, nil
		})
		if err != nil {
			return nil, nil
		}
		return m, cmd
	}
	return m, nil
}
func (m model) View() string {
	return m.tui.View()
}

func (m *model) editModel(addr string, edit func(tea.Model) (tea.Model, error)) error {
	if edit == nil {
		return fmt.Errorf("no edit function provided")
	}
	v, ok := m.tui.ModelMap[addr]
	if !ok {
		fmt.Errorf("no model with address '%s' found", addr)
	}
	v, err := edit(v)
	if err != nil {
		return err
	}
	m.tui.ModelMap[addr] = v
	return nil
}

type stringer string

func (s stringer) String() string {
	return string(s)
}

// satisfy the tea.Model interface
func (s stringer) Init() tea.Cmd                           { return nil }
func (s stringer) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return s, nil }
func (s stringer) View() string                            { return s.String() }

type spinnerHolder struct {
	m spinner.Model
}

func (s spinnerHolder) Init() tea.Cmd {
	return s.m.Tick
}

func (s spinnerHolder) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m, cmd := s.m.Update(msg)
	s.m = m
	return s, cmd
}
func (s spinnerHolder) View() string {
	return s.m.View()
}
