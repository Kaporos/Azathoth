package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kaporos/azathoth/components"
	"github.com/kaporos/azathoth/core"
)

type Model struct {
	Term     string
	Width    int
	Height   int
	TxtStyle lipgloss.Style
	player   core.Player
	History  string
	input    string
}

// This Init will update m , the bubbletea Init below cannot change m because of its signature
func (m *Model) InitMutable() {
	//Giving the player basics items at log-in
}

// bubbletea init method. required to have but it does not do anything.
func (m Model) Init() tea.Cmd {
	return nil
}

// bubbletea update method. called at each event (inside of msg var) and returns updated model, as well as tea commands
// tea commands are related to IO stuff, we don't use them ATM (except for disconnecting player on quit)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down":
			m.player.Move(0, 1)
		case "up":
			m.player.Move(0, -1)
		case "right":
			m.player.Move(1, 0)
		case "left":
			m.player.Move(-1, 0)
		case "enter":
			if len(m.input) == 0 {
				break
			}
			if m.input == "clear" {
				m.History = ""
				m.input = ""
				break
			}
			if m.input == "quit" {
				return m, tea.Quit
			}

			cmdResult := core.ProcessCommand(&m.player, m.input)
			m.History += components.RenderPrompt(m.input, false) + "\n" + cmdResult + "\n"
			m.input = ""
		case "backspace", "delete":
			if len(m.input) == 0 {
				break
			}
			m.input = m.input[:len(m.input)-1]
		case "esc":
			break
		case "alt":
			break
		case "shift":
			break
		case "tab":
			break
		default:
			m.input += msg.String()
		}

	}
	return m, nil
}

// bubbletea view method. It generates what the player sees.
func (m Model) View() string {
	map_view := components.RenderMap(&m.player)
	player_stats := components.RenderPlayerStats(&m.player)
	left_style := m.TxtStyle.MarginLeft(5).MarginTop(1)

	// Computing size of visible history (to avoid overflow)
	size := len(strings.Split(m.History, "\n")) - m.Height*2/3
	if size < 0 {
		size = 0
	}

	right_side := map_view + "\n\n\n" + player_stats

	//stripping history to only keep max 2/3 of terminal's height.
	left_side := strings.Join(strings.Split(m.History, "\n")[size:], "\n")

	//adding user prompt
	left_side += components.RenderPrompt(m.input, true)

	return lipgloss.JoinHorizontal(
		//We want to align views at top
		lipgloss.Top,
		right_side,
		left_style.Render(left_side),
	)
}
