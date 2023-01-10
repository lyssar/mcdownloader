package utils

import "github.com/chzyer/readline"

type NoBellStdout struct{}

func (n *NoBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *NoBellStdout) Close() error {
	return readline.Stdout.Close()
}

func NewNoBellStdout() NoBellStdout {
	return NoBellStdout{}
}
