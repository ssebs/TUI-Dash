package tui

import (
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type MainWindow struct {
	stopwatch   stopwatch.Model
	w, h        int
	currentTime time.Time
	// header      tea.Model
	// footer      tea.Model
	mainArea string
	cmd      string
}

func CreateMainWindow() MainWindow {
	return MainWindow{
		currentTime: time.Now(),
		mainArea:    "TUI Dash",
		stopwatch:   stopwatch.NewWithInterval(1 * time.Second),
		cmd:         "pwd",
	}
}

func (m MainWindow) Init() tea.Cmd {
	return m.stopwatch.Init()

}

func (m MainWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.currentTime = time.Now()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			// open text window
			m.mainArea = "enter pressed"
		case "c":
			_c := exec.Command(m.cmd)
			cmd := tea.ExecProcess(_c, nil)
			return m, cmd
		}

	case stopwatch.TickMsg:
		// m.mainArea = "current time tick'd: " + m.currentTime.Format("15:04:05")
	}

	// Update timer
	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m MainWindow) View() string {
	// re-renders every tick
	m.setWindowSize()

	s := m.renderHeader()
	s += "\n"
	s += renderBox(m.mainArea)
	s += "\n"

	// s += m.renderCmd()

	return s
}

func (m *MainWindow) renderCmd() string {
	return ""
}

func (m MainWindow) renderHeader() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		AlignHorizontal(lipgloss.Center).
		Width(m.w-3).Margin(0, 1) // magic #'s

	s := m.currentTime.Format("15:04:05")

	return headerStyle.Render(s)
}

func renderBox(msg string) string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		// Background(lipgloss.Color("#7D56F4")).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		AlignHorizontal(lipgloss.Center)

	return style.Render(msg)
}

func (m *MainWindow) setWindowSize() {
	w, h, _ := term.GetSize(os.Stdout.Fd())
	if w != m.w || h != m.h {
		m.w = w
		m.h = h
	}
}
