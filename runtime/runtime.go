package runtime

import (
	"fmt"
	"mahirocd/utils"
)

type Runtime struct {
	Name    string
	Path    string
	Command []string
}

func NewRuntime(name string, path string, command []string) *Runtime {
	return &Runtime{
		Name:    name,
		Path:    path,
		Command: command,
	}
}

func (r *Runtime) GetName() string {
	return r.Name
}

func (r *Runtime) GetCommand() []string {
	return r.Command
}

func (r *Runtime) GetHash() string {
	return utils.Md5Encode(r.Name)
}

func (r *Runtime) GetPath() string {
	return r.Path
}

func (r *Runtime) Exec() string {
	shell := NewShell(r.Path, r.Command)
	response, err := shell.Run()
	if err != nil {
		return fmt.Sprintf("%s\nruntime error occurred: %s", response, err)
	} else {
		return response
	}
}

func (r *Runtime) ExecWithLog() (string, error) {
	response := r.Exec()
	return response, WriteLog(r.GetHash(), response)
}

func (r *Runtime) ProcessAsync() {
	go func() {
		_, err := r.ExecWithLog()
		if err != nil {
			fmt.Println(err)
		}
	}()
}
