package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/chzyer/readline"
	"github.com/simonjefford/fourthbot"
	"github.com/simonjefford/fourthbot/mock"
	"github.com/simonjefford/fourthbot/mock/responders"
)

func main() {
	r := fourthbot.NewRobot(mock.NewMockResponseWriter())
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
		err = r.HandleCommand(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}
