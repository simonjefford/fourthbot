package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/simonjefford/fourthbot"
	"github.com/simonjefford/fourthbot/mock/responders"
)

type terminalWriter struct {
	io.Writer
}

func newTerminalWriter() *terminalWriter {
	return &terminalWriter{
		os.Stdout,
	}
}

func (t *terminalWriter) WriteStatus(s int) {
	fmt.Fprintln(t, "Written status - ", s)
}

func main() {
	r := fourthbot.NewRobot()
	rw := newTerminalWriter()
	responders.RegisterAll(r)

	l, err := readline.NewEx(&readline.Config{
		Prompt:      "\033[31m»\033[0m ",
		HistoryFile: ".history",
	})
	if err != nil {
		panic(err)
	}
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		c, err := fourthbot.ParseCommand(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = r.HandleCommand(context.Background(), c, rw)
		if err != nil {
			fmt.Println(err)
		}
	}
}
