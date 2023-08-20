package workflow

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"mahirocd/runtime"
	"mahirocd/utils"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkflowConfig() []string {
	list, err := os.ReadDir("workflows")
	if err != nil {
		panic(err)
	}

	config := make([]string, 0)
	for _, item := range list {
		ext := filepath.Ext(item.Name())
		if !item.IsDir() && (ext == ".yml" || ext == ".yaml") {
			config = append(config, fmt.Sprintf("workflows/%s", item.Name()))
		}
	}

	return config
}

func ReadWorkflow(path string) *Workflow {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var workflow Workflow
	if err := yaml.Unmarshal(file, &workflow); err != nil {
		return nil
	}

	return &workflow
}

func (w *Workflow) Save() {
	file, err := os.Create(fmt.Sprintf("workflows/%s.yml", w.Name))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(w); err != nil {
		panic(err)
	}
}

func GetCommand(data string) string {
	stack := make([]string, 0)
	for _, char := range strings.Split(data, "\n") {
		char = strings.TrimSpace(char)
		if len(char) == 0 {
			continue
		}
		stack = append(stack, strings.TrimSpace(char))
	}
	return strings.Join(stack, fmt.Sprintf(" %s ", utils.GetCommandSeparator()))
}

func (w *Workflow) GetCommands() []string {
	commands := make([]string, 0)
	for _, step := range w.Steps {
		commands = append(commands, GetCommand(step.Run))
	}
	return commands
}

func (w *Workflow) RunAsync() {
	instance := runtime.NewRuntime(w.Repo, w.Path, w.GetCommands())
	instance.ProcessAsync()
}
