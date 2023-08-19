package runtime

import (
	"os/exec"
)

type Shell struct {
	Path     string
	Commands []string
}

func (s *Shell) Run() (response string, err error) {
	instance := exec.Command("bash")
	instance.Dir = s.Path

	buffer := NewBuffer()

	stdin, err := instance.StdinPipe()
	if err != nil {
		return "", err
	}
	defer stdin.Close()

	stdout, err := instance.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	if err := instance.Start(); err != nil {
		return "", err
	}

	for _, command := range s.Commands {
		buffer.PushCommand(command)

		_, err := stdin.Write([]byte(command + "\n"))
		if err != nil {
			buffer.PushError(err)
			break
		}

		output := make([]byte, 1024*1024*5) // 5 MiB buffer size
		n, err := stdout.Read(output)
		if err != nil {
			buffer.PushReadError(err)
			break
		}

		buffer.PushOutput(output[:n])
	}

	if err := instance.Wait(); err != nil {
		buffer.PushQuitError(err)
	}

	buffer.PushExitCode(instance.ProcessState.ExitCode())
	return buffer.StringAll(), nil
}
