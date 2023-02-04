package terminal

import "fmt"

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

const (
	FgHiBlack = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func BlackString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgBlack, text)
}

func RedString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgRed, text)
}

func GreenString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, text)
}

func YellowString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgYellow, text)
}

func BlueString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgBlue, text)
}

func MagentaString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgMagenta, text)
}

func CyanString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgCyan, text)
}

func WhiteString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgWhite, text)
}

func HiBlackString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiBlack, text)
}

func HiRedString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiRed, text)
}

func HiYellowString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiYellow, text)
}

func HiGreenString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiGreen, text)
}

func HiBlueString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiBlue, text)
}

func HiMagentaString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiMagenta, text)
}

func HiCyanString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiCyan, text)
}

func HiWhiteString(text string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgHiWhite, text)
}
