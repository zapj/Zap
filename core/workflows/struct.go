package workflows

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
