package main

import (
    "fmt"
//    "strings"
    
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    algorithms  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    selected int
}

func initialModel() model {
    return model{
        algorithms: []string{"Bubble Sort", "Bogo Sort"},
        selected: 0,
    }
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.algorithms)-1 {
                m.cursor++
            }

        // The "enter" key and the spacebar (a literal space) toggle
        // the selected state for the item that the cursor is pointing at.
        case "enter", " ":
//            algo := m.algorithms[m.selected]
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "Select an Algorithm\n\n"

    // Iterate over our algorithms
    for i, algorithm := range m.algorithms {

        // Is the cursor pointing at this algorithm?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, algorithm)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}
