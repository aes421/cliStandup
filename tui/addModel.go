package tui

import (
	"github.com/aes421/cliStandup/state"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type addModel struct {
	textArea textarea.Model
	help     help.Model
}

func NewAddModel() addModel {
	m := addModel{
		textArea: textarea.New(),
		help:     help.New(),
	}
	m.textArea.Placeholder = "Enter your update here"
	m.textArea.Focus()
	m.textArea.SetWidth(state.WindowSize.Width)
	m.textArea.SetHeight(state.WindowSize.Height - 1)
	m.help.Width = state.WindowSize.Width
	m.textArea.CharLimit = 0 // unlimited
	m.help.ShowAll = false

	return m
}

func (m addModel) Init() tea.Cmd {
	return nil
}

func (m addModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			if m.textArea.Focused() {
				m.textArea.Blur()
				return m, nil
			} else {
				return NewListModel(true), nil
			}
		// Can fix this functionality later
		// case "ctrl+s":
		// 	log.Printf("Saving update: %v", m.textArea.Value())
		// 	return models["list"], m.SaveUpdateCmd
		case "enter":
			if !m.textArea.Focused() {
				return NewListModel(true), SaveUpdate(m.textArea.Value())
			}
		case "w":
			if !m.textArea.Focused() {
				m.textArea.Focus()
				return m, textarea.Blink
			}
		}
	case tea.WindowSizeMsg:
		m.textArea.SetWidth(state.WindowSize.Width)
		m.textArea.SetHeight(state.WindowSize.Height - 1)
		m.help.Width = state.WindowSize.Width
	}

	var cmd tea.Cmd
	m.textArea, cmd = m.textArea.Update(msg)
	return m, cmd
}

func (m addModel) View() string {
	helpView := m.help.View(m)
	textAreaView := m.textArea.View()

	return lipgloss.JoinVertical(lipgloss.Left, textAreaView, helpView)
}

func (m addModel) ShortHelp() []key.Binding {
	if m.textArea.Focused() {
		return getAddModelWriteModeKeys()
	}
	return getAddModelViewModeKeys()
}

// Noop to satisfy the interface
func (k addModel) FullHelp() [][]key.Binding { return nil }
