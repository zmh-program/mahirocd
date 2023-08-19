package runtime

import (
	"fmt"
	"strings"
)

type Buffer struct {
	Length int
	Stack  []string
}

func NewBuffer() *Buffer {
	return &Buffer{
		Length: 0,
		Stack:  []string{},
	}
}

func (b *Buffer) Push(s string) {
	b.Stack = append(b.Stack, s)
	b.Length++
}

func (b *Buffer) Pop() string {
	if b.Length == 0 {
		return ""
	}

	s := b.Stack[b.Length-1]
	b.Stack = b.Stack[:b.Length-1]
	b.Length--
	return s
}

func (b *Buffer) Peek() string {
	if b.Length == 0 {
		return ""
	}

	return b.Stack[b.Length-1]
}

func (b *Buffer) Clear() {
	b.Stack = []string{}
	b.Length = 0
}

func (b *Buffer) IsEmpty() bool {
	return b.Length == 0
}

func (b *Buffer) IsNotEmpty() bool {
	return b.Length > 0
}

func (b *Buffer) GetStack() []string {
	return b.Stack
}

func (b *Buffer) GetLength() int {
	return b.Length
}

func (b *Buffer) SetStack(stack []string) {
	b.Stack = stack
	b.Length = len(stack)
}

func (b *Buffer) SetLength(length int) {
	b.Length = length
}

func (b *Buffer) String() string {
	return b.Peek()
}

func (b *Buffer) StringAll() string {
	return strings.Join(b.Stack, "\n")
}

func (b *Buffer) PushCommand(cmd string) {
	b.Push("> " + cmd)
}

func (b *Buffer) PushOutput(output []byte) {
	b.Push(string(output))
}

func (b *Buffer) PushError(err error) {
	b.Push("Error: " + err.Error())
}

func (b *Buffer) PushReadError(err error) {
	b.Push("Error while reading output: " + err.Error())
}

func (b *Buffer) PushQuitError(err error) {
	b.Push("Error while quitting: " + err.Error())
}

func (b *Buffer) PushExitCode(code int) {
	b.Push(fmt.Sprintf("\nProcess finished with exit code %d\n", code))
}
