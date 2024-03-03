package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kaporos/azathoth/components"
	"github.com/kaporos/azathoth/core"
	"github.com/kaporos/azathoth/stores"
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

func (m *Model) InitMutable() {
	m.player.GiveItem(stores.CreateItemSure("apple"))
	m.player.GiveItem(stores.CreateItemSure("skin"))
}
func (m Model) Init() tea.Cmd {
	return nil
}

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

func (m Model) View() string {
	map_view := components.RenderMap(&m.player)
	player_stats := components.RenderPlayerStats(&m.player)
	history_style := m.TxtStyle.MarginLeft(5).MarginTop(1)

	//Computing size of visible history (to avoid overflow)
	size := len(strings.Split(m.History, "\n")) - m.Height*2/3
	if size < 0 {
		size = 0
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		map_view+"\n\n\n"+player_stats,
		history_style.Render(strings.Join(strings.Split(m.History, "\n")[size:], "\n")+components.RenderPrompt(m.input, true)))
}
