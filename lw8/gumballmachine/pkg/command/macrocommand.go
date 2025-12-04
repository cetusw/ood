package command

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand() *MacroCommand {
	return &MacroCommand{commands: []Command{}}
}

func (mc *MacroCommand) AddCommand(cmd Command) {
	mc.commands = append(mc.commands, cmd)
}

func (mc *MacroCommand) Execute() {
	for _, cmd := range mc.commands {
		cmd.Execute()
	}
}
