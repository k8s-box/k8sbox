// Package structs contain every k8sbox public structs
package structs

// Environment is your environment in a struct
type Environment struct {
	Name             string            `toml:"name"`
	ID               string            `toml:"id"`
	Namespace        string            `toml:"namespace"`
	Boxes            []Box             `toml:"boxes"`
	TempDirectory    string            `toml:"-"`
	Variables        string            `toml:"variables"`
	LoadBoxesFrom    string            `toml:"load_boxes_from"`
	LoadBoxesHeaders map[string]Header `toml:"load_boxes_headers"`
}

// EnvironmentService is a public EnvironmentService
type EnvironmentService struct {
	DeployEnvironment         func(*Environment, bool) error
	DeleteEnvironment         func(*Environment) error
	CreateTempDeployDirectory func(*Environment, bool) (string, error)
	ValidateEnvironment       func(*Environment) error
	ExpandVariables           func(*Environment)
}

// GetEnvironmentAliases return a slice of environment model name aliases
func GetEnvironmentAliases() []string {
	return []string{"environment", "environments", "env"}
}

func GetAvailableDownloadSchemes() []string {
	return []string{"http", "https"}
}
