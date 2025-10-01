package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4")).
			MarginBottom(1)

	optionStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)
)

// Model represents the state of our TUI
type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
	step     string // "main_menu", "select_dirs", "machine_type", "auth"
}

func initialModel() model {
	return model{
		choices: []string{
			"Initialize new sync",
			"Add directory to sync",
			"Remove directory from sync",
			"Check sync status",
			"Manual push",
			"Manual pull",
			"Start daemon",
			"Stop daemon",
			"Exit",
		},
		selected: make(map[int]struct{}),
		step:     "main_menu",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter":
			return m.handleSelection()
		}
	}

	return m, nil
}

func (m model) handleSelection() (tea.Model, tea.Cmd) {
	switch m.step {
	case "main_menu":
		switch m.cursor {
		case 0: // Initialize
			// TODO: Navigate to initialization wizard
			fmt.Println("\nInitializing omarchy-sync...")
			return m, tea.Quit
		case 1: // Add directory
			fmt.Println("\nAdd directory feature - coming soon!")
			return m, tea.Quit
		case 2: // Remove directory
			fmt.Println("\nRemove directory feature - coming soon!")
			return m, tea.Quit
		case 3: // Status
			fmt.Println("\nChecking status...")
			return m, tea.Quit
		case 4: // Push
			fmt.Println("\nPushing changes...")
			return m, tea.Quit
		case 5: // Pull
			fmt.Println("\nPulling changes...")
			return m, tea.Quit
		case 6: // Start daemon
			fmt.Println("\nStarting daemon...")
			return m, tea.Quit
		case 7: // Stop daemon
			fmt.Println("\nStopping daemon...")
			return m, tea.Quit
		case 8: // Exit
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("🔄 Omarchy Sync") + "\n\n"

	switch m.step {
	case "main_menu":
		s += "Select an option:\n\n"

		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = "›"
				s += selectedStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
			} else {
				s += optionStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
			}
		}

		s += helpStyle.Render("\n↑/k: up • ↓/j: down • enter: select • q: quit")
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
