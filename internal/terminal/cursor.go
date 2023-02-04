package terminal

import "fmt"

func MoveCursor(s string, x, y int) string {
	return fmt.Sprintf("%s \033[%d;%dH", s, x, y)
}
