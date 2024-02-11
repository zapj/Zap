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
	Author string `json:"author"`
	// OrganizationName is the name of the organization associated with the model.
	OrganizationName string `json:"organization_name"`
	// Icon is the icon associated with the model.
	Icon string `json:"icon"`

	// Tags is a list of tags associated with the model.
	Tags []string `json:"tags"`
	// Actions is a map of actions and their associated values.
	Actions map[string]string            `json:"actions"`
	Options map[string]map[string]string `json:"options"`
	// Steps            []Job
	// Build            []Job
	// Install          []Job
	// Uninstall        []Job
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
