package workflow

func (m *Manager) Load() {
	config := GetWorkflowConfig()
	m.Workflows = make(map[string]Workflow)
	for _, path := range config {
		workflow := ReadWorkflow(path)
		if workflow != nil {
			m.Workflows[workflow.Name] = *workflow
		}
	}
}

func (m *Manager) Get(name string) *Workflow {
	if workflow, ok := m.Workflows[name]; ok {
		return &workflow
	}
	return nil
}

func (m *Manager) List() []Workflow {
	list := make([]Workflow, 0)
	for _, workflow := range m.Workflows {
		list = append(list, workflow)
	}
	return list
}

func (m *Manager) Add(workflow Workflow) {
	m.Workflows[workflow.Name] = workflow
}

func (m *Manager) Remove(name string) {
	delete(m.Workflows, name)
}

func (m *Manager) Refresh() {
	m.Load()
}

func NewManager() Manager {
	manager := Manager{}
	manager.Load()
	return manager
}

func (m *Manager) RunAsync(name string) bool {
	workflow := m.Get(name)
	if workflow == nil {
		return false
	}

	workflow.RunAsync()
	return true
}
