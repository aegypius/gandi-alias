package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/qmuntal/stateless"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	email    textinput.Model
	alias    textinput.Model
	aliases  list.Model
	state    *stateless.StateMachine
	quitting bool
	err      error
}

var quitKeys = key.NewBinding(
	key.WithKeys("esc", "ctrl+c"),
	key.WithHelp("", "press ESC to quit"),
)
var validateKeys = key.NewBinding(
	key.WithKeys("enter", "tab"),
	key.WithHelp("", "press ENTER to continue"),
)

func InitModel() model {
	emailInput := textinput.New()
	emailInput.Placeholder = "Email"
	emailInput.Focus()
	emailInput.CharLimit = 156
	emailInput.Width = 50

	items := []list.Item{}
	aliasesList := list.New(items, list.NewDefaultDelegate(), 0, 0)

	return model{
		email:   emailInput,
		aliases: aliasesList,
		state:   initStateMachine(),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit

		}

		if key.Matches(msg, validateKeys) {
			log.Info("current state", "state", m.state.MustState())
			m.state.Fire(triggerEmailStored, m)
			return m, nil
		}

		m.email, _ = m.email.Update(msg)
		m.aliases, cmd = m.aliases.Update(msg)

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.aliases.SetSize(msg.Width-h, msg.Height-v)
	}

	return m, cmd
}

func (m model) View() string {
	var view string
	switch m.state.MustState() {
	case stateNew:
		view = m.email.View()
	case stateEmailStored, stateAliasesQueried, stateAliasAdded:
		view = m.aliases.View()
	}

	return docStyle.Render(view)
}
