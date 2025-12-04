package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gumballmachine/pkg/command"
)

type menuItem struct {
	shortcut    string
	description string
	command     command.Command
}

type Menu struct {
	items                []menuItem
	exit                 bool
	reader               *bufio.Scanner
	isRecording          bool
	newMacroShortcut     string
	newMacroDescription  string
	currentMacroCommands []command.Command
}

func NewMenu() *Menu {
	return &Menu{
		reader:               bufio.NewScanner(os.Stdin),
		items:                make([]menuItem, 0),
		currentMacroCommands: make([]command.Command, 0),
	}
}

func (m *Menu) AddItem(shortcut, description string, cmd command.Command) {
	m.items = append(m.items, menuItem{shortcut, description, cmd})
}

func (m *Menu) Run() {
	m.ShowInstructions()
	for {
		if m.isRecording {
			fmt.Print("rec> ")
		} else {
			fmt.Print("> ")
		}

		if !m.reader.Scan() {
			break
		}
		line := m.reader.Text()
		if !m.executeCommand(line) {
			break
		}
	}
}

func (m *Menu) ShowInstructions() {
	fmt.Println("Commands list:")
	for _, item := range m.items {
		fmt.Printf("  %s: %s\n", item.shortcut, item.description)
	}
	fmt.Println("Macro commands:")
	fmt.Println("  begin_macro: Start recording a new macro command.")
	fmt.Println("  end_macro:   Stop recording and save the macro command.")
	fmt.Println()
}

func (m *Menu) Exit() {
	m.exit = true
}

func (m *Menu) handleBeginMacro() {
	if m.isRecording {
		fmt.Println("Already recording a macro. Type 'end_macro' to finish.")
		return
	}

	fmt.Print("Enter macro shortcut: ")
	if !m.reader.Scan() {
		return
	}
	shortcut := m.reader.Text()

	for _, item := range m.items {
		if item.shortcut == shortcut {
			fmt.Println("Error: Command with this shortcut already exists.")
			return
		}
	}
	if shortcut == "begin_macro" || shortcut == "end_macro" {
		fmt.Println("Error: Cannot use reserved keywords for shortcut.")
		return
	}

	fmt.Print("Enter macro description: ")
	if !m.reader.Scan() {
		return
	}
	description := m.reader.Text()

	m.isRecording = true
	m.newMacroShortcut = shortcut
	m.newMacroDescription = description
	m.currentMacroCommands = []command.Command{}
	fmt.Println("Macro recording started. Type 'end_macro' to finish.")
}

func (m *Menu) handleEndMacro() {
	if !m.isRecording {
		fmt.Println("Error: Not recording a macro. Type 'begin_macro' to start.")
		return
	}

	newMacro := command.NewMacroCommand()
	for _, cmd := range m.currentMacroCommands {
		newMacro.AddCommand(cmd)
	}

	m.AddItem(m.newMacroShortcut, m.newMacroDescription, newMacro)
	fmt.Printf("Macro '%s' saved successfully.\n", m.newMacroShortcut)

	m.isRecording = false
	m.newMacroShortcut = ""
	m.newMacroDescription = ""
	m.currentMacroCommands = nil
}

func (m *Menu) executeCommand(shortcut string) bool {
	m.exit = false
	shortcut = strings.TrimSpace(shortcut)

	if shortcut == "begin_macro" {
		m.handleBeginMacro()
		return !m.exit
	}

	if shortcut == "end_macro" {
		m.handleEndMacro()
		return !m.exit
	}

	var foundItem *menuItem
	for i := range m.items {
		if m.items[i].shortcut == shortcut {
			foundItem = &m.items[i]
			break
		}
	}

	if m.isRecording {
		if foundItem != nil {
			m.addCommandToMacro(shortcut, foundItem)
		} else {
			fmt.Println("Unknown command, not added to macro.")
		}
	} else {
		if foundItem != nil {
			foundItem.command.Execute()
		} else {
			fmt.Println("Unknown command")
		}
	}

	return !m.exit
}

func (m *Menu) addCommandToMacro(shortcut string, item *menuItem) {
	if shortcut == "help" || shortcut == "exit" {
		fmt.Printf("Command '%s' cannot be recorded into a macro.\n", shortcut)
	} else {
		fmt.Printf("Added command '%s' to macro.\n", shortcut)
		m.currentMacroCommands = append(m.currentMacroCommands, item.command)
	}

}
