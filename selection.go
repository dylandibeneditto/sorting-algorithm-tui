package main

import (
    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)
var (
    docStyle = lipgloss.NewStyle().Margin(1, 2)

    titleStyle = lipgloss.NewStyle().
                    Foreground(lipgloss.Color("#FFFFFF")).
                    Padding(0, 1)

    selectedTitleStyle = lipgloss.NewStyle().
                            Border(lipgloss.NormalBorder(), false, false, false, true).
                            BorderForeground(lipgloss.AdaptiveColor{Light: "#fcba03", Dark: "#fcba03"}).
                            Foreground(lipgloss.AdaptiveColor{Light: "#fcba03", Dark: "#fcba03"}).
                            Padding(0, 0, 0, 1)
)

var items = []list.Item{
    item{title: "Bubble Sort", time: "O(n^2)"},
}

type item struct {
    title, time string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.time }
func (i item) FilterValue() string { return i.title }

type model struct {
    list list.Model
}

func NewList() model {
    d := list.NewDefaultDelegate()
    d.Styles.SelectedTitle = selectedTitleStyle
    d.Styles.SelectedDesc = selectedTitleStyle
    m := model{list: list.New(items, d, 0, 0)}
    m.list.Title = "Select an Algorithm"
    m.list.Styles.Title = titleStyle
    return m
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width-h, msg.Height-v)
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m model) View() string {
    return docStyle.Render(m.list.View())
}

