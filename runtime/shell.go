package runtime

import (
	"log"
	"os/exec"
	"strings"
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

func (s *Shell) run() (response string, err error) {
	buffer := NewBuffer()

	for _, command := range s.Commands {
		buffer.PushCommand(command)

		cmd := exec.Command(strings.Split(command, " ")[0], strings.Split(command, " ")[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			buffer.PushError(err)
			buffer.PushExitCode(cmd.ProcessState.ExitCode())
			return buffer.StringAll(), err
		}
		buffer.PushOutput(output)
	}

	return buffer.StringAll(), nil
}

func (s *Shell) Run() (response string, err error) {
	s.Quit = false
	response, err = s.run()
	s.Quit = true
	return response, err
}

func (s *Shell) RunAsync() {
	go func() {
		_, err := s.Run()
		if err != nil {
			log.Println(err)
		}
	}()
}

func (s *Shell) IsRunning() bool {
	return !s.Quit
}
