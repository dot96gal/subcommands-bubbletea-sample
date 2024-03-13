package bubbletea

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewPrintProgram(ctx context.Context, delay int, args []string) *tea.Program {
	m := newPrintModel(ctx, delay, args)

	return tea.NewProgram(m)
}

var _ tea.Msg = (*printMsg)(nil)

type printMsg struct {
	output string
}

var _ tea.Model = (*printModel)(nil)

type printModel struct {
	ctx     context.Context
	delay   int
	args    []string
	output  string
	spinner spinner.Model
	loading bool
}

func newPrintModel(ctx context.Context, delay int, args []string) printModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return printModel{
		ctx:     ctx,
		delay:   delay,
		args:    args,
		output:  "",
		spinner: s,
		loading: true,
	}
}

func (m printModel) Init() tea.Cmd {
	return tea.Batch(
		m.printCmd,
		m.spinner.Tick,
	)
}

func (m printModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case printMsg:
		m.loading = false
		m.output = msg.output
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m printModel) View() string {
	if m.loading {
		return fmt.Sprintf("%s Loading ...\n", m.spinner.View())
	}

	return fmt.Sprintf("%s\n", m.output)
}

var _ tea.Cmd = printModel{}.printCmd

func (m printModel) printCmd() tea.Msg {
	time.Sleep(time.Duration(m.delay) * time.Millisecond)
	output := strings.Join(m.args, " ")

	return printMsg{output: output}
}
