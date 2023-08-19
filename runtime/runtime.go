package runtime

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
	hash := md5.New()
	hash.Write([]byte(r.Name))
	return hex.EncodeToString(hash.Sum(nil))
}

func (r *Runtime) GetPath() string {
	return r.Path
}

func (r *Runtime) Exec() string {
	shell := NewShell(r.Path, r.Command)
	response, err := shell.Run()
	if err != nil {
		return fmt.Sprintf("runtime error occurred: %s", err.Error())
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
