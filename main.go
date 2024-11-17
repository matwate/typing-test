package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/matwate/typing-test/analysis"
	"github.com/matwate/typing-test/metrics"
)

type TestState int

const (
	Typing TestState = iota
	Stats
)

var prompt string = select_random_writing_prompt()

type model struct {
	something    int
	player_typed string
	strokes      []metrics.Stroke
	state        TestState
	ta           textarea.Model
}

func initialModel() model {
	ta := textarea.New()
	ta.Focus()
	ta.CharLimit = 140
	ta.SetWidth(100)
	return model{
		something:    0,
		player_typed: "",
		state:        Typing,
		ta:           ta,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if len(m.player_typed) > 0 {
				m.player_typed = m.player_typed[:len(m.player_typed)-1]
				m.strokes = append(m.strokes, metrics.Stroke{Char: "backspace", Time: time.Now()})
			}
		case "enter":
			// Finish the typing phase and go to the stats phase
			m.state = Stats
		default:
			switch m.state {
			case Typing:
				if len(msg.String()) == 1 {
					m.player_typed += msg.String()
					m.strokes = append(m.strokes, metrics.Stroke{Char: msg.String(), Time: time.Now()})

				}
			}
		}
	}

	m.ta, cmd = m.ta.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.state {
	case Typing:
		// Header
		s := "Your typing prompt is:"
		// Typing area
		s += "\n\n"
		s += prompt
		s += "\n\n"
		s += m.ta.View()
		return s
	case Stats:

		matched_text := analysis.Construct_target_sentence(m.player_typed)
		// Calculate the accuracy
		accuracy := analysis.Accuracy(m.player_typed)
		rpm := metrics.GetRawWpm(m.strokes)
		wpm := metrics.GetWpm(m.strokes, accuracy)
		s := fmt.Sprintf("Your accuracy is: %f", accuracy)
		s += "\n\n"
		s += fmt.Sprintf("Your writing prompt was: %s \n", prompt)
		s += "You typed: \n"
		s += m.player_typed
		s += "\n\n"
		s += "We thought you were trying to type: \n"
		s += matched_text
		s += "\n\n"
		s += "Your typing speed was: "
		s += fmt.Sprintf("%d", wpm)
		s += " words per minute"
		s += "\n\n"
		s += "Your raw typing speed was: "
		s += fmt.Sprintf("%d", rpm)
		s += " words per minute"
		s += "\n\n"
		s += "You took: "
		s += fmt.Sprintf("%f", metrics.GetTimeTaken(m.strokes))
		s += " seconds"
		s += "\n\n"
		s += "You lost: "
		s += fmt.Sprintf("%f", metrics.TimeLostByFixingMistakes(m.strokes))
		s += " seconds, by fixing mistakes"
		s += "\n\n"
		s += "You took: "
		s += fmt.Sprintf("%f", metrics.ThinkingTime(m.strokes))
		s += " seconds, thinking"
		s += "\n\n"
		s += "Press 'esc' to exit"
		return s

	}
	return ""
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func select_random_writing_prompt() string {
	prompts := []string{}
	file, err := os.Open("./prompts.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prompts = append(prompts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return prompts[rand.Intn(len(prompts))]
}
