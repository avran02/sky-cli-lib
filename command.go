package skyclilib

// Describe os command. Core will execute this command before project generation
type OsCommand struct {
	// Name of command like go, poetry, npm, etc
	Name string

	// Args and flags of command. Flugs will be applied in same order
	Args []CommandArg
}

// Describe argument or flag
type CommandArg struct {
	// Desrcribes type of argument or flag.
	//
	// If it should be gotten from user, use FromUser struct.
	//
	// If if user can choose if arg needed, use FromUserBool struct.
	//
	// If you want to use predefined value use FromPlugin struct
	Source source

	// Name of argument or flag. You can keep it empty with FromPlugin source
	Name string

	// Value of argument or flag. You shuold define this field only with FromPlugin source.
	// In other cases value will be overwritten by user's value
	Value string
}

// used if you need to get value from user
//
// example: enter plugin url: <user input>
type FromUser struct{}

func (FromUser) Get() string {
	return "FromUser"
}

// used if you want to ask user if this flug needed
//
// example: -l needed for command 'ls' [Y/n]
type FromUserBool struct{}

func (FromUserBool) Get() string {
	return "FromUserBool"
}

// used if you don'n want to ask user value
//
// example: init (for go mod <init> command)
type FromPlugin struct{}

func (FromPlugin) Get() string {
	return "FromPlugin"
}

// Should be implemented only in FromUser, FromUserBool and FromPlugin
type source interface {
	Get() string
}
