package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type TimeZone string

const (
	EST TimeZone = "EST"
	GMT TimeZone = "GMT"
	CST TimeZone = "CST"
)

type Exchange struct {
	name     string
	timezone TimeZone
}

type model struct {
	local_time TimeZone
	exchanges  []Exchange
}

var (
	closedTradingDay = lipgloss.NewStyle().
				Width(15).
				Height(15).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.HiddenBorder())

	openTradingDay = lipgloss.NewStyle().
			Width(15).
			Height(15).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.DoubleBorder())
)

func main() {
	p := tea.NewProgram(initModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}

func initModel() model {
	return model{
		local_time: EST,
		exchanges: []Exchange{
			{
				name:     "Toronto Stock Exchange",
				timezone: EST,
			},
			{
				name:     "London Stock Exchange",
				timezone: GMT,
			},
			{
				name:     "Chicago Stock Exchange",
				timezone: CST,
			},
		},
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
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	s += lipgloss.JoinHorizontal(lipgloss.Top, openTradingDay.Render("TORONTO"), closedTradingDay.Render("CHICAGO"))

	return s
}
