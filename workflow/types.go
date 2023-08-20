package workflow

type Workflow struct {
	Name  string `json:"name"`
	Repo  string `json:"repo"`
	Path  string `json:"path"`
	Steps []Step `json:"steps"`
}

type Step struct {
	Name string `json:"name"`
	Run  string `json:"run"`
}

type Manager struct {
	Workflows map[string]Workflow
}
