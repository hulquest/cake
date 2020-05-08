package progress

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type chanWriter struct {
	ch chan string
}

func NewChanWriter(ch chan string) *chanWriter {
	return &chanWriter{ch}
}

func (w *chanWriter) Chan() <-chan string {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := len(p)
	logString := string(p)
	logString = strings.TrimSuffix(logString, "\n")
	w.ch <- logString
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

type LogrusFormat struct {
}

func (l *LogrusFormat) Format(entry *logrus.Entry) ([]byte, error) {

	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}
