package word

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type words struct {
	w []string
}

func Read(path string) (*words, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var w words
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		w.w = append(w.w, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &w, nil
}

func (w *words) Get() []string {
	return w.w
}

func (w *words) Reverse() {
	if len(w.w) == 0 || len(w.w) == 1 {
		return
	}
	final := make([]string, 0, len(w.w))
	for _, word := range w.w {
		split := strings.Split(word, " - ")
		if len(split) < 2 {
			return
		}
		if rand.Intn(2) == 0 {
			final = append(final, split[0]+" - "+split[1])
		} else {
			final = append(final, split[1]+" - "+split[0])
		}
	}
	w.w = final
}

func (w *words) Shuffle() {
	for i := range w.w {
		j := rand.Intn(i + 1)
		w.w[i], w.w[j] = w.w[j], w.w[i]
	}
}
