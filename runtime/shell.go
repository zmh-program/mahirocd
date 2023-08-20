package runtime

import (
	"os/exec"
	"runtime"
)

type Shell struct {
	Path     string
	Commands []string
	Quit     bool
}

func NewShell(path string, commands []string) *Shell {
	return &Shell{
		Path:     path,
		Commands: commands,
		Quit:     false,
	}
}

func (s *Shell) run() string {
	buffer := NewBuffer()

	cwd := s.Path
	for _, command := range s.Commands {
		buffer.PushCommand(command)

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", "/C", command)
		} else {
			cmd = exec.Command("sh", "-c", command)
		}
		cmd.Dir = cwd
		output, err := cmd.CombinedOutput()
		if err != nil {
			buffer.PushError(err)
			buffer.PushExitCode(cmd.ProcessState.ExitCode())
			return buffer.StringAll()
		} else {
			cwd = cmd.Dir
		}
		buffer.PushOutput(output)
	}

	return buffer.StringAll()
}

func (s *Shell) Run() string {
	s.Quit = false
	response := s.run()
	s.Quit = true
	return response
}

func (s *Shell) RunAsync() {
	go s.Run()
}

func (s *Shell) IsRunning() bool {
	return !s.Quit
}
