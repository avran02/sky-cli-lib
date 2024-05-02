/*
Common library for plugins and sky-cli core.
Most nessary part of this package is common interface.
*/
package skyclilib

// Describe generating project structure
type PluginConfiger interface {
	// Return names of required files and folders
	// GetRequiredParams() []string

	// Return os commands that will be used in project.
	// This commands will be executed before generation.
	// It may be something like: "go mod init", "poetry init", etc
	GetOsCommands() []OsCommand

	// Return JSON-like virtual file system
	GetVirtualFs() *Folder
}
