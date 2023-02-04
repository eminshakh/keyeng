package app

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/eminshakh/keyeng/internal/terminal"
	"github.com/eminshakh/keyeng/internal/word"
)

const (
	enter     = 13
	backspace = 127
	ctrlZ     = 26
)

func Run() error {
	w, err := word.Read("words.txt")
	if err != nil {
		return err
	}
	w.Shuffle()
	w.Reverse()
	words := w.Get()
	err = keygen(words)
	if err != nil {
		return err
	}
	return nil
}

func keygen(words []string) error {
	symbols := 0
	typos := 0
	err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	now := time.Now()
	var char rune
Z:
	for _, s := range words {
		terminal.Clear()
		fmt.Printf(terminal.MoveCursor(terminal.HiYellowString(s), 1, 1))
		line := []rune(s)
		prev := New(len(line))
	E:
		for i := 0; i < len(line); i++ {
			char, _, err = bufio.NewReader(os.Stdin).ReadRune()
			if err != nil {
				return err
			}
			switch char {
			case enter:
				terminal.Clear()
				prev.Empty()
				break E
			case ctrlZ:
				break Z
			case backspace:
				peek, ok := prev.Peek()
				// if red symbol
				if ok && peek[2] == 57 && peek[3] == 49 {
					prev.Pop()
					i = prev.Len() - 1
					terminal.Clear()
					fmt.Printf(terminal.MoveCursor(fmt.Sprintf("%s%s", prev.String(), terminal.HiYellowString(string(line[prev.Len():]))), 1, prev.Len()+1))
				} else {
					i--
				}
			case line[i]:
				symbols++
				prev.Push(terminal.WhiteString(string(line[i])))
				terminal.Clear()
				fmt.Printf(terminal.MoveCursor(fmt.Sprintf("%s%s", prev.String(), terminal.HiYellowString(string(line[prev.Len():]))), 1, i+2))
			default:
				typos++
				prev.Push(terminal.HiRedString(string(line[i])))
				terminal.Clear()
				fmt.Printf(terminal.MoveCursor(fmt.Sprintf("%s%s", prev.String(), terminal.HiYellowString(string(line[prev.Len():]))), 1, i+2))
			}
		}
	}
	if symbols == 0 && typos > 0 {
		symbols = typos
	}
	spm := float64(symbols) / time.Since(now).Minutes()
	accuracy := (float64(symbols-typos) / float64(symbols)) * 100
	if symbols == 0 {
		accuracy = 0
	}
	terminal.Clear()
	fmt.Println(terminal.HiBlueString(fmt.Sprintf("Speed:[%d] Accuracy:[%.0f%%] Typos:[%d] Elapsed:[%.2f min]", int(spm), accuracy, typos, time.Since(now).Minutes())))
	return nil
}
