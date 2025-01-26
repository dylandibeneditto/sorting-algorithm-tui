package main

import (
	//"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define different program states
type state int

const (
	stateList state = iota
	stateSorting
)

var (
	docStyle           = lipgloss.NewStyle().Margin(1, 2)
	selectedTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#fcba03", Dark: "#fcba03"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#fcba03", Dark: "#fcba03"}).
				Padding(0, 0, 0, 1)
)

// Item structure for the list
type item struct {
	title string
	time  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.time }
func (i item) FilterValue() string { return i.title }

// Model to handle states
type model struct {
	state       state
	list        list.Model
	selectedAlg string
	numbers     []int
	step        int
}

func initialModel() model {
	items := []list.Item{
		item{"Bubble Sort", "O(n^2)"},
		item{"Insertion Sort", "O(n^2)"},
		item{"Selection Sort", "O(n^2)"},
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = selectedTitleStyle
	delegate.Styles.SelectedDesc = selectedTitleStyle
	l := list.New(items, delegate, 20, 10)
	l.Title = "Select a Sorting Algorithm"
	l.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 1)

	return model{
		state: stateList,
		list:  l,
	}
}

// Bubble Tea update function
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q": // Quit
			return m, tea.Quit
		case "enter":
			// Switch to sorting visualization state
			m.selectedAlg = m.list.SelectedItem().(item).title
			m.state = stateSorting
			m.numbers = generateRandomNumbers(100)
			m.step = 0
			return m, tick()
		}
	case tickMsg:
		// Advance the sorting step
		if m.step < len(m.numbers)-1 {
			m.numbers = bubbleSortStep(m.numbers, m.step)
			m.step++
			return m, tick()
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

// Bubble Tea view function
func (m model) View() string {
	switch m.state {
	case stateList:
		return docStyle.Render(m.list.View())
	case stateSorting:
		return docStyle.Render(visualizeArray(m.numbers))
	}
	return "Unknown state"
}

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func visualizeArray(arr []int) string {
	doc := strings.Builder{}
	maxHeight := findMax(arr)

	//divisor := maxHeight / 8 // Normalize height to fit block characters
	//if divisor == 0 {
	//	divisor = 1
	//}

	divisor := 8

	// Create vertical bars from top to bottom
	for i := maxHeight/divisor + 1; i > 0; i-- {
		for _, val := range arr {
			var c string
			v := val + divisor
			if v/divisor >= i {
				c = "█" // Full block
				if v/divisor == i {
					switch v % divisor {
					case 0:
						c = " "
					case 1:
						c = "▁"
					case 2:
						c = "▂"
					case 3:
						c = "▃"
					case 4:
						c = "▄"
					case 5:
						c = "▅"
					case 6:
						c = "▆"
					case 7:
						c = "▇"
					}
				}
				s := lipgloss.NewStyle().SetString(c).Foreground(lipgloss.Color("#0000FF"))
				doc.WriteString(s.String()) // Block for value
			} else {
				doc.WriteString(" ") // Empty space
			}
		}
		doc.WriteRune('\n')
	}

	return doc.String()
}

// Tick message for updating visualization
type tickMsg struct{}

func tick() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(100 * time.Millisecond)
		return tickMsg{}
	}
}

// Bubble Sort step-by-step function
func bubbleSortStep(arr []int, step int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	if step < len(arr)-1 {
		for j := 0; j < len(arr)-step-1; j++ {
			if newArr[j] > newArr[j+1] {
				newArr[j], newArr[j+1] = newArr[j+1], newArr[j]
			}
		}
	}
	return newArr
}

// Generate random numbers for sorting visualization
func generateRandomNumbers(size int) []int {
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = i + 1
	}
	for i := range numbers {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
