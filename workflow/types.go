package workflow

type GithubWebhook struct {
	HookId     int64 `json:"hook_id"`
	Repository struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
	}
}

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
