package skyclilib

// Describe os command. Core will execute this command before project generation
type OsCommand struct {
	// Name of command like go, poetry, npm, etc
	Name string

	// Args and flags of command. Flugs will be applied in same order
	Args []CommandArg
}

// Return required and optional command arguments and flags
func (cmd *OsCommand) ExtructArgs() map[string]bool {
	args := make(map[string]bool)
	for _, arg := range cmd.Args {
		args[arg.Name] = arg.NeedGetFromUser
	}
	return args
}

// Describe argument or flag
type CommandArg struct {
	// Name of argument or flag. If it should be get from user, name it like "version".
	// If if this arg is required, enter value like "v1.0.0"
	Name string

	// Is argument or flag should be gotten from user. If true, core will ask for value from user. All previously entered values will be shown
	NeedGetFromUser bool
}
