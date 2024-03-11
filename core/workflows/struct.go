package workflows

type AppInfo struct {
	Id         string `json:"id"`
	ConfigName string `json:"config_name"`
	Title      string `json:"title"`
	// Name is the name of the model.
	Name string `json:"name"`
	// Version is the version of the model.
	Version []string `json:"version"`
	// Description is the description of the model.
	Description string `json:"description"`
	// Category is the category of the model.
	Category string `json:"category"`
	// Author is the author of the model.
	Author string `json:"author" yaml:"author"`
	// OrganizationName is the name of the organization associated with the model.
	OrganizationName string `json:"organization_name" yaml:"organization_name"`
	// Icon is the icon associated with the model.
	Icon string `json:"icon"`

	AllowMultipleInstances string `json:"allow_multiple_instances" yaml:"allow_multiple_instances"`

	Tags    []string          `json:"tags" yaml:"tags"`
	Actions map[string]string `json:"actions" yaml:"actions"`
	Options map[string]any    `json:"options" yaml:"options"`
}

type WorkflowsInfo struct {
	Id               string
	ScriptName       string
	Title            string
	Name             string
	Version          string
	Description      string
	Author           string
	OrganizationName string
	Icon             string
	Actions          map[string]string
	// Steps            []Job
	// Build            []Job
	// Install          []Job
	// Uninstall        []Job
}

type Workflows struct {
	Id               string
	ScriptName       string
	Title            string
	Name             string
	Version          string
	Description      string
	Author           string
	OrganizationName string
	Icon             string
	Actions          map[string]string
	Steps            []Job
	Build            []Job
	Install          []Job
	Uninstall        []Job
	Upgrade          []Job
}

type Job struct {
	Name string
	Run  string
}
