package runtime

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type Runtime struct {
	Name        string
	Path        string
	CommandList [][]string
}

func NewRuntime(name string, path string, commandList [][]string) *Runtime {
	return &Runtime{
		Name:        name,
		Path:        path,
		CommandList: commandList,
	}
}

func (r *Runtime) GetName() string {
	return r.Name
}

func (r *Runtime) GetCommandList() [][]string {
	return r.CommandList
}

func (r *Runtime) GetHash() string {
	hash := md5.New()
	hash.Write([]byte(r.Name))
	return string(hash.Sum(nil))
}

func (r *Runtime) GetPath() string {
	return r.Path
}

func (r *Runtime) Exec() string {
	stack := make([]string, 0)
	for _, command := range r.CommandList {
		shell := NewShell(r.Path, command)
		response, err := shell.Run()
		if err != nil {
			stack = append(stack, fmt.Sprintf("runtime error occurred: %s", err.Error()))
			return strings.Join(stack, "\n")
		} else {
			stack = append(stack, response)
		}
	}
	return strings.Join(stack, "\n")
}

func (r *Runtime) ExecWithLog() (string, error) {
	response := r.Exec()
	return response, WriteLog(r.GetHash(), response)
}
